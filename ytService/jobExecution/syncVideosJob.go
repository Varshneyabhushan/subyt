package jobexecution

import (
	"errors"
	"log"
	"time"
	"ytservice/env"
	"ytservice/storage"
	"ytservice/videofetcher"
	"ytservice/videosservice"
)

func MakeSyncVideosJob(
	checkPoint *storage.CheckPoint,
	fetcher *videofetcher.VideoFetcher,
	service videosservice.VideosService,
	config env.SchedulerConfig,
) RecurringJob {
	return func() (time.Duration, error) {
		response, err := fetcher.GetNext(checkPoint.NextPageToken)
		if err != nil {
			return 0, errors.New("error while getting videos : " + err.Error())
		}

		totalValidVideos := 0
		*checkPoint, totalValidVideos = NewCheckpoint(*checkPoint, response)

		//add videos to videosService
		if err = service.AddVideos(response.Videos[:totalValidVideos]); err != nil {
			return 0, errors.New("error while adding videos : " + err.Error())
		}

		if err = checkPoint.Save(); err != nil {
			log.Fatal("error while updating checkpoint", err)
		}

		//when pagination continues, no need to cool down much
		if len(response.NextPageToken) != 0 {
			return time.Duration(config.RequestCoolDown) * time.Second, nil
		}

		return time.Duration(config.Period) * time.Second, nil
	}
}

func NewCheckpoint(checkPoint storage.CheckPoint, response videofetcher.VideosResponse) (
	storage.CheckPoint, int) {

	totalValidVideos := videosservice.GetValidVideos(response.Videos, checkPoint.VideoLimit)
	if totalValidVideos != 0 && checkPoint.NextVideoLimit.IsEmpty() {
		checkPoint.NextVideoLimit = response.Videos[0].VideoComparor
	}

	if totalValidVideos < len(response.Videos) {
		checkPoint.VideoLimit = checkPoint.NextVideoLimit
		checkPoint.NextVideoLimit = videosservice.VideoComparor{}
	}

	checkPoint.NextPageToken = response.NextPageToken

	return checkPoint, totalValidVideos
}
