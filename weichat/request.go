package weichat

import (
	"errors"
	"bytes"
	"net/http"
	"log"
	"net/url"
)

func prepareURL(reqUrl string, param *map[string]string) (string, error) {
	var buf bytes.Buffer
	acToken, err := getAccessToken()
	if err != nil {
		return "", errors.New("Could not send get accessToken | " + err.Error())
	}

	buf.WriteString(reqUrl + "?access_token=" + acToken)
	
	if param != nil {
		log.Fatal("http requests not support parameters yet")
	}

	reqUrl = url.PathEscape(buf.String())

	return reqUrl, nil
}

func readHttpResponseBody(resp *http.Response) ([]byte, error) {
	var b []byte
	_, err := resp.Body.Read(b)

	if err != nil {
		return nil, errors.New("Get body data failed | " + err.Error())
	}

	err = resp.Body.Close()

	if  err !=  nil {
		return nil, errors.New("Close body failed | " + err.Error())
	}

	err = checkError(b)

	if err != nil {
		return nil, errors.New("Weichat return error | " + err.Error())
	}

	return b, nil
}


func sendGETRequest(reqUrl string, param *map[string]string) ([]byte, error) {
	preparedUrl, err := prepareURL(reqUrl, param)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(preparedUrl)

	if err != nil {
		return nil, errors.New("Could not send get request | " + err.Error())
	}

	b, err := readHttpResponseBody(resp)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//TODO
func sendPOSTRequest(reqUrl string, param *map[string]string, contentType string, body []byte) ([]byte, error) {
	preparedUrl, err := prepareURL(reqUrl, param)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(body)

	resp, err := http.Post(preparedUrl, contentType, reader)
	if err != nil {
		return nil, errors.New("Could not send get request | " + err.Error())
	}

	b, err := readHttpResponseBody(resp)
	if err != nil {
		return nil, err
	}	
	return b, nil
}