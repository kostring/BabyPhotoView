package weichat

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"fmt"
	"log"
)

type Message struct {
	ToUserName string
	FromUserName string
	CreateTime int64
	MsgType string
	MsgId string

	Content string

	PicUrl string
	MediaId string

	Format string
	Recognition string

	ThumbMediaId string

	Location_X string
	Location_Y string
	Scale string
	Label string

	Title string
	Description string
	Url string

	
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
	var message Message 

	xml.Unmarshal(body, &message)

	log.Print(string(body))
	log.Printf("%+v", message)

	if message.MsgType == "text" {
		w.Write(weichatTextMsgHandler(message))
	}
	
	if message.MsgType == "image" {
		w.Write(weichatImageMsgHandler(message))
	}
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


func weichatTextMsgHandler(message Message) []byte {
	var ret []byte = []byte("text")

	return ret
}

func weichatImageMsgHandler(message Message) []byte {
	fmt.Printf("%+v", message)
	var ret []byte = []byte("image")
	return ret
}

func Init() {
	err := updateAccessToken()
	if err != nil {
		log.Fatal("Could not get access token!")
	}
}


