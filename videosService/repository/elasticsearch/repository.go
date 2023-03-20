package elasticsearch

import "github.com/elastic/go-elasticsearch/v8"

type Repository struct {
	esClient *elasticsearch.Client
}

func NewRepository(esClient *elasticsearch.Client) *Repository {
	return &Repository{esClient: esClient}
}

func (repo *Repository) Search(_ string) ([]Video, error) {
	return nil, nil
}
