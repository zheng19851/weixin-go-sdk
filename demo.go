package main

import "reflect"

func main() {

	req := NewGetAccessTokenRequest("wx41a8a8b76d194515", "bd6e5590cc1eeb6180d980628cf82b55")

	client := NewDefaultClient("https://api.weixin.qq.com/cgi-bin")

	response := client.Execute(req)

	println(reflect.TypeOf(response))

	println(response)

	println(response.IsSuccess())

}
