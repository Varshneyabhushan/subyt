package env

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type ServerConfig struct {
	Port int
}

type MongoConfig struct {
	Uri          string
	DatabaseName string
}

type ElasticSearchConfig struct {
	Uri string
}

type Config struct {
	ServerConfig        `json:"serverConfig"`
	MongoConfig         `json:"mongoConfig"`
	ElasticSearchConfig `json:"elasticSearchConfig"`
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

	return result, json.Unmarshal(envBytes, &result)
}

func GetConfigFromEnv() (result Config, err error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return result, errors.New("error while reading PORT from env")
	}

	return Config{
		ServerConfig: ServerConfig{Port: port},
		MongoConfig: MongoConfig{
			Uri:          os.Getenv("MONGO_URI"),
			DatabaseName: os.Getenv("MONGO_DATABASE"),
		},
		ElasticSearchConfig: ElasticSearchConfig{
			Uri: os.Getenv("ELASTICSEARCH_URI"),
		},
	}, nil
}
