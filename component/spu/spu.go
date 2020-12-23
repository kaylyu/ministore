package spu

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiAddSpu         = "/product/spu/add"
	apiDelSpu         = "/product/spu/del"
	apiGetSpu         = "/product/spu/get"
	apiGetSpuList     = "/product/spu/get_list"
	apiSearchSpuList  = "/product/spu/search"
	apiUpdateSpu      = "/product/spu/update"
	apiUpSpuListing   = "/product/spu/listing"
	apiUpSpuDelisting = "/product/spu/delisting"
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
func (s *Spu) AddSpu(request *models.SpuAddRequest) (response models.SpuAddResponse, err error) {
	err = s.core.HttpPostJson(apiAddSpu, request, &response)
	return
}

/*
删除商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/del_spu.html
*/
func (s *Spu) DelSpu(request *models.SpuDelRequest) (response models.SpuDelResponse, err error) {
	err = s.core.HttpPostJson(apiDelSpu, request, &response)
	return
}

/*
获取商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu.html
*/
func (s *Spu) GetSpu(request *models.SpuGetRequest) (response models.SpuGetResponse, err error) {
	err = s.core.HttpPostJson(apiGetSpu, request, &response)
	return
}

/*
获取商品列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/get_spu_list.html
*/
func (s *Spu) GetSpuList(request *models.SpuGetListRequest) (response models.SpuGetListResponse, err error) {
	err = s.core.HttpPostJson(apiGetSpuList, request, &response)
	return
}

/*
搜索商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/search_spu.html
*/
func (s *Spu) SearchSpu(request *models.SpuSearchRequest) (response models.SpuSearchResponse, err error) {
	err = s.core.HttpPostJson(apiSearchSpuList, request, &response)
	return
}

/*
更新商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu.html
*/
func (s *Spu) UpdateSpu(request *models.SpuUpdateRequest) (response models.SpuUpdateResponse, err error) {
	err = s.core.HttpPostJson(apiUpdateSpu, request, &response)
	return
}

/*
上架商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_listing.html
*/
func (s *Spu) UpSpu(request *models.SpuUpRequest) (response models.SpuUpResponse, err error) {
	err = s.core.HttpPostJson(apiUpSpuListing, request, &response)
	return
}

/*
下架商品
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/spu/up_spu_delisting.html
*/
func (s *Spu) DownSpu(request *models.SpuDownRequest) (response models.SpuDownResponse, err error) {
	err = s.core.HttpPostJson(apiUpSpuDelisting, request, &response)
	return
}
