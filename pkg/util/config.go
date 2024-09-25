package util

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	// The URL of the NatureWatch server. This can be an IP address or a hostname
	// but must include the protocol (generally "http://")
	NatureWatchURL string `mapstructure:"NATUREWATCH_URL"`
	// The number of seconds to wait between fetches. If this is 0 or negative,
	// the process will exit after the first fetch.
	FetchIntervalSeconds int `mapstructure:"FETCH_INTERVAL_SECONDS"`
	// The output path on the local filesystem where photos and videos will be saved.
	// This path must already exist. However, the photos and videos subdirectories
	// will be created if they do not exist.
	OutputPath string `mapstructure:"OUTPUT_PATH"`
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

	// Ensure OutputPath is a valid directory and exists
	if _, err := os.Stat(Config.OutputPath); os.IsNotExist(err) {
		return fmt.Errorf("OutputPath does not exist: %v", Config.OutputPath)
	}
	// Create subdirectories for photos and videos
	os.Mkdir(Config.OutputPath+"/photos", 0777)
	os.Mkdir(Config.OutputPath+"/videos", 0777)

	return nil
}
