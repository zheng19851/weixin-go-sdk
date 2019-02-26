package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Execute(req Request) Response
}

type DefaultClient struct {
	Api string
}

func NewDefaultClient() *DefaultClient {
	return &DefaultClient{}
}

func (c DefaultClient) Execute(req Request) Response {

	response := req.GetResponse()

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, "wx41a8a8b76d194515", "bd6e5590cc1eeb6180d980628cf82b55")

	resp, err := http.Get(apiUrl)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	//getAccessTokenResponse := &GetAccessTokenResponse{"test", 0}
	json.Unmarshal(body, response)

	return response
}
