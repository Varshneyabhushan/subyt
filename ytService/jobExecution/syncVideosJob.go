package jobexecution

import (
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

	delayTracker := NewDelayTracker(config.RequestCoolDown)
	syncCoolDown := config.SyncCoolDown

	return func() (time.Duration, error) {
		response, err := fetcher.GetNext(checkPoint.NextPageToken)
		if err != nil {
			log.Println("error while getting videos : ", err)
			return delayTracker.ExponentialBackOff(), nil
		}

		newCheckPoint, totalValidVideos := NewCheckpoint(*checkPoint, response)

		//add videos to videosService
		if err = service.AddVideos(response.Videos[:totalValidVideos]); err != nil {
			log.Println("error while adding videos : " + err.Error())
			return delayTracker.ProportionalDelay(), nil
		}

		delayTracker.Reset()

		//when checkpoint is not saved, oldCheckpoint is reused
		if err = newCheckPoint.Save(); err != nil {
			log.Println("error while updating checkpoint", err)
			return delayTracker.Delay(), nil
		}
		*checkPoint = newCheckPoint

		//when pagination continues, sync shouldn't cool down
		if len(response.NextPageToken) != 0 {
			return delayTracker.Delay(), nil
		}

		return syncCoolDown, nil
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
