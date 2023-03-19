package searchVideos

import (
	"videosservice/repository"
)

type Service interface {
	Search(term string) ([]repository.Video, error)
}
