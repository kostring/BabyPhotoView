package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"log"
)

func photoViewServer(w http.ResponseWriter, req *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	reqBody := buf.String()
	fmt.Print("Body:\n")
	fmt.Print(reqBody)
	fmt.Print("\nBody end\n")

	io.WriteString(w, "Welcome to photo viewer!\n")
}

func main() {
	fmt.Printf("Welcome to baby photo view backend!")

	http.HandleFunc("/weichat", photoViewServer)
	err := http.ListenAndServe(":80", nil)
	if err !=  nil {
		log.Fatal(err)
	}
}