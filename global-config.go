package main

import (
	"github.com/BurntSushi/toml"
	"os"
)

type GlobalConfig struct {
	BaseDirectory string
	StateFile     string
}

func setupConfig() GlobalConfig {
	var config GlobalConfig
	configFile := "/etc/shpack.conf"

	_, err := os.Stat(configFile)
	if err != nil {
		// Empty config. use default value
	} else {
		// TODO: Read config file
		_, err := toml.DecodeFile(configFile, &config)
		if err != nil {
			panic(err)
		}
	}
	return config
}
