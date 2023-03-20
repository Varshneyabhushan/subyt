package env

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type YoutubeSearchConfig struct {
	DevKey     string
	Query      string
	MaxResults int
}

type SchedulerConfig struct {
	SyncCoolDown    time.Duration //in seconds
	RequestCoolDown time.Duration //in seconds
}

type Config struct {
	YoutubeSearch   YoutubeSearchConfig
	Scheduler       SchedulerConfig
	CheckPointPath  string
	VideoServiceURL string
}

func GetConfigFromFile(filePath string) (result Config, err error) {
	envBytes, err := os.ReadFile(filePath)
	if err != nil {
		return result, errors.New("error while reading file : " + err.Error())
	}

	err = json.Unmarshal(envBytes, &result)
	if err != nil {
		return result, err
	}

	//convert to seconds
	result.Scheduler.SyncCoolDown *= time.Second
	result.Scheduler.RequestCoolDown *= time.Second
	return result, nil
}
