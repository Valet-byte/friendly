package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type FriendlyConfig struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Read   time.Duration `yaml:"read"`
			Server time.Duration `yaml:"server"`
			Write  time.Duration `yaml:"write"`
			Idle   time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
		Cache struct {
			Host     string        `yaml:"host"`
			Password string        `yaml:"password"`
			Duration time.Duration `yaml:"duration"`
		} `yaml:"cache"`
		Jwt struct {
			SecretKey    string        `yaml:"secretKey"`
			ValidTime    time.Duration `yaml:"validTime"`
			RecreateTime time.Duration `yaml:"recreateTime"`
		} `yaml:"jwt"`
	} `yaml:"server"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}

func LoadConfig(configPath string) (*FriendlyConfig, error) {
	config := &FriendlyConfig{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&config)

	return config, err
}
