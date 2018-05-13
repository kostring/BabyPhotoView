package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type WeichatConfig struct {
	Token string `json:"token"`
	AppID string `json:"appID"`
	Secret string `json:"secret"`
}

type StorageConfig struct {
	ImageFilePath string `json:"imageFilePath"`
}

type AppConfig struct {
	Weichat WeichatConfig `json:"weichatConfig"`
	Storage StorageConfig `json:"StorageConfig"`
	Test string
}


func loadConfig(configFilePath string) AppConfig {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Error: Could not open config file! Err: " + err.Error())
		return AppConfig{}
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error: failed reading config file! Err: " + err.Error())
		return AppConfig{}
	}

	fmt.Print(string(buf))

	var appConfig AppConfig

	if err = json.Unmarshal(buf, &appConfig); err != nil {
		log.Fatal("Error: could not unmarshal config file! Err: " + err.Error())
	}

	return appConfig
}