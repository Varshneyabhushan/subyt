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
	fmt.Println(videos)
	return nil
}

func NewVideoService(apiUrl string) VideosService {
	return VideosService{ apiUrl: apiUrl }
}