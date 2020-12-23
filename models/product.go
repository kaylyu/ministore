package models

//类目详情请求参数说明
type CategoryRequest struct {
	APIRequest
	FCatId uint64 `json:"f_cat_id"` //父类目ID，可先填0获取根部类目
}

//类目详情响应参数说明
type CategoryResponse struct {
	APIResponse
	Data []*CategoryList `json:"cat_list"`
}

type CategoryList struct {
	CatId  uint64 `json:"cat_id"`   //类目ID
	FCatId uint64 `json:"f_cat_id"` //类目父ID
	Name   string `json:"name"`     //类目名称
}

//品牌列表请求参数说明
type BrandRequest struct {
	APIRequest
}

//品牌列表响应参数说明
type BrandResponse struct {
	APIResponse
	Data []*Brands `json:"brands"`
}

type Brands struct {
	FirstCatId  uint64     `json:"first_cat_id"`  //第一级类目
	SecondCatId uint64     `json:"second_cat_id"` //第二级类目
	ThirdCatId  uint64     `json:"third_cat_id"`  //第三级类目
	BrandInfo   *BrandInfo `json:"brand_info"`    //品牌
}
type BrandInfo struct {
	BrandId   uint64 `json:"brand_id"`   //品牌ID
	BrandName string `json:"brand_name"` //品牌名称
}

//运费模板请求参数说明
type FreightTemplateRequest struct {
	APIRequest
}

//运费模板响应参数说明
type FreightTemplateResponse struct {
	APIResponse
	Data []*TemplateList `json:"template_list"`
}

type TemplateList struct {
	TemplateId    uint64        `json:"template_id"`    //模板ID
	Name          string        `json:"name"`           //模板名称
	ValuationType ValuationType `json:"valuation_type"` //计费类型
}

//店铺的商品分类请求参数说明
type ShopCatRequest struct {
	APIRequest
}

//店铺的商品分类响应参数说明
type ShopCatResponse struct {
	APIResponse
	Data []*ShopCatList `json:"shopcat_list"`
}

type ShopCatList struct {
	ShopcatId   uint64 `json:"shopcat_id"`   //分类ID
	ShopcatName string `json:"shopcat_name"` //分类名称
	FShopcatId  uint64 `json:"f_shopcat_id"` //父分类ID
	CatLevel    uint64 `json:"cat_level"`    //分类层级
}
