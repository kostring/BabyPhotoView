package mediaService

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

var imageDB []string


//TODO move this to real DB
func initImageDB() {
	err := filepath.Walk(imageFilePath, func(path string, info os.FileInfo, err error) error {
		if err !=  nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		imageDB = append(imageDB, path)
		return nil
	})

	if err != nil {
		log.Printf("Error init image DB! " + err.Error())
	}
}

func imageDBInsert(path string) {
	imageDB = append(imageDB, path)
}

func imageDBGet(index int) (string, error) {
	if len(imageDB) <= index{
		return "", errors.New("Index out of range!")
	}
	return imageDB[index], nil
}

func imageDBLen() int {
	return len(imageDB)
}
