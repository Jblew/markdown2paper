package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config is the config of markdown2paper
type Config struct {
	BibPath string `yaml:"bibPath"`
	OutlineFile string `yaml:"outlinePath"`
	OutFile string `yaml:"outPath"`
}

func loadConfigFromFile(configPath string) (Config, error) {
	contents, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	c := Config{}
	err = yaml.Unmarshal(contents, &c)
	return c, err
}
