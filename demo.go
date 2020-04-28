package main

import (
	"fmt"
	"reflect"
)

func main() {

	//testGetAccessToken2()

	testGetCallbackIpRequest2()

}

/**
获取微信服务器列表
https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
*/
func testGetCallbackIpRequest() {
	req := NewGetCallbackIpRequest("32_QR0AbpGf34moAA6csEetsKbJjQ4HYvgdF_C0wYp81E-WvCeA1P-vlwpmVvaFQtv3hBKXzWLj_OEbarI88Uagn-A3RxIAhze5bg5ZRP-WXynTU8IUWZyODRZ17W50cJGvGAuiBZbgMPQKJ983AONcACAPDC")
	client := NewDefaultClient("https://api.weixin.qq.com/cgi-bin")
	response := client.Execute(req)
	fmt.Println(response)
}

/**
获取微信服务器列表
https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
*/
func testGetCallbackIpRequest2() {

	token := "32_l72IYKSbAaBX1Go-bMnL_JInTxiBuJiIJYELx2q5FKBa0aMOGFNO556OgDcMtv5M8yVLMC0nhCGVvGCIGI88CauST29rzcG9HbFiJ2dvYyc95ZiUkOXYWbfbLAFi6ikia5sTeVcIJnRB0bnXLLXgAFACCO"
	req := NewGetCallbackIpRequest(token)
	//client := NewDefaultClient("https://api.weixin.qq.com/cgi-bin")
	response := req.Execute()
	fmt.Printf("ip list:%v", response.IpList)
	fmt.Println(response.IsSuccess())
}

/**
通过client实现http请求
*/
func testGetAccessToken1() {
	req := NewGetAccessTokenRequest("wx41a8a8b76d194515", "bd6e5590cc1eeb6180d980628cf82b55")
	fmt.Print(req.Appid)
	client := NewDefaultClient("https://api.weixin.qq.com/cgi-bin")
	response := client.Execute(req)
	println(reflect.TypeOf(response))
	println(response)
	println(response.IsSuccess())
}

/**
每个请求对象实现自己的http请求
*/
func testGetAccessToken2() {
	req := NewGetAccessTokenRequest("wx41a8a8b76d194515", "bd6e5590cc1eeb6180d980628cf82b55")
	fmt.Print(req.Appid)
	response := req.Execute()
	println(reflect.TypeOf(response))
	println("token:" + response.Access_token)
	println(response.IsSuccess())

}
