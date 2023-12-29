package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB DB `mapstructure:"db"`
}

type DB struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	URL      string `mapstructure:"url"`
	Options  string `mapstructure:"options"`
}

func NewConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file. err = %v", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal config file. err = %v", err)
	}

	return &config, nil
}
