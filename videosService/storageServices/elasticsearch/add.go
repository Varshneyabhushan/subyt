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

func newBulkIndexer(client *elasticsearch.Client, indexName string) (esutil.BulkIndexer,
	func() error) {
	bulkIndexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		NumWorkers: 1,
		Client:     client,
		Index:      indexName,
	})

	closeIndexer := func() error {
		if err := bulkIndexer.Close(context.Background()); err != nil {
			return errors.New("error while closing bi : " + err.Error())
		}

		return nil
	}

	return bulkIndexer, closeIndexer
}

func addDocToBulkIndexer[T Document](indexer esutil.BulkIndexer, doc T) error {
	docBytes, _ := json.Marshal(doc)
	return indexer.Add(context.Background(), esutil.BulkIndexerItem{
		Action:     "index",
		DocumentID: doc.GetId(),
		Body:       bytes.NewReader(docBytes),
	})
}

func (service *Index[T]) Add(documents []T) error {
	bulkIndexer, closeIndex := newBulkIndexer(service.esClient, service.indexName)
	for _, doc := range documents {
		if err := addDocToBulkIndexer(bulkIndexer, doc); err != nil {
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
