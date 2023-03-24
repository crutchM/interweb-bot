package config

import (
	"github.com/spf13/viper"
	"os"
)

func New(path string, config *interface{}) (*interface{}, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(path)
	viper.SetConfigFile(currentDir + "/.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	viper.WatchConfig()
	return config, nil
}
