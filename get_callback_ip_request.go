package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetCallbackIpRequest struct {
	AccessToken string
}

type GetCallbackIpResponse struct {
	IpList []string `json:"ip_list"`
}

func NewGetCallbackIpRequest(accessToken string) *GetCallbackIpRequest {
	return &GetCallbackIpRequest{accessToken}
}

func (req *GetCallbackIpRequest) ApiName() string {
	return "getcallbackip"
}

func (req *GetCallbackIpRequest) GetParams() map[string]string {
	return map[string]string{"access_token": req.AccessToken}
}

func (req *GetCallbackIpRequest) GetHttpMethod() string {
	return "get"
}

func (req *GetCallbackIpRequest) GetResponse() Response {
	return &GetCallbackIpResponse{}
}

func (res *GetCallbackIpResponse) IsSuccess() bool {
	return res.IpList != nil
}

func (req *GetCallbackIpRequest) Execute() *GetCallbackIpResponse {

	apiUrlPattern := "https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.AccessToken)

	resp, err := http.Get(apiUrl)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	response := &GetCallbackIpResponse{}
	json.Unmarshal(body, response)

	return response

}
