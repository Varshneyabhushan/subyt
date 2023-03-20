package elasticsearch

type Repository struct{}

func (repo *Repository) Search(_ string) ([]Video, error) {
	return nil, nil
}
