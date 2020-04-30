package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UpdateTagRequest struct {
	AccessToken string `json:"access_token"`
	Id          int
	Name        string `json:"name"`
}

type TagInfo4Update struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateTagResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func NewUpdateTagRequest(accessToken string, id int, name string) *UpdateTagRequest {
	return &UpdateTagRequest{accessToken, id, name}
}

func (req *UpdateTagRequest) GetResponse() *UpdateTagResponse {
	return &UpdateTagResponse{}
}

func (res *UpdateTagResponse) IsSuccess() bool {
	return res.Errcode == 0
}

func (req *UpdateTagRequest) Execute() *UpdateTagResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	tag := &TagInfo4Update{req.Id, req.Name}

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

	response := &UpdateTagResponse{}
	json.Unmarshal(body, response)

	return response

}
