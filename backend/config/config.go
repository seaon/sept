package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

var conf config

func init() {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigFile(defaultConfigFile)

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
}

func GetApp() App {
	return conf.App
}

func GetServer() Server {
	return conf.Server
}

func GetDatabase() Database {
	return conf.Database
}
