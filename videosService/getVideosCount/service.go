package getVideosCount

import (
	"videosservice/storageServices/mongo"
)

func MakeService(mongoService mongo.Collection[mongo.Video]) Service {
	return func() (int64, error) {
		return mongoService.GetCount()
	}
}
