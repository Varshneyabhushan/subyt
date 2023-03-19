package videos

type Repository interface {
	Add(videos []Video) (int, error)
	Get(skip, limit int) ([]Video, error)
	Search(term string) ([]Video, error)
}
