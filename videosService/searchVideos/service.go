package searchVideos

import (
	"videosservice/addVideos"
	elasticsearch2 "videosservice/storageServices/elasticsearch"
	"videosservice/storageServices/mongo"
)

func MakeSearchService(
	mongoService mongo.Collection[mongo.Video],
	esService elasticsearch2.Index[elasticsearch2.Video],
) Service {
	return func(term string, skip, limit int) ([]addVideos.VideoResponse, error) {
		var result []addVideos.VideoResponse
		esVideos, err := esService.Search(term, skip, limit)
		if err != nil || len(esVideos) == 0 {
			return result, err
		}

		var ytIds []string
		for _, esVideo := range esVideos {
			ytIds = append(ytIds, esVideo.Id)
		}

		mongoVideos, err := mongoService.FindByKey("ytid", ytIds)
		if err != nil {
			return nil, err
		}

		return addVideos.ToVideoResponses(mongoVideos), nil
	}
}
