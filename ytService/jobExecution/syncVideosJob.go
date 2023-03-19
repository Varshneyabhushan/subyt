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

		//fetch videos from YouTube
		response, toDelay, err := fetchVideoResponse(*fetcher, delayTracker, checkPoint.NextPageToken)
		if err != nil || toDelay {
			return delayTracker.Delay(), err
		}

		newCheckPoint, totalValidVideos := NewCheckpoint(*checkPoint, response)

		//add videos to videosService
		delay, err := saveVideos(service, delayTracker, response.Videos[:totalValidVideos])
		if delay || err != nil {
			return delayTracker.Delay(), err
		}

		delayTracker.Reset()

		//when checkpoint is not saved, oldCheckpoint is reused
		if err = newCheckPoint.Save(); err != nil {
			log.Println("error while updating checkpoint", err)
			return delayTracker.Delay(), nil
		}
		*checkPoint = newCheckPoint

		//when pagination continues, sync shouldn't cool down
		if len(checkPoint.NextPageToken) != 0 {
			return delayTracker.Delay(), nil
		}

		log.Println("sync cool down")
		return syncCoolDown, nil
	}
}

func NewCheckpoint(checkPoint storage.CheckPoint, response videofetcher.VideosResponse) (
	storage.CheckPoint, int) {

	if len(response.Videos) == 0 {
		return checkPoint, 0
	}

	totalValidVideos := videosservice.GetValidVideos(response.Videos, checkPoint.VideoLimit)
	if totalValidVideos != 0 && checkPoint.NextVideoLimit.IsEmpty() {
		checkPoint.NextVideoLimit = response.Videos[0].VideoComparor
	}

	if totalValidVideos == len(response.Videos) {
		checkPoint.NextPageToken = response.NextPageToken
		return checkPoint, totalValidVideos
	}

	//don't go to next page
	checkPoint.NextPageToken = ""
	if totalValidVideos == 0 {
		return checkPoint, 0
	}

	checkPoint.VideoLimit = checkPoint.NextVideoLimit
	checkPoint.NextVideoLimit = videosservice.VideoComparor{}
	return checkPoint, totalValidVideos
}
