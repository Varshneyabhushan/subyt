package storage

import (
	"github.com/elastic/go-elasticsearch/v8"
	"videosservice/env"
)

func GetESClient(config env.ElasticSearchConfig) (*elasticsearch.Client, error) {
	esConfig := elasticsearch.Config{
		Addresses: []string{
			config.Uri,
		},
	}

	client, err := elasticsearch.NewClient(esConfig)
	return client, err
}
