package env

import (
	"encoding/json"
	"errors"
	"os"
)

type YoutubeSearchConfig struct {
	DevKey string
	Query string
	MaxResults int
	RequestCoolDown int //in seconds
}

type SchedulerConfig struct {
	Period int
}

type Config struct {
	YoutubeSearch YoutubeSearchConfig
	Scheduler SchedulerConfig
	CheckPointPath string
	VideoServiceURL string
}

func GetConfig() (result Config, err error) {
	envBytes, err := os.ReadFile("env.json")
	if err != nil {
		return result, errors.New("error while reading file : " + err.Error())
	}

	return result, json.Unmarshal(envBytes, &result)
}