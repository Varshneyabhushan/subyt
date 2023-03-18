package limitidentifier

import (
	"encoding/json"
	"errors"
	"os"
)

type LimitIdentifierStorage struct {
	LimitIdentifier
	filePath string
}

func New(filePath string) (LimitIdentifierStorage, error) {
	result := LimitIdentifierStorage {}
	result.filePath = filePath
	
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return result, errors.New("error while reading file : " + err.Error())
	}

	err = json.Unmarshal(bytes, &result.LimitIdentifier)
	return result, err
}

func (t LimitIdentifierStorage) Save() error {
	bytes, err := json.Marshal(t.LimitIdentifier)
	if err != nil {
		return errors.New("error while converting to bytes : " + err.Error())
	}

	return os.WriteFile(t.filePath, bytes, os.ModeAppend)
}