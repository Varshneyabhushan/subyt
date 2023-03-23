package elasticsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
)

type EsHit[T interface{}] struct {
	Id     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source T       `json:"_source"`
}

type HitsResult[T interface{}] struct {
	Total struct{ Value int64 }
	Hits  []EsHit[T]
}

type queryResponse[T interface{}] struct {
	Took       int
	TimedOut   bool          `json:"timed_out"`
	HitsResult HitsResult[T] `json:"hits"`
}

func GetDocsFromResponse[T interface{}](response *esapi.Response) ([]T, int64, error) {
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, errors.New("error while reading response : " + err.Error())
	}

	var esResult queryResponse[T]
	err = json.Unmarshal(responseBytes, &esResult)
	if err != nil {
		return nil, 0, errors.New("error while reading result : " + err.Error())
	}

	var result []T
	for _, hit := range esResult.HitsResult.Hits {
		result = append(result, hit.Source)
	}

	return result, esResult.HitsResult.Total.Value, nil
}

func buildQuery(term string, skip int, limit int) any {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Title": term,
			},
		},
		"from": skip,
		"size": limit,
	}
}

func (service *Index[T]) Search(term string, skip, limit int) ([]T, int64, error) {
	var buf bytes.Buffer
	query := buildQuery(term, skip, limit)
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, 0, errors.New("Error encoding query: " + err.Error())
	}

	search := service.esClient.Search
	response, err := search(search.WithIndex(service.indexName), search.WithBody(&buf))
	if err != nil {
		return nil, 0, err
	}

	return GetDocsFromResponse[T](response)
}
