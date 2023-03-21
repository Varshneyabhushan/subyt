package getVideos

import (
	"go.mongodb.org/mongo-driver/bson"
	"videosservice/addVideos"
	mongo2 "videosservice/repository/mongo"
	"videosservice/storageServices/mongo"
)

func MakeGetVideosService(mongoService mongo.Collection[mongo2.Video]) Service {
	return func(skip, limit int64) ([]addVideos.VideoResponse, error) {
		mongoVideos, err := mongoService.Get(skip, limit, bson.M{"publishedat": -1})
		if err != nil {
			return nil, err
		}

		return addVideos.ToVideoResponses(mongoVideos), nil
	}
}
