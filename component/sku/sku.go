package sku

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiAddSku         = "/product/sku/add"
	apiBatchAddSku    = "/product/sku/batch_add"
	apiDelSku         = "/product/sku/del"
	apiGetSku         = "/product/sku/get"
	apiGetListSku     = "/product/sku/get_list"
	apiUpdateSku      = "/product/sku/update"
	apiUpdateSkuPrice = "/product/sku/update_price"
	apiUpdateStock    = "/product/sku/stock/update"
	apiGetStock       = "/product/sku/stock/get"
)

//Sku
type Sku struct {
	core *component.Core
}

/*
创建Sku实例
*/
func NewSku(core *component.Core) *Sku {
	return &Sku{core}
}

/*
添加SKU
https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/add_sku.html
*/
func (s *Sku) AddSku(request *models.SkuAddRequest) (response models.SkuAddResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiAddSku, request, &response)
	return
}

/*
批量添加SKU
https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/batch_add_sku.html
*/
func (s *Sku) BatchAddSku(request *models.SkuBatchAddRequest) (response models.SkuBatchAddResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiBatchAddSku, request, &response)
	return
}

/*
删除SKU
https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/del_sku.html
*/
func (s *Sku) DelSku(request *models.SkuDelRequest) (response models.SkuDelResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiDelSku, request, &response)
	return
}

/*
获取SKU信息
https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/get_sku.html
*/
func (s *Sku) GetSku(request *models.SkuGetRequest) (response models.SkuGetResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiGetSku, request, &response)
	return
}

/*
批量获取SKU信息
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/sku/get_sku_list.html
*/
func (s *Sku) GetListSku(request *models.SkuGetListRequest) (response models.SkuGetListResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiGetListSku, request, &response)
	return
}

/*
更新SKU
https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/up_sku.html
*/
func (s *Sku) UpdateSku(request *models.SkuUpdateRequest) (response models.SkuUpdateResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiUpdateSku, request, &response)
	return
}

/*
更新SKU价格
https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/up_sku_price.html
*/
func (s *Sku) UpdateSkuPrice(request *models.SkuUpdatePriceRequest) (response models.SkuUpdatePriceResponse, err error) {
	err = s.core.HttpPostJson(apiUpdateSkuPrice, request, &response)
	return
}

/*
更新库存
https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/up_stock.html
*/
func (s *Sku) UpdateStock(request *models.SkuUpdateStockRequest) (response models.SkuUpdateStockResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiUpdateStock, request, &response)
	return
}

/*
获取库存
https://developers.weixin.qq.com/doc/MiniStore/minishopopencomponent/API/sku/get_stock.html
*/
func (s *Sku) GetStock(request *models.SkuGetStockRequest) (response models.SkuGetStockResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiGetStock, request, &response)
	return
}
