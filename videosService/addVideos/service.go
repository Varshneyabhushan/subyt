package addVideos

import (
	"log"
	"videosservice/storageServices/elasticsearch"
	"videosservice/storageServices/mongo"
)

func MakeAddService(
	mongoService mongo.Collection[mongo.Video],
	esService elasticsearch.Index[elasticsearch.Video],
) Service {
	return func(videos []VideoResponse) (int, error) {
		var mongoVideos []mongo.Video
		var esVideos []elasticsearch.Video
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
