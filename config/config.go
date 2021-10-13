package config

import (
	"log"
	"reflect"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	App `mapstructure:"app"`
}

type App struct {
	Debug bool   `mapstructure:"debug"`
	Port  string `mapstructure:"port"`
}

var config Config

func (cfg Config) IsInited() bool {
	return !reflect.DeepEqual(cfg, Config{})
}

func Init(cfgPath, env string) {
	switch env {
	case "local":
		viper.SetConfigName("config.local")
	case "develop":
		viper.SetConfigName("config.develop")
	default:
		viper.SetConfigName("config")
	}
	viper.AddConfigPath(cfgPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file has changed: ", e.Name)
	})
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}

}

func GetViper() *Config {
	return &config
}
