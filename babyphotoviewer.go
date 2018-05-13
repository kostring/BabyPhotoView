package main

import (
	"fmt"
	"net/http"
	"log"
	"./weichat"
	"./mediaService"
)

const (
	configFile = "./config.json"
)

var appConfig AppConfig

func main() {
	fmt.Printf("Welcome to baby photo viewer backend!\n")

	appConfig = loadConfig(configFile)

	fmt.Printf("%+v\n", appConfig)

	mediaService.Init(appConfig.MediaService.ImageFilePath, appConfig.MediaService.MediaServicePath, appConfig.MediaService.MediaServicePort)
	mediaService.StartMediaServer()

	go weichatServerFunc()

	blockChan := make(chan struct {})
	<-blockChan

}

func weichatServerFunc() {
	//TODO this should move into package weichat
	weichat.Init(appConfig.Weichat.Token, appConfig.Weichat.AppID, appConfig.Weichat.Secret)
	weichat.RegisterMsgHandler("image", handleImageMessage)
	http.HandleFunc("/weichat", weichat.WeichatHandleFunction)
	err := http.ListenAndServe(":80", nil)
	if err !=  nil {
		log.Fatal(err)
	}
}

func handleImageMessage(message weichat.Message) ([]byte, error) {
	log.Printf("Handle image message, will download it")
	log.Printf("%+v", message)
	mediaService.ImageDownloaderInsertWork(message.FromUserName, message.PicUrl)
	return nil, nil
}

