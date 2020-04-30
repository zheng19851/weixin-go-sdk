package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CreateTagRequest struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}

type TagInfo struct {
	Name string `json:"name"`
}

type CreateTagResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func NewCreateTagRequest(accessToken string, name string) *CreateTagRequest {
	return &CreateTagRequest{accessToken, name}
}

func (req *CreateTagRequest) GetResponse() *CreateTagResponse {
	return &CreateTagResponse{}
}

func (res *CreateTagResponse) IsSuccess() bool {
	return res.Errcode == 0
}

func (req *CreateTagRequest) Execute() *CreateTagResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	tag := &TagInfo{req.Name}

	requestParams := map[string]interface{}{
		"tag": tag, // 定义时指定的初始key/value, 后面可以继续添加
	}

	b, err := json.Marshal(requestParams)

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

	fmt.Println(string(body))

	response := &CreateTagResponse{}
	json.Unmarshal(body, response)

	return response

}
