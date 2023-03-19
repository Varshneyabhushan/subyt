package storage

import (
	"encoding/json"
	"errors"
	"os"
	"time"
	"ytservice/videosservice"
)

//this is about storing the temporary data to a file
//so that the app is resilient
//when app is abruptly shutdown, this helps restoring to previous check point

// CheckPoint struct to be stored in json
type CheckPoint struct {
	filePath       string
	VideoLimit     videosservice.VideoComparor
	NextVideoLimit videosservice.VideoComparor
	NextPageToken  string
}

var defaultLimitVideo, _ = time.Parse(time.RFC3339, "2023-03-10T00:00:00Z")

func LoadCheckPoint(filePath string) (*CheckPoint, error) {
	newCheckPoint := &CheckPoint{
		filePath: filePath,
		VideoLimit: videosservice.VideoComparor{
			PublishedAt: defaultLimitVideo,
		},
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return newCheckPoint, err
	}

	err = json.Unmarshal(bytes, newCheckPoint)
	return newCheckPoint, err
}

func (p *CheckPoint) Save() error {
	bytes, err := json.Marshal(p)
	if err != nil {
		return errors.New("error while converting to bytes : " + err.Error())
	}

	return os.WriteFile(p.filePath, bytes, os.ModeAppend)
}
