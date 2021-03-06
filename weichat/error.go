package weichat

import (
	"fmt"
	"errors"
	"encoding/json"
)

type ErrorInfo struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
} 

func checkError(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	var errorInfo ErrorInfo
	err := json.Unmarshal(b, &errorInfo)
	if err != nil {
		return nil
	}

	if errorInfo.Errcode != 0 || errorInfo.Errmsg != "" {
		return errors.New("Weichat error: code: " + fmt.Sprintf("%d", errorInfo.Errcode)  + " , msg: " + errorInfo.Errmsg)
	}
	return nil
}