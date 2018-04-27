package weichat

import (
	"errors"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AccessTokenData struct {
	AccessToken string 	`json:"access_token"`
	ExpiresIn int 		`jsion:"expires_in"`
}

const (
	invalidAccessToken = "Invald_Token"
)

var accessToken string = invalidAccessToken

func getAccessToken() (string, error) {
	if accessToken == invalidAccessToken {
		return "", errors.New("Invalid access token!")
	}
	return accessToken, nil
}

func updateAccessToken() error {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appID + "&secret=" + secret)
	if err != nil {
		log.Printf("Failed to send request! Err: %s", err.Error())
		accessToken = invalidAccessToken
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read body! Err: %s", err.Error())
		resp.Body.Close()
		accessToken = invalidAccessToken
		return	err	
	}

	resp.Body.Close()

	err = checkError(b)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	var accessTokenData AccessTokenData 
	err = json.Unmarshal(b, &accessTokenData)
	if err != nil {
		log.Printf("Failed to decode JSON! Err: %s", err.Error())
		accessToken = invalidAccessToken
		return err
	}

	accessToken = accessTokenData.AccessToken

	log.Print("Updated Access Token: " + accessToken)
	return nil
}
