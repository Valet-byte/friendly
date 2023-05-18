package config

import (
	"encoding/csv"
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type PathsConfig struct {
	Path       string
	Method     string
	IsOpen     bool
	AuthMethod string
}

type RestPathsConfig struct {
	Data *[]PathsConfig
}

func NewRestPathsConfig(csvFilePath string) (*RestPathsConfig, error) {
	var data []PathsConfig

	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		config := PathsConfig{
			Path:       record[0],
			Method:     record[1],
			IsOpen:     record[2] == "true",
			AuthMethod: record[3],
		}
		data = append(data, config)
	}

	return &RestPathsConfig{Data: &data}, nil
}

func (rc *RestPathsConfig) FindConfig(path, method string) (*PathsConfig, error) {

	for _, config := range *rc.Data {
		if matchPath(config.Path, path) && strings.Compare(config.Method, method) == 0 {
			logrus.Println(path, " ", method)
			if config.Method == method {
				return &config, nil
			}
		}
	}

	return nil, errors.New("not find path Data")
}

func matchPath(pattern, path string) bool {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i := range patternParts {
		if !strings.HasPrefix(patternParts[i], ":") {
			if patternParts[i] != pathParts[i] {
				return false
			}
		}
	}

	return true
}
