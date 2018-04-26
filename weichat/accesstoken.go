package weichat

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AccessTokenData struct {
	AccessToken string 	`json:"access_token"`
	ExpiresIn int 		`jsion:"expires_in"`
}

var accessToken string = ""

func getAccessToken() string {
	return accessToken
}

func updateAccessToken() error {
	resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + appID + "&secret=" + secret)
	if err != nil {
		log.Printf("Failed to send request! Err: %s", err.Error())
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read body! Err: %s", err.Error())
		resp.Body.Close()
		return	err	
	}

	resp.Body.Close()

	var accessTokenData AccessTokenData 
	err = json.Unmarshal(b, accessTokenData)
	if err != nil {
		log.Printf("Failed to decode JSON! Err: %s", err.Error())
		return err
	}

	accessToken = accessTokenData.AccessToken
	return nil
}
