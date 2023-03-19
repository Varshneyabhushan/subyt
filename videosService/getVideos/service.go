package getVideos

import "videosservice/videos"

type Service interface {
	Get(skip, limit int) ([]videos.Video, error)
}
