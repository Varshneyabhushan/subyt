package addVideos

import "videosservice/videos"

/**
as this service is being used by POST request,
this is bound to be used with repository that should
store and get back the videos.
*/

type Service interface {
	Add(videos []videos.Video) (int, error)
}
