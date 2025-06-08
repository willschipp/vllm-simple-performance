package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Endpoint struct {
		Url string `mapstructure:"url"`
		Prompt string `mapstructure:"prompt"`
		Model string `mapstructure:"model"`
	}
	Metrics struct {
		Url string `mapstructure:"url"`
		Interval int `mapstructure:"interval"`
		Output string `mapstructure:"output"`
	}
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./") //TODO set passable

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}