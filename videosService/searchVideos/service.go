package searchVideos

import "videosservice/videos"

type Service interface {
	Search(term string) ([]videos.Video, error)
}
