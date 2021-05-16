package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Mysql struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     int    `json:"port"`
	Address  string `json:"address"`
}

type Configuration struct {
	Mysql Mysql `json:"mysql"`
}

func Load() *Configuration {
	var configuration Configuration
	jsonFile, err := os.Open("./app/config.json")
	if err != nil {
		panic(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteValue, &configuration)
	if err != nil {
		panic(err)
	}
	return &configuration
}
