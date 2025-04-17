package config

import (
	"fmt"
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var mods = []string{"dev", "prod"}

type PostgresConfig struct {
	Host               string `mapstructure:"host" validate:"required,hostname|ip"`
	Port               int    `mapstructure:"port" validate:"required,min=1,max=65535"`
	User               string `mapstructure:"user" validate:"required"`
	Password           string `mapstructure:"password" validate:"required"`
	DBName             string `mapstructure:"dbname" validate:"required"`
	SSLMode            string `mapstructure:"sslmode" validate:"required,oneof=disable require verify-ca verify-full"`
	MaxOpenConns       int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns       int    `mapstructure:"max_idle_conns" validate:"required,min=0"`
	ConnMaxLifetimeMin int    `mapstructure:"conn_max_lifetime_minutes" validate:"required,min=1"`
	ConnMaxIdleTimeMin int    `mapstructure:"conn_max_idle_time_minutes" validate:"required,min=0"`
}

type TokenConfig struct {
	PrivateKeyPath        string `mapstructure:"private_key_path" validate:"required"`
	PublicKeyPath         string `mapstructure:"public_key_path" validate:"required"`
	AccessTokenExpiresIn  int    `mapstructure:"access_token_expires_in" validate:"required"`  // minutes
	RefreshTokenExpiresIn int    `mapstructure:"refresh_token_expires_in" validate:"required"` // days
}

type Config struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Token    TokenConfig    `mapstructure:"token"`
}

func Load(mod string) (Config, error) {
	var config Config

	if !slices.Contains(mods, mod) {
		return config, fmt.Errorf("unsupported mode: %s", mod)
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(mod)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, fmt.Errorf("failed to validate config file: %w", err)
	}

	return config, nil
}
