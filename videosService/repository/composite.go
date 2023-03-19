package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"videosservice/repository/mongo"
)

// CompositeRepository
// should implement all the functionalities using its sub repositories
type CompositeRepository struct {
	mongoRepository *mongo.MongoRepository
}

func NewCompositeRepository(mongoRepository *mongo.MongoRepository) Repository {
	return &CompositeRepository{mongoRepository: mongoRepository}
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

	var videos []Video
	for _, mongoVideo := range mongoVideos {
		videos = append(videos, ToVideo(mongoVideo))
	}

	return videos, nil
}

func (repo *CompositeRepository) Search(term string) ([]Video, error) {
	return []Video{}, nil
}
