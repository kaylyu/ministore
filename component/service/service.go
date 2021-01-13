// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package service 服务商接口
package service

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiCheckAuth           = "/product/service/check_auth"
	apiGetServiceList      = "/product/service/get_list"
	apiGetServiceOrderList = "/product/service/get_order_list"
)

//服务市场
type Service struct {
	core *component.Core
}

/*
创建服务市场实例
*/
func NewService(core *component.Core) *Service {
	return &Service{core}
}

/*
登录验证
See: https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/check_auth.html
*/
func (p *Service) CheckAuth(request *models.CheckAuthRequest, componentAccessToken string) (response models.CheckAuthResponse, body string, err error) {
	//设置
	request.SetComponentAccessToken(componentAccessToken)
	body, err = p.core.HttpPostJson(apiCheckAuth, request, &response)
	return
}

/*
获取用户购买的在有效期内的服务列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/get_service_list.html
*/
func (p *Service) GetServiceList(request *models.GetListRequest) (response models.GetListResponse, body string, err error) {
	body, err = p.core.HttpPostJson(apiGetServiceList, request, &response)
	return
}

/*
获取用户购买的服务列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/service/get_service_order_list.html
*/
func (p *Service) GetServiceOrderList(request *models.GetServiceOrderListRequest) (response models.GetServiceOrderListListResponse, body string, err error) {
	body, err = p.core.HttpPostJson(apiGetServiceOrderList, request, &response)
	return
}
