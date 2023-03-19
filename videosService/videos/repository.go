package videos

type Repository interface {
	Add(videos []Video) (int, error)
	Get(skip, limit int) ([]Video, error)
	GetByIds(ids []string) ([]Video, error)
}
