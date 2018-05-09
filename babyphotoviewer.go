package main

import (
	"fmt"
	"net/http"
	"log"
	"./weichat"
	"./imageDownloader"
)

func main() {
	fmt.Printf("Welcome to baby photo view backend!\n")

	imageDownloader.ImageDownloaderInit("")

	weichat.Init()

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
	imageDownloader.ImageDownloaderInsertWork(message.FromUserName, message.PicUrl)
	return nil, nil
}