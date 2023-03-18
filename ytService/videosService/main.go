package videosservice

import (
	"fmt"
)

type VideosService struct {
	apiUrl string
}

func (service VideosService) AddVideos(videos []Video) error {
	if len(videos) == 0 {
		return nil
	}

	//send these videos to videosService
	for _, video := range videos {
		fmt.Println(video.Id, video.Title, video.PublishedAt)
	}
	return nil
}

func NewVideoService(apiUrl string) VideosService {
	return VideosService{ apiUrl: apiUrl }
}