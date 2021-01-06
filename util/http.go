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
	req := ghttp.NewClient()
	req.SetTimeOut(timeout)

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
func (c *Client) HttpPostJson(path string, params string, bComponentAccessToken bool, headers ...map[string]string) (body string, err error) {
	//请求
	body, err = c.request("POST", path, bComponentAccessToken, params, headers...)

	return
}

//上传图片
func (c *Client) Upload(path, filename string) (body string, err error) {
	uri, err := c.applyAccessToken(path)
	if err != nil {
		return
	}
	//上传文件
	rsp, err := ghttp.Post(WxServerUrl+uri, "upload-file=@media:"+filename)
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
