package main

import (
	"fmt"
	"reflect"
)

func main() {

	//testGetAccessToken2()

	//testGetCallbackIpRequest2()

	//testCreateMenuRequest()

	//testGetMenuRequest()
	//testDeleteMenuRequest()

	//testCreateTagRequest()
	//testGetTagsRequest()
	//testUpdateTagRequest()
	//testDeleteTagRequest()

	testGetTagFansRequest()

}

func testGetTagFansRequest() {
	token := "32_BcYNq8KPUi5tyOPpY2eJtYiWZJCWDz9uXwHAeh-jwL0HvGYclWt29Dho908zB6O4twOS5YKPV4ZujKMoxXqbl4w9E0r0ANb1leeHfzuqhTyJUmI-lemdE7mDHR53uhwI-xup-0ilxGP9eK-pOZDeAFATEI"

	req := NewGetTagFansRequest(token, 100, "")

	response := req.Execute()

	fmt.Println(response.IsSuccess())
	fmt.Println(response.Errcode)
	fmt.Println(response.Errmsg)
	fmt.Println(response.TagFansData)
}

func testDeleteTagRequest() {
	token := "32_BcYNq8KPUi5tyOPpY2eJtYiWZJCWDz9uXwHAeh-jwL0HvGYclWt29Dho908zB6O4twOS5YKPV4ZujKMoxXqbl4w9E0r0ANb1leeHfzuqhTyJUmI-lemdE7mDHR53uhwI-xup-0ilxGP9eK-pOZDeAFATEI"

	req := NewDeleteTagRequest(token, 104)

	response := req.Execute()

	fmt.Println(response.IsSuccess())
	fmt.Println(response.Errcode)
	fmt.Println(response.Errmsg)
}

func testUpdateTagRequest() {
	// 100
	token := "32_BcYNq8KPUi5tyOPpY2eJtYiWZJCWDz9uXwHAeh-jwL0HvGYclWt29Dho908zB6O4twOS5YKPV4ZujKMoxXqbl4w9E0r0ANb1leeHfzuqhTyJUmI-lemdE7mDHR53uhwI-xup-0ilxGP9eK-pOZDeAFATEI"

	req := NewUpdateTagRequest(token, 100, "test1122")

	response := req.Execute()

	fmt.Println(response.IsSuccess())
	fmt.Println(response.Errcode)
	fmt.Println(response.Errmsg)
}

func testGetTagsRequest() {
	token := "32_BcYNq8KPUi5tyOPpY2eJtYiWZJCWDz9uXwHAeh-jwL0HvGYclWt29Dho908zB6O4twOS5YKPV4ZujKMoxXqbl4w9E0r0ANb1leeHfzuqhTyJUmI-lemdE7mDHR53uhwI-xup-0ilxGP9eK-pOZDeAFATEI"

	req := NewGetTagsRequest(token)

	response := req.Execute()

	fmt.Println(response.IsSuccess())
	fmt.Println(response.Errcode)
	fmt.Println(response.Errmsg)
	fmt.Println(len(response.Tags))

	for e := range response.Tags {
		fmt.Println(response.Tags[e])
	}

}

func testCreateTagRequest() {
	token := "32_looq8breXAKpt_YAtzeIAKMYkHY_VMGf2DEcgJYyjFthpklujNsCANHukp0HJ64tA0dloS7wQGy16tl3fIEATqM1h7-8zr0Ccli-HuUnX80eDYp7SXOq9km7cP4Eb7-WhdbJAcsfvZbyMSBfJDLeABATNV"

	req := NewCreateTagRequest(token, "shanghai")

	response := req.Execute()

	fmt.Println(response.IsSuccess())
}

/*

删除菜单
*/
func testDeleteMenuRequest() {
	token := "32_C5cblJWbfXzYegq7Y5NvDhRj3UO0gck0SrZaQv5aOeJ0lOUmO9u-DhSzcXdETyagck0ycgJJr4U0EtMaWJrIArlKxMcYYY0XIjdwIOUaK-PtJGz_MZafQFeTrC7RANeFkqwkOSgwu3w3UxgmRCKaAJAWHP"

	req := NewDeleteMenuRequest(token)

	response := req.Execute()

	fmt.Println(response.IsSuccess())
	fmt.Println(response.Errcode)
}

func testGetMenuRequest() {
	token := "32_uoKpK1fnJmVYNVYXOHDzgthJihH8FMj_e63QWl3LfoeVKcRAmLBCo5NuL-0WjPwE4A6IqLtPHOVnLLvNt3tyOqEnxGz4r6w5jUBaun3r1CnuR__VdZL5h9OmO-hdCsKfjm14X96HD77JqaryBLMcAIANXY"

	req := NewGetMenuRequest(token)

	response := req.Execute()

	fmt.Println(response.IsSuccess())
	fmt.Println(response.Is_menu_open)
	fmt.Println(len(response.Selfmenu_info.Button))
	fmt.Println(response.Selfmenu_info.Button)
}

/*
创建菜单
*/
func testCreateMenuRequest() {

	token := "32_looq8breXAKpt_YAtzeIAKMYkHY_VMGf2DEcgJYyjFthpklujNsCANHukp0HJ64tA0dloS7wQGy16tl3fIEATqM1h7-8zr0Ccli-HuUnX80eDYp7SXOq9km7cP4Eb7-WhdbJAcsfvZbyMSBfJDLeABATNV"

	//menu1 := Button{"view", "test1", "http://www.baidu.com", "", []Button{}}
	menu1 := Button{Menu_type: "view", Name: "test", Url: "http://www.baidu.com"}
	buttons := []Button{menu1}

	req := NewCreateMenuRequest(token, buttons)

	response := req.Execute()

	fmt.Println(response.IsSuccess())
}

/*
获取微信服务器列表
https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
*/
func testGetCallbackIpRequest() {
	req := NewGetCallbackIpRequest("32_QR0AbpGf34moAA6csEetsKbJjQ4HYvgdF_C0wYp81E-WvCeA1P-vlwpmVvaFQtv3hBKXzWLj_OEbarI88Uagn-A3RxIAhze5bg5ZRP-WXynTU8IUWZyODRZ17W50cJGvGAuiBZbgMPQKJ983AONcACAPDC")
	client := NewDefaultClient("https://api.weixin.qq.com/cgi-bin")
	response := client.Execute(req)
	fmt.Println(response)
}

/*
获取微信服务器列表
https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
*/
func testGetCallbackIpRequest2() {

	token := "32_l72IYKSbAaBX1Go-bMnL_JInTxiBuJiIJYELx2q5FKBa0aMOGFNO556OgDcMtv5M8yVLMC0nhCGVvGCIGI88CauST29rzcG9HbFiJ2dvYyc95ZiUkOXYWbfbLAFi6ikia5sTeVcIJnRB0bnXLLXgAFACCO"
	req := NewGetCallbackIpRequest(token)
	response := req.Execute()
	fmt.Printf("ip list:%v", response.IpList)
	fmt.Println(response.IsSuccess())
}

/*
通过client实现http请求
*/
func testGetAccessToken1() {
	req := NewGetAccessTokenRequest("wxe58afcd99f7a997e", "5dcf8eac1e99e983fc58e42376ab0267")
	fmt.Print(req.Appid)
	client := NewDefaultClient("https://api.weixin.qq.com/cgi-bin")
	response := client.Execute(req)
	println(reflect.TypeOf(response))
	println(response)
	println(response.IsSuccess())
}

/*
每个请求对象实现自己的http请求
*/
func testGetAccessToken2() {
	req := NewGetAccessTokenRequest("wxe58afcd99f7a997e", "5dcf8eac1e99e983fc58e42376ab0267")
	fmt.Print(req.Appid)
	response := req.Execute()
	println(reflect.TypeOf(response))
	println("token:" + response.Access_token)
	println(response.IsSuccess())

}
