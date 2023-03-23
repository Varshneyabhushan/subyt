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
	return func(term string, skip, limit int) ([]addVideos.VideoResponse, int64, error) {
		var result []addVideos.VideoResponse
		esVideos, count, err := esService.Search(term, skip, limit)
		if err != nil || len(esVideos) == 0 {
			return result, count, err
		}

		var ytIds []string
		for _, esVideo := range esVideos {
			ytIds = append(ytIds, esVideo.Id)
		}

		mongoVideos, err := mongoService.FindByKey("ytid", ytIds)
		if err != nil {
			return nil, count, err
		}

		return addVideos.ToVideoResponses(mongoVideos), count, nil
	}
}
