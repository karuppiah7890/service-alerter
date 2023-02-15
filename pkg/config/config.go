package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// All configuration is through environment variables

const CONFIG_FILE_PATH_ENV_VAR = "CONFIG_FILE_PATH"
const DEFAULT_CONFIG_FILE_PATH = "service-alerter.yaml"
const ENVIRONMENT_NAME_ENV_VAR = "ENVIRONMENT_NAME"
const DEFAULT_ENVIRONMENT_NAME = "Production"
const SLACK_TOKEN_ENV_VAR = "SLACK_TOKEN"
const SLACK_CHANNEL_ENV_VAR = "SLACK_CHANNEL"

type Config struct {
	configFilePath  string
	environmentName string
	slackToken      string
	slackChannel    string
}

func NewConfigFromEnvVars() (*Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting nginx port: %v", err)
	}

	environmentName := getEnvironmentName()

	slackToken, err := getSlackToken()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting slack token: %v", err)
	}

	slackChannel, err := getSlackChannel()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting slack channel: %v", err)
	}

	return &Config{
		configFilePath:  configFilePath,
		environmentName: environmentName,
		slackToken:      slackToken,
		slackChannel:    slackChannel,
	}, nil
}

// Get nginx port number. Default is "80"
func getConfigFilePath() (string, error) {
	configFilePath, ok := os.LookupEnv(CONFIG_FILE_PATH_ENV_VAR)
	if !ok {
		configFilePath = DEFAULT_CONFIG_FILE_PATH
	}

	_, err := os.Stat(configFilePath)

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return "", fmt.Errorf("config file does not exist at path %s", configFilePath)
		}

		return "", fmt.Errorf("could not find file info of the config file at path %s: %v", configFilePath, err)
	}

	return configFilePath, nil
}

// Get optional environment name for the environment where
// the services are running. Default is "production". This name will
// be used in the alert messages
func getEnvironmentName() string {
	environmentName, ok := os.LookupEnv(ENVIRONMENT_NAME_ENV_VAR)
	if !ok {
		return DEFAULT_ENVIRONMENT_NAME
	}

	return environmentName
}

func getSlackToken() (string, error) {
	slackToken, ok := os.LookupEnv(SLACK_TOKEN_ENV_VAR)
	if !ok {
		return "", fmt.Errorf("%s environment variable is not defined and is required. Please define it", SLACK_TOKEN_ENV_VAR)
	}
	return slackToken, nil
}

func getSlackChannel() (string, error) {
	slackChannel, ok := os.LookupEnv(SLACK_CHANNEL_ENV_VAR)
	if !ok {
		return "", fmt.Errorf("%s environment variable is not defined and is required. Please define it", SLACK_CHANNEL_ENV_VAR)
	}
	return slackChannel, nil
}

func (c *Config) GetConfigFilePath() string {
	return c.configFilePath
}

func (c *Config) GetEnvironmentName() string {
	return c.environmentName
}

func (c *Config) GetSlackToken() string {
	return c.slackToken
}

func (c *Config) GetSlackChanel() string {
	return c.slackChannel
}
