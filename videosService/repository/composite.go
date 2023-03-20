package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	for _, video := range videos {
		mongoVideos = append(mongoVideos, mongoVideo(video))
	}

	return repo.mongoRepository.Add(mongoVideos)
}

func (repo *CompositeRepository) Get(skip, limit int64) ([]Video, error) {
	mongoVideos, err := repo.mongoRepository.Get(skip, limit, bson.M{"publishedat": -1})
	if err != nil {
		return nil, err
	}

	return ToVideos(mongoVideos), nil
}

func (repo *CompositeRepository) Search(term string) ([]Video, error) {
	esVideos, err := repo.esRepository.Search(term)
	if err != nil {
		return nil, err
	}

	var ids []primitive.ObjectID
	for _, esVideo := range esVideos {
		if objectId, err := primitive.ObjectIDFromHex(esVideo.Id); err != nil {
			ids = append(ids, objectId)
		}
	}

	mongoVideos, err := repo.mongoRepository.FindByIds(ids)
	if err != nil {
		return nil, err
	}

	return ToVideos(mongoVideos), nil
}
