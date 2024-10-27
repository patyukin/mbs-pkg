package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

func LoadConfig[T any](configFilePathEnv string, config *T) error {
	yamlConfigFilePath := os.Getenv(configFilePathEnv)
	if yamlConfigFilePath == "" {
		return fmt.Errorf("environment variable %s is not set", configFilePathEnv)
	}

	f, err := os.Open(yamlConfigFilePath)
	if err != nil {
		return fmt.Errorf("unable to open config file: %w", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("unable to close config file: %v\n", err)
		}
	}()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(config); err != nil {
		return fmt.Errorf("unable to decode config file: %w", err)
	}

	validate := validator.New()
	if err = validate.Struct(config); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}

	return nil
}
