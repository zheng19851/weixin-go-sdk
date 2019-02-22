package main

func main() {

	req := NewGetAccessTokenRequest("wx41a8a8b76d194515", "")

	res := req.Execute()

	println(res.Access_token)
	println(res.Expires_in)

}
