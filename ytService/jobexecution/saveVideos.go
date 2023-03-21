package jobexecution

import (
	"errors"
	"log"
	"syscall"
	"ytservice/videosservice"
)

func saveVideos(service videosservice.VideosService, _ *DelayTracker,
	videos []videosservice.Video) (bool, error) {
	if err := service.AddVideos(videos); err != nil {
		if errors.Is(err, syscall.ECONNREFUSED) ||
			errors.Is(err, syscall.ECONNABORTED) ||
			errors.Is(err, syscall.ECONNRESET) {
			return false, errors.New("videosService not reachable : " + err.Error())
		}

		if errors.Is(err, videosservice.HttpError(404)) {
			return false, errors.New("api not found")
		}

		if errors.Is(err, videosservice.HttpError(400)) {
			log.Println("bad request to videosService : " + err.Error())
			return false, nil
		}

		log.Println("error while adding videos : " + err.Error())
		return true, nil
	}

	return false, nil
}
