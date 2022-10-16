package main

import (
	"fmt"
	"gotemplate/pkg/config"
)

// Config is the configuration for the application
type Config struct {
	Appname string
	Port    int

	Mongo struct {
		URI string
	}
}

func readConfig(env string) Config {
	var conf Config

	filepath := fmt.Sprintf(".config/%s.yml", env)

	if err := config.Read(filepath, "GOTEMPLATE", &conf); err != nil {
		panic(err)
	}

	return conf
}
