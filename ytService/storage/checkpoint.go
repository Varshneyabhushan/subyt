package storage

import (
	"encoding/json"
	"errors"
	"os"
	"ytservice/videosservice"
)

//this is about storing the temporary data to a file
//so that the app is resilient
//when app is abruptly shutdown, this helps restoring to previous check point


//struct to be stored in json
type CheckPoint struct {
	filePath string
	NextPageToken string 
	LimitVideo videosservice.VideoComparor
}

func LoadCheckPoint(filePath string) (CheckPoint, error) {
	newCheckPoint := CheckPoint { filePath: filePath }
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return newCheckPoint, errors.New("error while reading file : " + err.Error())
	}

	err = json.Unmarshal(bytes, &newCheckPoint)
	return newCheckPoint, err
}

func (p CheckPoint) Save() error {
	bytes, err := json.Marshal(p)
	if err != nil {
		return errors.New("error while converting to bytes : " + err.Error())
	}

	return os.WriteFile(p.filePath, bytes, os.ModeAppend)
}