package servicecheck

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Instance struct {
	StatusUrl string `yaml:"statusUrl"`
}

type Instances []Instance

type HttpService struct {
	Name      string    `yaml:"name"`
	Instances Instances `yaml:"instances"`
}

type HttpServices []HttpService

type Config struct {
	HttpServices HttpServices `yaml:"httpServices"`
}

func NewConfig(configFilePath string) (*Config, error) {
	configFileData, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error occurred while reading config file at path %s: %v", configFilePath, err)
	}

	var config Config

	err = yaml.Unmarshal(configFileData, &config)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing yaml config file at path %s: %v", configFilePath, err)
	}

	return &config, nil
}
