package elasticsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
)

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
	Took       int
	TimedOut   bool       `json:"timed_out"`
	HitsResult HitsResult `json:"hits"`
}

func GetVideosFromResponse(response *esapi.Response) ([]Video, error) {
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("error while reading response : " + err.Error())
	}

	var esResult queryResponse
	err = json.Unmarshal(responseBytes, &esResult)
	if err != nil {
		return nil, errors.New("error while reading result : " + err.Error())
	}

	var result []Video
	for _, hit := range esResult.HitsResult.Hits {
		result = append(result, hit.Source)
	}

	return result, nil
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

func (repo *Repository) Search(term string, skip, limit int) ([]Video, error) {
	var buf bytes.Buffer
	query := buildQuery(term, skip, limit)
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errors.New("Error encoding query: " + err.Error())
	}

	search := repo.esClient.Search
	response, err := search(search.WithIndex("videos"), search.WithBody(&buf))
	if err != nil {
		return nil, err
	}

	return GetVideosFromResponse(response)
}
