package main

import (
	"errors"
	"bytes"
	"sort"
	"net/url"
	"fmt"
	"net/http"
	"log"
	"crypto/sha1"
)

const (
	token = "kostring112233"
)

func processVerification(parameters url.Values) ([]byte, error) {
	inputArgs := [3]string{token, parameters["timestamp"][0], parameters["nonce"][0]}
	argsSlice := inputArgs[:]

	sort.Strings(argsSlice)
	var buf bytes.Buffer
	for _, s := range argsSlice {
		buf.WriteString(s)
	}

	sha1Result := sha1.Sum(buf.Bytes())

	//Reuse buf to get the correct sha1 result string
	buf.Reset()

	for _, b := range sha1Result {
		buf.WriteString(fmt.Sprintf("%.2x", b))
	}

	sha1String := buf.String()

	if sha1String != parameters["signature"][0] {
		log.Printf("Invalid verification request! Signature %s, expect: %s", parameters["signature"][0], sha1String)
		return nil, errors.New("Invalid verification request!")
	}

	return []byte(parameters["echostr"][0]), nil
}

func photoViewServer(w http.ResponseWriter, req *http.Request) {
	if req.ParseForm() != nil {
		log.Fatal("Failed to parse req!")
	}

	parameters := req.Form
	returnInfo, err := processVerification(parameters)
	if err != nil {
		return
	}
	w.Write(returnInfo)
}

func main() {
	fmt.Printf("Welcome to baby photo view backend!")

	http.HandleFunc("/weichat", photoViewServer)
	err := http.ListenAndServe(":80", nil)
	if err !=  nil {
		log.Fatal(err)
	}
}