package main

import (
	"errors"
	"time"
	"ytservice/storage"
	"ytservice/videofetcher"
	"ytservice/videosservice"
)

func AddVideos(
	checkPoint storage.CheckPoint,
	fetcher videofetcher.VideoFetcher,
	videosService videosservice.VideosService,
	coolDown int,
) error {

	if checkPoint.LimitVideo.IsEmpty() {
		return errors.New("limit on video is empty")
	}

	var nextLimitVideo videosservice.VideoComparor
	fetcher.NextPageToken = checkPoint.NextPageToken

	for fetcher.HasNext() {
		timeout := time.After(time.Duration(coolDown) * time.Second)
		videos, err := fetcher.GetNext()
		if err != nil {
			return errors.New("error while getting videos : " + err.Error())
		}

		validVideosCount := videosservice.GetValidVideos(videos, checkPoint.LimitVideo)
		if validVideosCount != 0 && nextLimitVideo.IsEmpty() {
			nextLimitVideo = videos[0].VideoComparor
		}

		//if this is the limit, then limitVideo has to be updated and call should be ended
		isLimiting := validVideosCount < len(videos)
		if isLimiting {
			checkPoint.LimitVideo = nextLimitVideo
			nextLimitVideo = videosservice.VideoComparor{}
		}

		//add videos to videosService
		if err = videosService.AddVideos(videos[:validVideosCount]); err != nil {
			return errors.New("error while adding videos : " + err.Error())
		}

		//save checkPoint for every fetch: nextPageToken, limitVideo are saved
		checkPoint.NextPageToken = fetcher.NextPageToken
		if err := checkPoint.Save(); err != nil {
			return errors.New("error while storing limits : " + err.Error())
		}

		<-timeout
		if isLimiting {
			return nil
		}
	}

	return nil
}
