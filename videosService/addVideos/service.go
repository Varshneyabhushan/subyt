package addVideos

import (
	"videosservice/repository"
)

/**
as this service is being used by POST request,
this is bound to be used with repository that should
store and get back the repository.
*/

type Service interface {
	Add(videos []repository.Video) (int, error)
}
