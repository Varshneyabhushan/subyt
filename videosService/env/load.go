package env

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	Port int
}

type MongoConfig struct {
	Uri               string
	DatabaseName      string
	ConnectionTimeout time.Duration //in seconds
}

type Config struct {
	ServerConfig ServerConfig
	MongoConfig  MongoConfig
}

const envFilePathKey = "ENV_FILEPATH"

// LoadEnv load env from file if envVariable is present
// else try to read from env variables themselves
func LoadEnv() (Config, error) {
	filePath := os.Getenv(envFilePathKey)
	if len(filePath) == 0 {
		return GetConfigFromEnv()
	}

	return GetConfigFromFile(filePath)
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

	result.MongoConfig.ConnectionTimeout *= time.Second
	return result, json.Unmarshal(envBytes, &result)
}

func GetConfigFromEnv() (result Config, err error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return result, errors.New("error while reading PORT from env")
	}

	result.ServerConfig = ServerConfig{Port: port}

	timeout, _ := strconv.Atoi(os.Getenv("MONGO_CONNECTION_TIMEOUT"))
	result.MongoConfig = MongoConfig{
		Uri:               os.Getenv("MONGO_URI"),
		DatabaseName:      os.Getenv("MONGO_DATABASE"),
		ConnectionTimeout: time.Duration(timeout) * time.Second,
	}

	return result, nil
}
