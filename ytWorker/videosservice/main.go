package videosservice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type VideosService struct {
	apiUrl string
}

func (service VideosService) AddVideos(videos []Video) error {
	if len(videos) == 0 {
		return nil
	}

	payload, err := json.Marshal(videos)
	if err != nil {
		return errors.New("error while marshalling : " + err.Error())
	}

	url := fmt.Sprintf("%s/", service.apiUrl)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return nil
	}

	return HttpError(resp.StatusCode)
}

func NewVideoService(apiUrl string) VideosService {
	return VideosService{apiUrl: apiUrl}
}
