package util

import (
	"fmt"
	"github.com/kaylyu/ministore/context"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
微信 api 服务器地址
*/
var WxServerUrl = "https://api.weixin.qq.com"

//
type Client struct {
	ctx *context.Context
}

//获取实例
func NewHttpClient(ctx *context.Context) (client *Client) {
	return &Client{ctx}
}

//请求
func (c *Client) request(method string, reqUrl string, bComponentAccessToken bool, reqBody string, headers ...map[string]string) (body string, err error) {
	//判断使用的token
	var uri string
	if bComponentAccessToken {
		uri, err = c.applyComponentAccessToken(reqUrl)
	} else {
		uri, err = c.applyAccessToken(reqUrl)
	}
	if err != nil {
		return
	}
	if c.ctx.Logger != nil {
		c.ctx.Logger.Printf("POST %s, reqBody:%s", uri, reqBody)
	}

	//默认超时10S
	timeout := c.ctx.Config.Timeout
	if timeout == 0 {
		timeout = time.Duration(10) * time.Second
	}
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, WxServerUrl+uri, strings.NewReader(reqBody)) //set body
	if err != nil {
		return
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Charset", "utf-8;")
	req.Header.Set("Content-Type", "application/json;")
	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Set(k, v)
		}
	}
	//req
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status: %s", resp.Status)
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body = string(content)

	return
}

//POST
func (c *Client) HttpPostJson(path string, params string, bComponentAccessToken bool, headers ...map[string]string) (body string, err error) {
	//请求
	body, err = c.request("POST", path, bComponentAccessToken, params, headers...)

	return
}

/*
在请求地址上附加上 access_token
*/
func (c *Client) applyAccessToken(oldUrl string) (newUrl string, err error) {
	accessToken := c.ctx.Config.AccessToken
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
	componentAccessToken := c.ctx.Config.ComponentAccessToken
	if strings.Contains(oldUrl, "?") {
		newUrl = oldUrl + "&component_access_token=" + componentAccessToken
	} else {
		newUrl = oldUrl + "?component_access_token=" + componentAccessToken
	}
	return
}

/*
筛查微信 api 服务器响应，判断以下错误：
- http 状态码 不为 200
- 接口响应错误码 errcode 不为 0
*/
