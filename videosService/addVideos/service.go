package addVideos

import (
	"log"
	es2 "videosservice/repository/elasticsearch"
	mongo2 "videosservice/repository/mongo"
	"videosservice/storageServices/elasticsearch"
	"videosservice/storageServices/mongo"
)

func MakeAddService(
	mongoService mongo.Collection[mongo2.Video],
	esService elasticsearch.Index[es2.Video],
) Service {
	return func(videos []VideoResponse) (int, error) {
		var mongoVideos []mongo2.Video
		var esVideos []es2.Video
		for _, video := range videos {
			mongoVideos = append(mongoVideos, ToMongoVideo(video))
			esVideos = append(esVideos, ToESVideo(video))
		}

		//add videos to es in background
		go func() {
			err := esService.Add(esVideos)
			if err != nil {
				log.Println("error while adding videos to es : ", err)
			}
		}()

		return mongoService.Add(mongoVideos)
	}
}
