package main

import (
	"fmt"
	"resty-client/config"
	"resty-client/pkg/get"
)

func main() {
	config.RestClient.InitRestClient()
	resp, trace := get.GetClient("https://www.baidu.com", true)
	fmt.Println(string(resp))
	fmt.Println(string(trace))
}
