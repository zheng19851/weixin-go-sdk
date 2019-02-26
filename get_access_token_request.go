package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetAccessTokenRequest struct {
	Appid  string
	Secret string
}

type GetAccessTokenResponse struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}

func NewGetAccessTokenRequest(appid string, secret string) *GetAccessTokenRequest {
	return &GetAccessTokenRequest{appid, secret}
}

func (req GetAccessTokenRequest) GetResponse() Response {
	return &GetAccessTokenResponse{}
}

func (res GetAccessTokenResponse) IsSuccess() bool {
	return true
}

func (req GetAccessTokenRequest) Execute() *GetAccessTokenResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.Appid, req.Secret)

	resp, err := http.Get(apiUrl)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	getAccessTokenResponse := &GetAccessTokenResponse{"test", 0}
	json.Unmarshal(body, getAccessTokenResponse)

	return getAccessTokenResponse

}
