package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"strconv"
)

type Repository struct {
	esClient *elasticsearch.Client
}

func NewRepository(esClient *elasticsearch.Client) *Repository {
	return &Repository{esClient: esClient}
}

func (repo *Repository) Add(esVideos []Video) error {

	bulkIndexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		NumWorkers: 1,
		Client:     repo.esClient,
		Index:      "videos",
	})

	for _, esVideo := range esVideos {
		esVideoBytes, _ := json.Marshal(esVideo)
		err := bulkIndexer.Add(context.Background(), esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: esVideo.Id,
			Body:       bytes.NewReader(esVideoBytes),
		})

		if err != nil {
			return errors.New("unexpected error : " + err.Error())
		}
	}

	if err := bulkIndexer.Close(context.Background()); err != nil {
		return errors.New("error while closing bi : " + err.Error())
	}

	failedCount := bulkIndexer.Stats().NumFailed
	if failedCount > 0 {
		return errors.New("failed bulk indexing : " + strconv.Itoa(int(failedCount)))
	}

	return nil
}

func (repo *Repository) Search(_ string) ([]Video, error) {
	return nil, nil
}
