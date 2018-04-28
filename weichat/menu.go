package weichat

import (
	"log"
	"errors"
	"encoding/json"
)

type MenuButton struct {
	Type string `json:type`
	Name string `json:name`
	Key string `json:key`
	URL string `json:url`
	SubButton []MenuButton `json:sub_button`
}

type Menu struct {
	Buttons []MenuButton `json:button`
}
/*
{
    "menu": {
        "button": [
            {
                "type": "click", 
                "name": "今日歌曲", 
                "key": "V1001_TODAY_MUSIC", 
                "sub_button": [ ]
            }, 
            {
                "type": "click", 
                "name": "歌手简介", 
                "key": "V1001_TODAY_SINGER", 
                "sub_button": [ ]
            }, 
            {
                "name": "菜单", 
                "sub_button": [
                    {
                        "type": "view", 
                        "name": "搜索", 
                        "url": "http://www.soso.com/", 
                        "sub_button": [ ]
                    }, 
                    {
                        "type": "view", 
                        "name": "视频", 
                        "url": "http://v.qq.com/", 
                        "sub_button": [ ]
                    }, 
                    {
                        "type": "click", 
                        "name": "赞一下我们", 
                        "key": "V1001_GOOD", 
                        "sub_button": [ ]
                    }
                ]
            }
        ]
    }
}
*/
func QueryMenu() (*Menu, error) {
	b, err := sendGETRequest("https://api.weixin.qq.com/cgi-bin/menu/get", nil)
	if err !=  nil {
		return  nil, errors.New("Failed to request menu | " + err.Error())
	}
	
	var menu Menu;
	err = json.Unmarshal(b, &menu)
	if err !=  nil {
		return  nil, errors.New("Failed to decode json \"" + string(b) + "\" | " + err.Error())
	}

	return &menu, nil
}

func CreateMenu(menu *Menu) error {

	err := checkMenu(menu)
	if err != nil {
		return errors.New("Invalid Menu! | " + err.Error())
	}

	b, err := json.Marshal(menu)
	if err != nil {
		return errors.New("Menu Json format invald! | " + err.Error())
	}

	b, err = sendPOSTRequest("https://api.weixin.qq.com/cgi-bin/menu/create", nil, b)
	if err != nil {
		return errors.New("Failed to send request to weichat! | " + err.Error())
	}

	return nil
}


//TODO
func checkMenu(menu *Menu) error {
	log.Fatal("checkMenu not done yet")
	return nil
}
