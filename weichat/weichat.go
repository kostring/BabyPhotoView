package weichat

import (
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
	"log"
)

const (
	token = "kostring112233"
)

type TextMessage struct {
	ToUserName string
	FromUserName string
	CreateTime time.Time
	MsgType string
	Content string
	MsgId string
}

func ProcessTextMessage() {
	fmt.Print("Process text message.\n")

}

func WeichatHandleFunction(w http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		weichatGetReqHandler(w, req)
		return
	}

	if req.Method == "POST" {
		weichatPostReqHandler(w, req)
		return
	}

}

func weichatPostReqHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print("Error while read body, error info: " + err.Error())
		return
	}

	log.Print(string(body))
}


func weichatGetReqHandler(w http.ResponseWriter, req *http.Request) {
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

