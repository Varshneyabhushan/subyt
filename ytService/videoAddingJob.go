package main

import (
	"errors"
	"time"
	"ytservice/limitidentifier"
	"ytservice/videofetcher"
	"ytservice/videosservice"
)

type VideoIdProvider videosservice.Video

func (v VideoIdProvider) GetId() string {
	return v.Id
}

func VideoIdProviders(videos []videosservice.Video) []limitidentifier.IdProvider {
	var result []limitidentifier.IdProvider
	for _, video := range videos {
		result = append(result, VideoIdProvider(video))
	}

	return result
}

func AddVideos(
	limitIdentifier *limitidentifier.LimitIdentifierStorage,
	fetcher *videofetcher.VideoFetcher,
	videosService videosservice.VideosService,
	coolDown int,
 ) error {
	if(len(limitIdentifier.Limit) == 0) {
		return errors.New("limit is zero in limit identifier")
	}

	for fetcher.HasNext() {
		timeout := time.After(time.Duration(coolDown)* time.Second)
		videos, err := fetcher.GetNext()
		if err != nil {
			return errors.New("error while getting videos : " + err.Error())
		}

		validVideosCount := limitIdentifier.Scrutinise(VideoIdProviders(videos))
		isLimiting := validVideosCount < len(videos)
		if isLimiting {
			limitIdentifier.AdvanceLimit()
			err := limitIdentifier.Save()
			if err != nil {
				return errors.New("error while storing limits : " + err.Error())
			}
		}

		if err = videosService.AddVideos(videos[:validVideosCount]); err != nil {
			return errors.New("error while adding videos : " + err.Error())
		}

		<- timeout
		if isLimiting {
			return nil
		}
	}

	return nil
}