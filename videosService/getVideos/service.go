package getVideos

import (
	"videosservice/repository"
)

type Service interface {
	Get(skip, limit int64) ([]repository.Video, error)
}
