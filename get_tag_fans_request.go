package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GetTagFansRequest struct {
	AccessToken string

	TagId      int    `json:tagid`
	NextOpenid string `json:next_openid`
}

type TagFansData struct {
	Openids []string `json:"openid"`
}

type GetTagFansResponse struct {
	Errcode     int         `json:"errcode"`
	Errmsg      string      `json:"errmsg"`
	Count       int         `json:"count"`
	TagFansData TagFansData `json:"data"`

	NextOpenid string `json:"next_openid"`
}

func NewGetTagFansRequest(accessToken string, tagId int, nextOpenId string) *GetTagFansRequest {
	return &GetTagFansRequest{accessToken, tagId, nextOpenId}
}

func (req *GetTagFansRequest) GetResponse() *GetTagFansResponse {
	return &GetTagFansResponse{}
}

func (res *GetTagFansResponse) IsSuccess() bool {
	return res.Errcode == 0
}

func (req *GetTagFansRequest) Execute() *GetTagFansResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	b, err := json.Marshal(req)

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

	response := &GetTagFansResponse{}
	json.Unmarshal(body, response)

	return response

}
