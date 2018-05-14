package mediaService

import (
	"fmt"
	"errors"
	"path/filepath"
	"io"
	"os"
	"net/http"
	"log"
)

func imageServerInit() {
	http.HandleFunc(serverPath + "/image/by_id/", getImageById)
	http.HandleFunc(serverPath + "/image/random/", getImageRandom)
	http.HandleFunc(serverPath + "/image", getImageList)
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
	if req.Method != "GET" {
		//TODO not support yet
		w.Write([]byte("Don't support method" + req.Method))
		return
	}

	relPath, err := filepath.Rel(serverPath + "/image/by_id/", req.URL.Path)
	if err != nil {
		log.Printf("Decode URL path fail, path: %s, err: %s", req.URL.Path, err.Error())
	}

	filePath := filepath.Join(imageFilePath, relPath)

	if err = getImage(filePath, w); err != nil {
		log.Print(err.Error())
	}

}

func getImage(path string, w http.ResponseWriter) error {
	file, err := os.Open(path)
	if err != nil {
		w.Write([]byte("File not exist!"))
		return errors.New("Requested not exist file, path: " + path)
	}
	defer file.Close()

	buf := make([]byte, 1000)
	exit := false
	for len, err := file.Read(buf);!exit;len, err = file.Read(buf) {
		if err == io.EOF {
			return nil
		} else if err != nil {
			return errors.New(fmt.Sprintf("Error reading file %s: %s", file.Name(), err.Error()))
		}

		_, err = w.Write(buf[:len])
		if err != nil {
			return errors.New(fmt.Sprintf("Write respond fail: %s", err.Error()))
		}
	}

	//Should never get here
	return nil
}