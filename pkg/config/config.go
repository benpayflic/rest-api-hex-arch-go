package config

import (
	"github.com/spf13/viper"
)

// Config stores all the configurations of the application.
// The values are read by viper from a configuration file or enviroment variable
type Config struct {
	AppPort string `mapstructure:"APP_PORT"`
	AppHost string `mapstructure:"APP_HOST"`
}

// Load configurations from path and unmarshal to Config
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
