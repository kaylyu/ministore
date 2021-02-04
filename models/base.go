package models

type IAPIRequest interface {
}

type APIRequest struct {
	AccessToken          string `json:"-"`
	ComponentAccessToken string `json:"-"`
}

//返回码
// https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/return/errcode.html
type APIResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
