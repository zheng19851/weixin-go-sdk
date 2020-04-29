package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetMenuRequest struct {
	AccessToken string
}

type SelfmenuInfo struct {
	Button []Button
}

type GetMenuResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`

	Is_menu_open string `json:"is_menu_open"`

	Selfmenu_info SelfmenuInfo `json:"selfmenu_info"`
}

func NewGetMenuRequest(accessToken string) *GetMenuRequest {
	return &GetMenuRequest{accessToken}
}

func (req *GetMenuRequest) GetResponse() *GetMenuResponse {
	return &GetMenuResponse{}
}

func (res *GetMenuResponse) IsSuccess() bool {
	return res.Errcode == 0
}

func (req *GetMenuRequest) Execute() *GetMenuResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	resp, err := http.Get(apiUrl)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(string(body))

	response := &GetMenuResponse{}
	json.Unmarshal(body, response)

	return response

}
