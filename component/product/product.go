package product

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiGetCatList         = "/product/category/get"
	apiGetBrand           = "/product/brand/get"
	apiGetFreightTemplate = "/product/delivery/get_freight_template"
	apiGetShopCat         = "/product/store/get_shopcat"
)

//商品
type Product struct {
	core *component.Core
}

/*
创建商品实例
*/
func NewProduct(core *component.Core) *Product {
	return &Product{core}
}

/*
获取类目详情
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_cat_list.html
*/
func (p *Product) GetCategory(request *models.CategoryRequest) (response models.CategoryResponse, body string, err error) {
	body, err = p.core.HttpPostJson(apiGetCatList, request, &response)
	return
}

/*
获取品牌列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_brand.html
*/
func (p *Product) GetBrand(request *models.BrandRequest) (response models.BrandResponse, body string, err error) {
	body, err = p.core.HttpPostJson(apiGetBrand, request, &response)
	return
}

/*
获取运费模板
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/cat/get_freight_template.html
*/
func (p *Product) GetFreightTemplate(request *models.FreightTemplateRequest) (response models.FreightTemplateResponse, body string, err error) {
	body, err = p.core.HttpPostJson(apiGetFreightTemplate, request, &response)
	return
}

/**
获取店铺的商品分类
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/store/get_shopcat.html
*/
func (p *Product) GetCats(request *models.ShopCatRequest) (response models.ShopCatResponse, body string, err error) {
	body, err = p.core.HttpPostJson(apiGetShopCat, request, &response)
	return
}
