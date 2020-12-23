package models

type IAPIRequest interface {
	SetAccessToken(appId string)
}

type APIRequest struct {
	AccessToken string `json:"-"`
}

func (req *APIRequest) SetAccessToken(appId string) {
	req.AccessToken = appId
}

//返回码
// https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/return/errcode.html
type APIResponse struct {
	Errcode uint   `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
