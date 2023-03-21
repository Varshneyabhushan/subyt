package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"videosservice/repository/elasticsearch"
	"videosservice/repository/mongo"
)

// CompositeRepository
// should implement all the functionalities using its sub repositories
type CompositeRepository struct {
	mongoRepository *mongo.MongoRepository
	esRepository    *elasticsearch.Repository
}

func NewCompositeRepository(
	mongoRepository *mongo.MongoRepository,
	esRepository *elasticsearch.Repository,
) Repository {
	return &CompositeRepository{
		mongoRepository: mongoRepository,
		esRepository:    esRepository,
	}
}

func (repo *CompositeRepository) Add(videos []Video) (int, error) {
	var mongoVideos []mongo.Video
	var esVideos []elasticsearch.Video
	for _, video := range videos {
		mongoVideos = append(mongoVideos, toMongoVideo(video))
		esVideos = append(esVideos, toEsVideo(video))
	}

	//add videos to es in background
	go func() {
		err := repo.esRepository.Add(esVideos)
		if err != nil {
			log.Println("error while adding videos to es : ", err)
		}
	}()

	return repo.mongoRepository.Add(mongoVideos)
}

func (repo *CompositeRepository) Get(skip, limit int64) ([]Video, error) {
	mongoVideos, err := repo.mongoRepository.Get(skip, limit, bson.M{"publishedat": -1})
	if err != nil {
		return nil, err
	}

	return ToVideos(mongoVideos), nil
}

func (repo *CompositeRepository) Search(term string, skip, limit int) ([]Video, error) {
	esVideos, err := repo.esRepository.Search(term, skip, limit)
	if err != nil || len(esVideos) == 0 {
		return []Video{}, err
	}

	var ids []string
	for _, esVideo := range esVideos {
		ids = append(ids, esVideo.Id)
	}

	mongoVideos, err := repo.mongoRepository.FindByIds(ids)
	if err != nil {
		return nil, err
	}

	return ToVideos(mongoVideos), nil
}
