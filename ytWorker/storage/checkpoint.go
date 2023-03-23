package storage

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"time"
	"ytworker/videosservice"
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

// limited to 120 days by default
var defaultLimitVideo = time.Now().Add(-time.Hour * 24 * 120)

func LoadCheckPoint(filePath string) *CheckPoint {
	newCheckPoint := &CheckPoint{
		filePath: filePath,
		VideoLimit: videosservice.VideoComparor{
			PublishedAt: defaultLimitVideo,
		},
	}

	file, err := os.Open(filePath)
	//when doesn't exist, create the file and return the newCheckpoint
	//when this is saved, faulty file gets overwritten with this newCheckpoint
	if errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(filePath); err != nil {
			log.Println("error while creating file for checkPoint : ", err)
		}
		return newCheckPoint
	}

	//file exists, but it can be a directory also
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("error while reading file stats : ", err)
		return newCheckPoint
	}

	//when file is a directory
	if fileInfo.IsDir() {
		_ = os.RemoveAll(filePath)
		if _, err = os.Create(filePath); err != nil {
			log.Println("error while creating file for checkPoint : ", err)
		}
		return newCheckPoint
	}

	//even when unmarshalling is succeeded or not, we will return the newCheckpoint
	bytes, _ := io.ReadAll(file)
	_ = json.Unmarshal(bytes, newCheckPoint)
	return newCheckPoint
}

func (p *CheckPoint) Save() error {
	bytes, err := json.Marshal(p)
	if err != nil {
		return errors.New("error while converting to bytes : " + err.Error())
	}

	return os.WriteFile(p.filePath, bytes, os.ModeAppend)
}
