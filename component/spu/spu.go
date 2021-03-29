package spu

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiAddSpu                      = "/product/spu/add"
	apiDelSpu                      = "/product/spu/del"
	apiGetSpu                      = "/product/spu/get"
	apiGetSpuList                  = "/product/spu/get_list"
	apiSearchSpuList               = "/product/spu/search"
	apiUpdateSpu                   = "/product/spu/update"
	apiUpSpuListing                = "/product/spu/listing"
	apiUpSpuDelisting              = "/product/spu/delisting"
	apiLimiteddiscountAdd          = "/product/limiteddiscount/add"
	apiLimiteddiscountList         = "/product/limiteddiscount/get_list"
	apiLimiteddiscountUpdateStatus = "/product/limiteddiscount/update_status"
)

//SPU
type Spu struct {
	core *component.Core
}

/*
创建SPU实例
*/
func NewSpu(core *component.Core) *Spu {
	return &Spu{core}
}

/*
添加商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/add_spu.html
*/
func (s *Spu) AddSpu(request *models.SpuAddRequest) (response models.SpuAddResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiAddSpu, request, &response)
	return
}

/*
删除商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/del_spu.html
*/
func (s *Spu) DelSpu(request *models.SpuDelRequest) (response models.SpuDelResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiDelSpu, request, &response)
	return
}

/*
获取商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu.html
*/
func (s *Spu) GetSpu(request *models.SpuGetRequest) (response models.SpuGetResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiGetSpu, request, &response)
	return
}

/*
获取商品列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu_list.html
*/
func (s *Spu) GetSpuList(request *models.SpuGetListRequest) (response models.SpuGetListResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiGetSpuList, request, &response)
	return
}

/*
搜索商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/search_spu.html
*/
func (s *Spu) SearchSpu(request *models.SpuSearchRequest) (response models.SpuSearchResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiSearchSpuList, request, &response)
	return
}

/*
更新商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu.html
*/
func (s *Spu) UpdateSpu(request *models.SpuUpdateRequest) (response models.SpuUpdateResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiUpdateSpu, request, &response)
	return
}

/*
上架商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_listing.html
*/
func (s *Spu) UpSpu(request *models.SpuUpRequest) (response models.SpuUpResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiUpSpuListing, request, &response)
	return
}

/*
下架商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_delisting.html
*/
func (s *Spu) DownSpu(request *models.SpuDownRequest) (response models.SpuDownResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiUpSpuDelisting, request, &response)
	return
}

/*
添加抢购任务
https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent/API/spu/add_limiteddiscount.html
*/
func (s *Spu) SpuLimitedDiscountAdd(request *models.SpuLimitedDiscountRequest) (response models.SpuLimitedDiscountResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiLimiteddiscountAdd, request, &response)
	return
}

/*
拉取抢购任务列表
https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent/API/spu/get_limiteddiscount_list.html
*/
func (s *Spu) SpuLimitedGetList(request *models.SpuLimitedListRequest) (response models.SpuLimitedListResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiLimiteddiscountList, request, &response)
	return
}

/*
修改抢购任务状态
https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent/API/spu/update_limiteddiscount_status.html
*/
func (s *Spu) SpuLimitedDiscountUpdate(request *models.SpuLimitedDiscountUpdateRequest) (response models.SpuLimitedDiscountUpdateResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiLimiteddiscountUpdateStatus, request, &response)
	return
}
