package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type conf struct {
	WeatherAPIKey string `mapstructure:"WEATHERAPI_API_KEY"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg conf

	viper.AutomaticEnv()

	if apiKey := os.Getenv("WEATHERAPI_API_KEY"); apiKey != "" {
		fmt.Println("Using API key from environment variable")
		cfg.WeatherAPIKey = apiKey
	} else {
		viper.SetConfigName("app_config")
		viper.SetConfigType("env")
		viper.AddConfigPath(path)
		viper.SetConfigFile(path + "/.env")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return nil, err
			}
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}
