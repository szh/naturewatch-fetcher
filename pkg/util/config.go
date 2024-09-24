package util

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	NatureWatchURL       string `mapstructure:"NATUREWATCH_URL"`
	FetchIntervalSeconds int    `mapstructure:"FETCH_INTERVAL_SECONDS"`
	OutputPath           string `mapstructure:"OUTPUT_PATH"`
}

var Config ConfigStruct

func LoadConfig(path string) (config ConfigStruct, err error) {
	// config file
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// environment variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	Config = config
	return
}

func ValidateConfig() error {
	// Ensure NatureWatchURL is a valid URL
	_, err := url.Parse(Config.NatureWatchURL)
	if err != nil {
		return fmt.Errorf("Invalid NatureWatchURL: %v", err)
	}

	// Ensure FetchIntervalSeconds is a positive integer
	if Config.FetchIntervalSeconds <= 0 {
		return fmt.Errorf("Invalid FetchIntervalSeconds: %v", Config.FetchIntervalSeconds)
	}

	// Ensure OutputPath is a valid directory and exists
	if _, err := os.Stat(Config.OutputPath); os.IsNotExist(err) {
		return fmt.Errorf("OutputPath does not exist: %v", Config.OutputPath)
	}
	// Create subdirectories for photos and videos
	os.Mkdir(Config.OutputPath+"/photos", 0777)
	os.Mkdir(Config.OutputPath+"/videos", 0777)

	return nil
}
