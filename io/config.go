package io

import (
	"encoding/json"
	"github.com/apache/brooklyn-client/error_handler"
	"os"
	"path/filepath"
)

type Config struct {
	FilePath string
	Map      map[string]interface{}
}

func GetConfig() (config *Config) {
	// check to see if $BRCLI_HOME/.brooklyn_cli or $HOME/.brooklyn_cli exists
	// Then parse it to get user credentials
	config = new(Config)
	if os.Getenv("BRCLI_HOME") != "" {
		config.FilePath = filepath.Join(os.Getenv("BRCLI_HOME"), ".brooklyn_cli")
	} else {
		config.FilePath = filepath.Join(os.Getenv("HOME"), ".brooklyn_cli")
	}
	if _, err := os.Stat(config.FilePath); os.IsNotExist(err) {
		config.Map = make(map[string]interface{})
		config.Write()
	}
	config.Read()
	return
}

func (config *Config) Write() {

	// Create file as read/write by user (but does not change perms of existing file)
	fileToWrite, err := os.OpenFile(config.FilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		error_handler.ErrorExit(err)
	}

	enc := json.NewEncoder(fileToWrite)
	enc.Encode(config.Map)
}

func (config *Config) Read() {
	fileToRead, err := os.Open(config.FilePath)
	if err != nil {
		error_handler.ErrorExit(err)
	}
	dec := json.NewDecoder(fileToRead)
	dec.Decode(&config.Map)
}
