package mediaService

import (
	"io"
	"os"
	"fmt"
	"net/http"
	"log"
)

func imageServerInit() {
	http.HandleFunc(serverPath + "/image/by_id/", getImageById)
	http.HandleFunc(serverPath + "/image/random/", getImageById)
	http.HandleFunc(serverPath + "/image", getImageById)
}

func getImageList(w http.ResponseWriter, req *http.Request) {
	//TODO
	w.Write([]byte("Get image list, not done yet"))
}

func getImageRandom(w http.ResponseWriter, req *http.Request) {
	//TODO
	w.Write([]byte("Get random image, not done yet"))
}

func getImageById(w http.ResponseWriter, req *http.Request) {
	fmt.Print("web handler! " + req.URL.Path)
	if req.Method == "GET" {

		
		file, err := os.Open(imageFilePath + req.URL.Path)
		if err != nil {
			w.Write([]byte("File not exist!"))
			return
		}
		defer file.Close()

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
		w.Write([]byte("Don't support post method"))
		return
	}

}