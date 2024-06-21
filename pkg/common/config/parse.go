package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	baseConfigPath := "./config/config.yaml"

	v := viper.New()
	v.SetConfigFile(baseConfigPath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error base config file: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
	})

	if err := v.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	envConfigPath := fmt.Sprintf("./config/config-%s.yaml", Config.App.Env)
	v.SetConfigFile(envConfigPath)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error %s config file: %s \n", Config.App.Env, err))
	}

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
	})

	if err := v.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
