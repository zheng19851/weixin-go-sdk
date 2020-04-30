package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DeleteTagRequest struct {
	AccessToken string `json:"access_token"`
	Id          int
}

type TagInfo4Delete struct {
	Id int `json:"id"`
}

type DeleteTagResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func NewDeleteTagRequest(accessToken string, id int) *DeleteTagRequest {
	return &DeleteTagRequest{accessToken, id}
}

func (req *DeleteTagRequest) GetResponse() *DeleteTagResponse {
	return &DeleteTagResponse{}
}

func (res *DeleteTagResponse) IsSuccess() bool {
	return res.Errcode == 0
}

func (req *DeleteTagRequest) Execute() *DeleteTagResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	tag := &TagInfo4Delete{req.Id}

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

	response := &DeleteTagResponse{}
	json.Unmarshal(body, response)

	return response

}
