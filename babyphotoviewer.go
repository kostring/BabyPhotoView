package main

import (
	"fmt"
	"net/http"
	"log"
	"./weichat"
)

func main() {
	fmt.Printf("Welcome to baby photo view backend!\n")

	weichat.Init()

	http.HandleFunc("/weichat", weichat.WeichatHandleFunction)
	err := http.ListenAndServe(":80", nil)
	if err !=  nil {
		log.Fatal(err)
	}
}