package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	Execute(req Request) Response
}

type DefaultClient struct {
	Server string
}

func NewDefaultClient(server string) *DefaultClient {
	return &DefaultClient{server}
}

func (c DefaultClient) Execute(req Request) Response {

	apiUrlPattern := c.Server + "/%s"

	apiUrl := fmt.Sprintf(apiUrlPattern, req.ApiName())

	fmt.Println("apiUrl:", apiUrl)

	httpArgs := url.Values{}

	for k, v := range req.GetParams() {
		httpArgs.Add(k, v)
	}

	httpReq, err := http.NewRequest(req.GetHttpMethod(), apiUrl, strings.NewReader(httpArgs.Encode()))
	if err != nil {
		// handle error
		return nil
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	response := req.GetResponse()
	json.Unmarshal(body, response)

	return response
}
