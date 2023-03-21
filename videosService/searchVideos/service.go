package searchVideos

import (
	"videosservice/repository"
)

type Service interface {
	Search(term string, skip, limit int) ([]repository.Video, error)
}
