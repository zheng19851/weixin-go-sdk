package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeleteMenuRequest struct {
	AccessToken string
}

type DeleteMenuResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func NewDeleteMenuRequest(accessToken string) *DeleteMenuRequest {
	return &DeleteMenuRequest{accessToken}
}

func (req *DeleteMenuRequest) GetResponse() *DeleteMenuResponse {
	return &DeleteMenuResponse{}
}

func (res *DeleteMenuResponse) IsSuccess() bool {
	return res.Errcode == 0
}

func (req *DeleteMenuRequest) Execute() *DeleteMenuResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	resp, err := http.Get(apiUrl)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(string(body))

	response := &DeleteMenuResponse{}
	json.Unmarshal(body, response)

	return response

}
