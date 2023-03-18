package main

import (
	"fmt"
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
 ) {
	
	for fetcher.HasNext() {
		videos, err := fetcher.GetNext()
		if err != nil {
			fmt.Println("error while getting videos : " + err.Error())
			break
		}

		validVideosCount := limitIdentifier.Scrutinise(VideoIdProviders(videos))
		isLimiting := validVideosCount < len(videos)
		if isLimiting {
			limitIdentifier.AdvanceLimit()
			err := limitIdentifier.Save()
			if err != nil {
				fmt.Println("error while storing limits : " + err.Error())
				return
			}
		}

		if err = videosService.AddVideos(videos[:validVideosCount]); err != nil {
			fmt.Println("error while adding videos : " + err.Error())
			return
		}

		if isLimiting {
			break;
		}
	}
}