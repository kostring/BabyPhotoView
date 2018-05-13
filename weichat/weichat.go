package weichat

import (
	"errors"
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

type MessageHandler struct {
	MsgType string
	HandlerFunc func(Message) ([]byte, error)
}

var handlers []MessageHandler

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

	for _, handler := range handlers {
		if handler.MsgType == message.MsgType {
			ret, err := handler.HandlerFunc(message)
			if err != nil {
				log.Printf("Failed to process message: %+v\n Error: %s", message, err.Error())
				return
			}
			
			w.Write(ret)
			return
		}
	}

	// If no handler for this function
	log.Panicf("Unhandled message: ----\n%+v\n----", message)
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

	if returnInfo != nil {
		w.Write(returnInfo)
	}
	
}

var token string
var appID string
var secret string

func Init(inputToken, inputAppID, inputSecret string) {
	token = inputToken
	appID = inputAppID
	secret = inputSecret
	err := updateAccessToken()
	if err != nil {
		log.Fatal("Could not get access token!")
	}
}

func RegisterMsgHandler(msgType string, handlerFunc func(Message) ([]byte, error)) error{

	if handlerFunc == nil {
		return errors.New(fmt.Sprintf("Register message handler function fail. Type: %s, nil function!", msgType))
	}

	for _, handler := range handlers {
		if handler.MsgType == msgType {
			return errors.New(fmt.Sprintf("Register message handler function fail. Message type %s exists: %v", msgType, handler.HandlerFunc))
		}
	}

	handlers = append(handlers,MessageHandler{msgType, handlerFunc})
	return nil
}