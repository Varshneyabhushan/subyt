package storage

import (
	"github.com/elastic/go-elasticsearch/v8"
	"net/http"
	"videosservice/env"
)

func GetESClient(config env.ElasticSearchConfig) (*elasticsearch.Client, error) {
	esConfig := elasticsearch.Config{
		Addresses: []string{
			config.Uri,
		},
		Transport: &http.Transport{
			ResponseHeaderTimeout: config.ConnectionTimeout,
		},
	}

	client, err := elasticsearch.NewClient(esConfig)
	return client, err
}
