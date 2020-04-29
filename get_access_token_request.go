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

func (req *GetAccessTokenRequest) ApiName() string {
	return "token"
}

func (req *GetAccessTokenRequest) GetParams() map[string]string {

	return map[string]string{"grant_type": "client_credential", "appid": req.Appid, "secret": req.Secret}
}

func (req *GetAccessTokenRequest) GetHttpMethod() string {
	return "get"
}

func (req *GetAccessTokenRequest) GetResponse() Response {
	return &GetAccessTokenResponse{}
}

func (res *GetAccessTokenResponse) IsSuccess() bool {
	return res.Access_token != ""
}

func (req *GetAccessTokenRequest) Execute() *GetAccessTokenResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.Appid, req.Secret)

	resp, err := http.Get(apiUrl)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	//fmt.Println(string(body))

	response := &GetAccessTokenResponse{}
	json.Unmarshal(body, response)

	return response

}
