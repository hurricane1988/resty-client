package config

import "github.com/go-resty/resty/v2"

// 定义resty客户端结构体
type restyClient struct {
	Client *resty.Client
}

// RestClient 定义resty全局类
var RestClient restyClient

// InitRestClient 初始化RestClient
func (r *restyClient) InitRestClient() {
	client := resty.New()
	r.Client = client
}
