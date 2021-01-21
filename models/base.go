package models

type IAPIRequest interface {
	SetAccessToken(accessToken string)
	GetAccessToken() string
	SetComponentAccessToken(componentAccessToken string)
	GetComponentAccessToken() string
}

type APIRequest struct {
	AccessToken          string `json:"-"`
	ComponentAccessToken string `json:"-"`
}

func (req *APIRequest) SetAccessToken(accessToken string) {
	req.AccessToken = accessToken
}

func (req *APIRequest) GetAccessToken() string {
	return req.AccessToken
}

func (req *APIRequest) SetComponentAccessToken(componentAccessToken string) {
	req.ComponentAccessToken = componentAccessToken
}
func (req *APIRequest) GetComponentAccessToken() string {
	return req.ComponentAccessToken
}

//返回码
// https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/return/errcode.html
type APIResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
