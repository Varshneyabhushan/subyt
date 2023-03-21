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

func newBulkIndexer(client *elasticsearch.Client) (esutil.BulkIndexer, func() error) {
	bulkIndexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		NumWorkers: 1,
		Client:     client,
		Index:      "videos",
	})

	closeIndexer := func() error {
		if err := bulkIndexer.Close(context.Background()); err != nil {
			return errors.New("error while closing bi : " + err.Error())
		}

		return nil
	}

	return bulkIndexer, closeIndexer
}

func addVideoToBulkIndexer(indexer esutil.BulkIndexer, esVideo Video) error {
	esVideoBytes, _ := json.Marshal(esVideo)
	return indexer.Add(context.Background(), esutil.BulkIndexerItem{
		Action:     "index",
		DocumentID: esVideo.Id,
		Body:       bytes.NewReader(esVideoBytes),
	})
}

func (repo *Repository) Add(esVideos []Video) error {
	bulkIndexer, closeIndex := newBulkIndexer(repo.esClient)
	for _, esVideo := range esVideos {
		if err := addVideoToBulkIndexer(bulkIndexer, esVideo); err != nil {
			return errors.New("unexpected error : " + err.Error())
		}
	}

	if err := closeIndex(); err != nil {
		return err
	}

	failedCount := bulkIndexer.Stats().NumFailed
	if failedCount > 0 {
		return errors.New("failed bulk indexing : " + strconv.Itoa(int(failedCount)))
	}

	return nil
}
