package config

import (
	"log"

	"hexagonal-template/pkg/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ",squash" will ignore mapstructure for that field

type Config struct {
	App appConfig `mapstructure:",squash"`
}

type appConfig struct {
	HTTPPort string `mapstructure:"APP_HTTP_PORT"`
	Env      string `mapstructure:"APP_ENV"`
}

var config Config

func Init(cfgPath string) {

	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.AppLogger.Infof("Config file has changed from %s with %s operation", e.Name, e.Op.String())
	})
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetViper() *Config {
	return &config
}
