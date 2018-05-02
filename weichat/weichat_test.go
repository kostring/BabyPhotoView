package weichat

import (
	"fmt"
	"testing"
)

func TestMenu(t *testing.T) {
	err := updateAccessToken()
	if err != nil {
		t.Error(err)
	}

	var menu *Menu = new(Menu)
	for i := 0; i < 3; i++ {
		var button MenuButton
		button.Name = fmt.Sprintf("Button%d", i)
		for j := 0; j < 5; j++ {
			var subButton MenuSubButton
			subButton.Name = fmt.Sprintf("B%dS%d", i, j)
			subButton.URL = "http://www.baidu.com"
			subButton.Type = "view"
			button.SubButton = append(button.SubButton, subButton)
		}

		menu.Buttons = append(menu.Buttons, button)
	}

	t.Log(fmt.Sprintf("%+v",menu))

	err = CreateMenu(menu)
	if err !=  nil {
		t.Error(err)
	}

	

	menu, err = QueryMenu()
	if err != nil {
		t.Error(err)
		return
	}
	t.Error(menu)
}