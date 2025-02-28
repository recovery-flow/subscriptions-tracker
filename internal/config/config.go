package config

import (
	"os"
	"time"

	_ "github.com/lib/pq" // postgres driver don`t delete
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	BasePath string `mapstructure:"base_path"`
	TestMode bool   `mapstructure:"test_mode"`
	Log      struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"log"`
}

type JWTConfig struct {
	AccessToken struct {
		SecretKey     string        `mapstructure:"secret_key"`
		TokenLifetime time.Duration `mapstructure:"token_lifetime"`
	} `mapstructure:"access_token"`
}

type DatabaseConfig struct {
	SQL struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"sql"`

	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
		Lifetime int    `mapstructure:"lifetime"`
	} `mapstructure:"redis"`
}

type RabbitConfig struct {
	URL      string `mapstructure:"url"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Exchange string `mapstructure:"exchange"`
}

type SwaggerConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	URL     string `mapstructure:"url"`
	Port    string `mapstructure:"port"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Rabbit   RabbitConfig   `mapstructure:"rabbit"`
	Database DatabaseConfig `mapstructure:"database"`
	Swagger  SwaggerConfig  `mapstructure:"swagger"`
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("KV_VIPER_FILE")
	if configPath == "" {
		return nil, errors.New("KV_VIPER_FILE env var is not set")
	}
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Errorf("error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Errorf("error unmarshalling config: %s", err)
	}

	return &config, nil
}
