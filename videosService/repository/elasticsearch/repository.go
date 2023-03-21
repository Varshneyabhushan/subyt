package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"io"
	"log"
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

type EsHit struct {
	Id     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Video   `json:"_source"`
}

type HitsResult struct {
	Total struct{ Value int }
	Hits  []EsHit
}

type queryResponse struct {
	Took     int
	TimedOut bool `json:"timed_out"`
	Hits     HitsResult
}

func GetResultFromResponse(response queryResponse) []Video {
	var result []Video
	for _, hit := range response.Hits.Hits {
		result = append(result, hit.Source)
	}

	return result
}

func (repo *Repository) Search(term string, skip, limit int) ([]Video, error) {
	search := repo.esClient.Search

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Title": term,
			},
		},
		"from": skip,
		"size": limit,
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errors.New("Error encoding query: " + err.Error())
	}

	response, err := search(
		search.WithIndex("videos"),
		search.WithBody(&buf),
		search.WithPretty(),
	)

	if err != nil {
		return nil, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("error while reading response : " + err.Error())
	}

	log.Println(string(responseBytes))

	var esResult queryResponse
	err = json.Unmarshal(responseBytes, &esResult)
	if err != nil {
		return nil, errors.New("error while reading result : " + err.Error())
	}

	return GetResultFromResponse(esResult), nil
}
