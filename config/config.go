package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Redis struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}
	N       int     `json:"N"`
	K       int     `json:"K"`
	UserIds []int64 `json:"UserIds"`
}

func Read(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, err
}
