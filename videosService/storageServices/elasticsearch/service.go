package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"videosservice/env"
)

type Document interface {
	GetId() string
}

type Index[T Document] struct {
	esClient  *elasticsearch.Client
	indexName string
}

type Service *elasticsearch.Client

func NewService(config env.ElasticSearchConfig) (Service, error) {
	esConfig := elasticsearch.Config{
		Addresses: []string{
			config.Uri,
		},
	}

	return elasticsearch.NewClient(esConfig)
}

func GetIndex[T Document](service Service, indexName string) Index[T] {
	return Index[T]{
		indexName: indexName,
		esClient:  service,
	}
}
