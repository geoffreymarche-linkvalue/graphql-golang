package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	files, err := ioutil.ReadDir(exPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
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
