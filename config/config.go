package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Env string
	}
	DB struct {
		SQL struct {
			Username     string
			Password     string
			Host         string
			Port         string
			DatabaseName string
		}
		Redis struct {
			Address string
		}
	}
}

func New(fileName string) (*Config, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
