package main

import (
	"io"
	"os"
	"fmt"
	"net/http"
	"log"
	"./weichat"
	"./imageDownloader"
)

func main() {
	fmt.Printf("Welcome to baby photo view backend!\n")

	imageDownloader.ImageDownloaderInit("")

	go weichatServerFunc()

	go webServerFunc()

	blockChan := make(chan struct {})
	<-blockChan

}

func webServerFunc() {
	http.HandleFunc("/", webHandlerFunction)
	err := http.ListenAndServe(":8080", nil)
	if err !=  nil {
		log.Fatal(err)
	}
}

func weichatServerFunc() {

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

func webHandlerFunction(w http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		file, err := os.Open("/kostring" + req.URL.Path)
		if err != nil {
			w.Write([]byte("File not exist!"))
			return
		}

		buf := make([]byte, 1000)
		exit := false
		for len, err := file.Read(buf);!exit;len, err = file.Read(buf) {
			if err == io.EOF {
				exit = true
			} else if err != nil {
				log.Printf("Error reading file %s: %s", file.Name(), err.Error())
				exit = true
			}

			_, err = w.Write(buf[:len])
			if err != nil {
				log.Printf("Write respond fail: %s", err.Error())
			}
		}
		return
	}

	if req.Method == "POST" {
		//TODO not support yet
		return
	}

}