package mediaService

import (
	"log"
	"net/http"
	"fmt"
)

var serverPath string
var imageFilePath string
var mediaServerPort int


func Init(inputImageFilePath, inputImageServerPath string, inputMediaServerPort int) {
	if inputImageFilePath == "" || inputImageServerPath == "" {
		log.Fatal("Error: Empty imageFilePath or imageServerPath!")
	}

	serverPath = inputImageServerPath
	imageFilePath = inputImageFilePath
	mediaServerPort = inputMediaServerPort

	initImageDB()
	imageDownloaderInit()
	imageServerInit()
}

func StartMediaServer() {
	go mediaServerFunc()
}

func StopMediaServer() {
	//TODO How to stop?
}

func mediaServerFunc() {
	log.Printf("media Server func, port: %d", mediaServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", mediaServerPort), nil)
	if err !=  nil {
		log.Fatal(err)
	}
}

