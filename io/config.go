package io

import(
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	FilePath string
	Map map[string]interface{}
}

func GetConfig() (config *Config) {
	// check to see if ~/.brooklyn_cli exists
	// Then Parse it to get user credentials
	config = new(Config)
	config.FilePath = filepath.Join(os.Getenv("HOME"), ".brooklyn_cli")
	if _, err := os.Stat(config.FilePath); os.IsNotExist(err) {
		config.Map = make(map[string]interface{})
		config.Write()
	}
	config.Read()
	return 
}

func (config *Config) Write() {
	
	fileToWrite, err := os.Create(config.FilePath)
	if err != nil {
		fmt.Println(err)
	}
	
	enc := json.NewEncoder(fileToWrite)
    enc.Encode(config.Map)
}

func (config *Config) Read() {
	fileToRead, err := os.Open(config.FilePath)
	if err != nil {
		fmt.Println(err)
	}
	dec := json.NewDecoder(fileToRead)
	dec.Decode(&config.Map)
}