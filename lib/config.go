package lib

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Application `yaml:"application"`
}

type Application struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	Mode string `yaml:"mode"`
}

func GetConfiguration(path string) (*Configuration, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("Can't close the file descriptor {%v}", err)
		}
	}(f)

	var cfg Configuration
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
