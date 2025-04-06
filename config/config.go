package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Test string `mapstructure:"TEST"`
}

func (c Config) Validate() error {
	return nil
}

func Load() (Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.UnmarshalExact(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	if err := config.Validate(); err != nil {
		return config, fmt.Errorf("failed to validate config: %w", err)
	}

	return config, nil
}
