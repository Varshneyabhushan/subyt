package videosservice

import (
	"fmt"
	"ytservice/videofetcher"
)

type VideosService struct {
	apiUrl string
}

func (service VideosService) AddVideos(videos []videofetcher.Video) error {
	if len(videos) == 0 {
		return nil
	}

	//send these videos to videosService
	fmt.Println(videos)
	return nil
}

func (service VideosService) GetTopVideoId() (string, error) {
	//get from videosService
	return "", nil
}

func NewVideoService(apiUrl string) VideosService {
	return VideosService{ apiUrl: apiUrl }
}