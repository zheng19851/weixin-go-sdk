package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Button struct {
	Menu_type  string   `json:"type"`
	Name       string   `json:"name"`
	Url        string   `json:"url"`
	Key        string   `json:"key"`
	Sub_button []Button `json:"sub_button"`
}

type CreateMenuRequest struct {
	AccessToken string   `json:"access_token"`
	Button      []Button `json:"button"`
}

type CreateMenuResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func NewCreateMenuRequest(accessToken string, button []Button) *CreateMenuRequest {
	return &CreateMenuRequest{accessToken, button}
}

func (req *CreateMenuRequest) GetResponse() *CreateMenuResponse {
	return &CreateMenuResponse{}
}

func (res *CreateMenuResponse) IsSuccess() bool {
	return res.Errcode == 0
}

func (req *CreateMenuRequest) Execute() *CreateMenuResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	b, err := json.Marshal(req)

	//fmt.Printf("requestBody:%s", string(b))

	if err != nil {
		log.Println("json format error:", err)
		return nil
	}

	request_body := bytes.NewBuffer(b)

	resp, err := http.Post(apiUrl, "application/json;charset=utf-8", request_body)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(string(body))

	response := &CreateMenuResponse{}
	json.Unmarshal(body, response)

	return response

}
