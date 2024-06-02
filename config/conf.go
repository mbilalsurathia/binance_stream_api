package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func GetConf(path string) Config {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Read config file failed, error: %v \n", err.Error())
		os.Exit(0)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Unmarshal config data failed, error: %v \n", err.Error())
		os.Exit(0)
	}
	return config
}
