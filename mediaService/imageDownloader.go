package mediaService

import (
	"path/filepath"
	"log"
	"time"
	"os/exec"
	"fmt"
)

const (
	maxOutstandingWorkItems = 1000
	outstandingWorkItemWarn = 900
)

type DownloadWorkItem struct {
	OpenId string
	Url    string
}

var downloadWorkItems chan DownloadWorkItem

func downloadRoutine(items <-chan DownloadWorkItem) {

	log.Printf("Downloader start listening.")

	for item := range items {

		fmt.Print(item)

		if len(items) > outstandingWorkItemWarn {
			log.Printf("Warn: Outstanding work item high: %d", len(items))
		}

		//TODO decide extension name by format
		fileName := item.OpenId + time.Now().Format("2006_01_02_15_04_05.99999") + ".jpg"
		cmd := exec.Command("wget", item.Url, "-O", filepath.Join(imageFilePath, fileName))

		log.Print("Download file: " + filepath.Join(imageFilePath, fileName))

		_, err := cmd.Output()

		if err != nil {
			log.Printf("Error: Failed to download image, error: " + err.Error())
		}

		//Update the imageDB
		imageDBInsert(filepath.Join(imageFilePath, fileName))
	}

	log.Print("Downloader end listen")
}

func imageDownloaderInit() {
	downloadWorkItems = make(chan DownloadWorkItem, maxOutstandingWorkItems)
	go downloadRoutine(downloadWorkItems)
}

func ImageDownloaderInsertWork(openId, url string) {
	var item DownloadWorkItem
	item.OpenId = openId
	item.Url = url
	downloadWorkItems <- item
}

func ImageDownloaderGetOutstandingWorks() int {
	return len(downloadWorkItems)
}