package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetTagsRequest struct {
	AccessToken string
}

type Tag struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type GetTagsResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Tags    []Tag  `json:"tags"`
}

func NewGetTagsRequest(accessToken string) *GetTagsRequest {
	return &GetTagsRequest{accessToken}
}

func (req *GetTagsRequest) GetResponse() *GetTagsResponse {
	return &GetTagsResponse{}
}

func (res *GetTagsResponse) IsSuccess() bool {
	return res.Errcode == 0
}

func (req *GetTagsRequest) Execute() *GetTagsResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	resp, err := http.Get(apiUrl)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(string(body))

	response := &GetTagsResponse{}
	json.Unmarshal(body, response)

	return response

}
