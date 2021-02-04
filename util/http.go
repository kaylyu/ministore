package util

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/kaylyu/ministore/context"
	"io/ioutil"
	"strings"
	"time"
)

/*
微信 api 服务器地址
*/
var WxServerUrl = "https://api.weixin.qq.com"

//
type Client struct {
	*context.Context
}

//获取实例
func NewHttpClient(ctx *context.Context) (client *Client) {
	return &Client{ctx}
}

//请求
func (c *Client) request(method string, reqUrl string, reqBody string, headers ...map[string]string) (body string, err error) {
	//判断使用的token
	var uri string
	if len(c.Config.ComponentAccessToken) > 0 {
		uri, err = c.applyComponentAccessToken(reqUrl)
	} else {
		uri, err = c.applyAccessToken(reqUrl)
	}
	if err != nil {
		return
	}
	if c.Logger != nil {
		c.Logger.Printf("POST %s, reqBody:%s", uri, reqBody)
	}

	//默认超时10S
	timeout := c.Config.Timeout
	if timeout == 0 {
		timeout = time.Duration(10) * time.Second
	}
	req := ghttp.NewClient()
	req.Timeout(timeout)

	req.SetHeader("Accept", "*/*")
	req.SetHeader("Accept-Charset", "utf-8;")
	req.SetHeader("Content-Type", "application/json;")
	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.SetHeader(k, v)
		}
	}

	body = req.RequestContent(method, WxServerUrl+uri, reqBody)

	return
}

//POST
func (c *Client) HttpPostJson(path string, params string, headers ...map[string]string) (body string, err error) {
	//请求
	body, err = c.request("POST", path, params, headers...)

	return
}

//上传图片
func (c *Client) Upload(path, filename string) (body string, err error) {
	uri, err := c.applyAccessToken(path)
	if err != nil {
		return
	}
	//上传文件
	rsp, err := ghttp.NewClient().Post(WxServerUrl+uri, "media=@file:"+filename)
	if err != nil {
		return
	}
	defer rsp.Close()
	bs, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}

	body = string(bs)
	return
}

/*
在请求地址上附加上 access_token
*/
func (c *Client) applyAccessToken(oldUrl string) (newUrl string, err error) {
	accessToken := c.Config.AccessToken
	if strings.Contains(oldUrl, "?") {
		newUrl = oldUrl + "&access_token=" + accessToken
	} else {
		newUrl = oldUrl + "?access_token=" + accessToken
	}
	return
}

/*
在请求地址上附加上 component_access_token
*/
func (c *Client) applyComponentAccessToken(oldUrl string) (newUrl string, err error) {
	componentAccessToken := c.Config.ComponentAccessToken
	if strings.Contains(oldUrl, "?") {
		newUrl = oldUrl + "&component_access_token=" + componentAccessToken
	} else {
		newUrl = oldUrl + "?component_access_token=" + componentAccessToken
	}
	return
}
