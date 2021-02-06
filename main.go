package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Assets []Asset `yaml:"assets"`
}
type Asset struct {
	Name  string      `yaml:"name"`
	Asset string      `yaml:"asset"`
	Users []AssetUser `yaml:"users"`
}
type AssetUser struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}

func config() *Config {
	buffer, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Please create a config.yaml")
	}

	config := &Config{}
	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	config := config()
	fmt.Println(config)
}
