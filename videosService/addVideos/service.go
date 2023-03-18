package addVideos

type VideosAddingService interface {
	AddVideos(videos []Video) (int, error)
}
