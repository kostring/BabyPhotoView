package weichat

import (
	"fmt"
	"errors"
	"encoding/json"
	"net/url"
)

const (
	menuButtonCountMin = 1
	menuButtonCountMax = 3
	menuSubButtonCountMin = 1
	menuSubButtonCountMax = 5

	menuButtonNameByteMax = 16
	menuSubButtonNameByteMax = 50

	menuKeyByteMax = 128
	menuURLByteMax = 1024
)

type MenuSubButton struct {
	Type string `json:type`
	Name string `json:name`
	Key string `json:key`
	URL string `json:url`
	MediaId string `json:media_id`
	AppId string `json:appid`             //Weichat mini program only
	Pagepath string `json:pagepath`       //Weichat mini program only
}

type MenuButton struct {
	Type string `json:type`
	Name string `json:name`
	Key string `json:key`
	URL string `json:url`
	MediaId string `json:media_id`
	AppId string `json:appid`             //Weichat mini program only
	Pagepath string `json:pagepath`       //Weichat mini program only
	SubButton []MenuSubButton `json:sub_button`
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

func DeleteMenu() error {
	_, err := sendGETRequest("https://api.weixin.qq.com/cgi-bin/menu/delete", nil)
	if err !=  nil {
		return  errors.New("Failed to request menu | " + err.Error())
	}

	return nil	
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

func checkSingleButtonContent(buttonType, key, buttonUrl, media_id, appid, pagepath string) error {
	if len([]byte(key)) > menuKeyByteMax {
		return errors.New(fmt.Sprintf("Key too long: %s", key))
	}

	//Url check
	if buttonUrl != "" {
		escapedUrl := url.PathEscape(buttonUrl)
		if escapedUrl != buttonUrl {
			return errors.New(fmt.Sprintf("URL include invalid character!: %s", buttonUrl))
		}
	
		if len(buttonUrl) > menuURLByteMax {
			return errors.New(fmt.Sprintf("Url too long: %s", buttonUrl))
		}
	}

	//Type check
	//Do not do the type check, let weichat do this
	//TODO add the type check

	//
	return nil
}



func checkMenu(menu *Menu) error {
	if len(menu.Buttons) < menuButtonCountMin || len(menu.Buttons) > menuButtonCountMax {
		return errors.New(fmt.Sprintf("Invalid menu button count: %d", len(menu.Buttons)))
	}

	for index, button := range menu.Buttons {
		if len([]byte(button.Name)) > menuButtonNameByteMax {
			return errors.New(fmt.Sprintf("Button index %d mane too long: %s", index, button.Name))
		}
		
		//If any sub button exist, the button should noly have name attribute
		if len(button.SubButton) != 0 {
			if button.Type != "" || button.URL != "" || button.Key != "" || button.MediaId != "" || button.AppId != "" || button.Pagepath != "" {
				return errors.New("Buttons with sub button should not have attribute other than name")
			}

			if len(button.SubButton) > menuSubButtonCountMax {
				return errors.New(fmt.Sprintf("Button index %d sub button count invalid: %d", index, len(button.SubButton)))
			}

			for subButtonIndex, subButton := range button.SubButton {
				if len([]byte(subButton.Name)) > menuSubButtonNameByteMax {
					return errors.New(fmt.Sprintf("Button index %d sub button index: %d mane too long: %s", index, subButtonIndex, subButton.Name))
				}

				err := checkSingleButtonContent(subButton.Type, subButton.Key, subButton.URL, subButton.MediaId, subButton.AppId, subButton.Pagepath)
				if err != nil {
					return errors.New("Button index %d sub button index: %d content error | " + err.Error())
				}
			}

		} else {
			err := checkSingleButtonContent(button.Type, button.Key, button.URL, button.MediaId, button.AppId, button.Pagepath)
			if err != nil {
				return errors.New("Button index %d content error | " + err.Error())
			}		
		}
	}

	return nil
}
