package env

import (
	"encoding/json"
	"errors"
	"os"
)

type ServerConfig struct {
	Port int
}

type Config struct {
	ServerConfig ServerConfig
}

func GetConfig() (result Config, err error) {
	envBytes, err := os.ReadFile("env.json")
	if err != nil {
		return result, errors.New("error while reading file : " + err.Error())
	}

	return result, json.Unmarshal(envBytes, &result)
}
