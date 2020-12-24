package models

//添加商品
type SpuAddRequest struct {
	APIRequest
	OutProductId string       `json:"out_product_id"` //商家自定义商品ID，与product_id二选一
	Title        string       `json:"title"`          //标题
	SubTitle     string       `json:"sub_title"`      //副标题
	HeadImg      []string     `json:"head_img"`       //主图,多张,列表
	DescInfo     DescInfo     `json:"desc_info"`      //商品详情，图文(目前只支持图片)
	BrandId      uint64       `json:"brand_id"`       //品牌ID，商家需要申请品牌并通过获取品牌接口brand/get获取，如果是无品牌请填2100000000
	Cats         []*Cat       `json:"cats"`           //类目
	Attrs        []*Attr      `json:"attrs"`          //属性
	Model        string       `json:"model"`          //商品型号
	ExpressInfo  *ExpressInfo `json:"express_info"`   //运费
	Skus         []*Sku       `json:"skus"`           //该 skus 列表非必填，可另行通过 BatchAddSKU 添加 SKU
}

//商品详情，图文(目前只支持图片)
type DescInfo struct {
	Imgs []string `json:"imgs"` //商品详情，图文(目前只支持图片)
}

//类目
type Cat struct {
	CatId uint64 `json:"cat_id"` //类目ID，如果brand_id=2100000000，需要先通过获取类目接口category/get拿到可用的cat_id；如果brand_id!=2100000000，则这里的cat_id需要与获取品牌接口brand/get中的1,2,3级类目一一对应
	Level uint64 `json:"level"`  //类目层级
}

//属性
type Attr struct {
	AttrKey   string `json:"attr_key"`   //属性键key（属性自定义用）
	AttrValue string `json:"attr_value"` //属性值（属性自定义用）
}

//运费模板
type ExpressInfo struct {
	TemplateId uint64 `json:"template_id"` //运费模板ID（先通过获取运费模板接口delivery/get_freight_template拿到）
}

//Sku
type Sku struct {
	OutProductId string     `json:"out_product_id"`   //商家自定义商品ID
	OutSkuId     string     `json:"out_sku_id"`       //商家自定义skuID
	SkuId        uint64     `json:"sku_id,omitempty"` //商家自定义skuID
	ThumbImg     string     `json:"thumb_img"`        //sku小图
	SalePrice    uint64     `json:"sale_price"`       //售卖价格,以分为单位
	MarketPrice  uint64     `json:"market_price"`     //市场价格,以分为单位
	StockNum     uint64     `json:"stock_num"`        //库存
	Barcode      string     `json:"barcode"`          //	条形码
	SkuCode      string     `json:"sku_code"`         //商品编码
	SkuAttrs     []*SkuAttr `json:"sku_attrs"`        //属性
}

//属性
type SkuAttr struct {
	AttrKey   string `json:"attr_key"`   //属性键key（属性自定义用）
	AttrValue string `json:"attr_value"` //属性值（属性自定义用）
}

//添加商品
type SpuAddResponse struct {
	APIResponse
	Data *SpuAddData `json:"data"`
}

//添加商品Data
type SpuAddData struct {
	ProductId    uint64 `json:"product_id"`     //小商店内部商品ID
	OutProductId string `json:"out_product_id"` //商家自定义商品ID
	CreateTime   string `json:"create_time"`    //创建时间
}

//删除商品
type SpuDelRequest struct {
	APIRequest
	ProductId    uint64 `json:"product_id"`     //小商店内部商品ID，与out_product_id二选一
	OutProductId string `json:"out_product_id"` //商家自定义商品ID，与product_id二选一
}

//删除商品
type SpuDelResponse struct {
	APIResponse
}

//获取商品
type SpuGetRequest struct {
	APIRequest
	ProductId    uint64      `json:"product_id"`     //小商店内部商品ID，与out_product_id二选一
	OutProductId string      `json:"out_product_id"` //商家自定义商品ID，与product_id二选一
	NeedEditSpu  NeedEditSpu `json:"need_edit_spu"`  //默认0:获取线上数据, 1:获取草稿数据
}

//获取商品
type SpuGetResponse struct {
	APIResponse
	Data *SpuGetData `json:"data"`
}

//商品获取数据
type SpuGetData struct {
	Spu []*SpuGetDataSpu `json:"spu"`
}
type SpuGetDataSpu struct {
	ProductId    uint64         `json:"product_id"`     //小商店内部商品ID，与out_product_id二选一
	OutProductId string         `json:"out_product_id"` //商家自定义商品ID，与product_id二选一
	Title        string         `json:"title"`          //标题
	SubTitle     string         `json:"sub_title"`      //副标题
	HeadImg      []string       `json:"head_img"`       //主图,多张,列表
	DescInfo     *DescInfo      `json:"desc_info"`      //商品详情，图文(目前只支持图片)
	BrandId      uint64         `json:"brand_id"`       //品牌ID，商家需要申请品牌并通过获取品牌接口brand/get获取，如果是无品牌请填2100000000
	Cats         []*Cat         `json:"cats"`           //类目
	Attrs        []*Attr        `json:"attrs"`          //属性
	Model        string         `json:"model"`          //商品型号
	ExpressInfo  *ExpressInfo   `json:"express_info"`   //运费
	Skus         []*Sku         `json:"skus"`           //该skus 列表非必填，可另行通过 BatchAddSKU 添加 SKU
	Status       SpuStatus      `json:"status"`         //商品线上状态
	EditStatus   SpuEditStatus  `json:"edit_status"`    //商品草稿状态
	MinPrice     uint64         `json:"min_price"`      //商品SKU最小价格（单位：分）
	ShopCatLists []*ShopCatList `json:"shop_cat_list"`  //分类ID，对应信息通过 get_shop_cat 接口拿到
}

//获取商品列表
type SpuGetListRequest struct {
	APIRequest
	NeedEditSpu NeedEditSpu `json:"need_edit_spu"` //默认0:获取线上数据, 1:获取草稿数据
	Status      SpuStatus   `json:"status"`        //商品状态
	Page        uint64      `json:"page"`          //页码
	PageSize    uint64      `json:"page_size"`     //每页个数
}

//获取商品列表
type SpuGetListResponse struct {
	APIResponse
	Data []*SpuGetListDataSpu `json:"spus"`
}

//商品列表获取数据
type SpuGetListDataSpu struct {
	ProductId    uint64         `json:"product_id"`        //小商店内部商品ID，与out_product_id二选一
	OutProductId string         `json:"out_product_id"`    //商家自定义商品ID，与product_id二选一
	Title        string         `json:"title"`             //标题
	SubTitle     string         `json:"sub_title"`         //副标题
	HeadImg      []string       `json:"head_img"`          //主图,多张,列表
	DescInfo     *DescInfo      `json:"desc_info"`         //商品详情，图文(目前只支持图片)
	BrandId      uint64         `json:"brand_id"`          //品牌ID，商家需要申请品牌并通过获取品牌接口brand/get获取，如果是无品牌请填2100000000
	Cats         []*Cat         `json:"cats,omitempty"`    //类目
	Attrs        []*Attr        `json:"attrs,omitempty"`   //属性
	Model        string         `json:"model"`             //商品型号
	ExpressInfo  *ExpressInfo   `json:"express_info"`      //运费
	Skus         []*Sku         `json:"skus"`              //该skus 列表非必填，可另行通过 BatchAddSKU 添加 SKU
	Status       SpuStatus      `json:"status"`            //商品线上状态
	EditStatus   SpuEditStatus  `json:"edit_status"`       //商品草稿状态
	MinPrice     uint64         `json:"min_price"`         //商品SKU最小价格（单位：分）
	ShopCatLists []*ShopCatList `json:"shopcat,omitempty"` //分类ID，对应信息通过 get_shop_cat 接口拿到
}

//搜索商品列表
type SpuSearchRequest struct {
	APIRequest
	NeedEditSpu NeedEditSpu `json:"need_edit_spu"` //默认0:获取线上数据, 1:获取草稿数据
	Status      SpuStatus   `json:"status"`        //商品状态
	Page        uint64      `json:"page"`          //页码
	PageSize    uint64      `json:"page_size"`     //每页个数
	Keyword     string      `json:"keyword"`       //关键字
}

//搜索商品列表
type SpuSearchResponse struct {
	APIResponse
	Data *SpuSearchDataSpu `json:"spus"`
}

//搜索商品列表获取数据
type SpuSearchDataSpu struct {
	ProductId    uint64         `json:"product_id"`     //小商店内部商品ID，与out_product_id二选一
	OutProductId string         `json:"out_product_id"` //商家自定义商品ID，与product_id二选一
	Title        string         `json:"title"`          //标题
	SubTitle     string         `json:"sub_title"`      //副标题
	HeadImg      []string       `json:"head_img"`       //主图,多张,列表
	DescInfo     *DescInfo      `json:"desc_info"`      //商品详情，图文(目前只支持图片)
	BrandId      uint64         `json:"brand_id"`       //品牌ID，商家需要申请品牌并通过获取品牌接口brand/get获取，如果是无品牌请填2100000000
	Cats         []*Cat         `json:"cats"`           //类目
	Attrs        []*Attr        `json:"attrs"`          //属性
	Model        string         `json:"model"`          //商品型号
	ExpressInfo  *ExpressInfo   `json:"express_info"`   //运费
	Skus         []*Sku         `json:"skus"`           //该skus 列表非必填，可另行通过 BatchAddSKU 添加 SKU
	Status       SpuStatus      `json:"status"`         //商品线上状态
	EditStatus   SpuEditStatus  `json:"edit_status"`    //商品草稿状态
	MinPrice     uint64         `json:"min_price"`      //商品SKU最小价格（单位：分）
	ShopCatLists []*ShopCatList `json:"shop_cat_list"`  //分类ID，对应信息通过 get_shop_cat 接口拿到
}

//更新商品
type SpuUpdateRequest struct {
	APIRequest
	OutProductId string       `json:"out_product_id"` //商家自定义商品ID，与product_id二选一
	Title        string       `json:"title"`          //标题
	SubTitle     string       `json:"sub_title"`      //副标题
	HeadImg      []string     `json:"head_img"`       //主图,多张,列表
	DescInfo     DescInfo     `json:"desc_info"`      //商品详情，图文(目前只支持图片)
	BrandId      uint64       `json:"brand_id"`       //品牌ID，商家需要申请品牌并通过获取品牌接口brand/get获取，如果是无品牌请填2100000000
	Cats         []*Cat       `json:"cats"`           //类目
	Attrs        []*Attr      `json:"attrs"`          //属性
	Model        string       `json:"model"`          //商品型号
	ExpressInfo  *ExpressInfo `json:"express_info"`   //运费
}

//更新商品
type SpuUpdateResponse struct {
	APIResponse
	Data *SpuUpdateData `json:"spus"`
}

//更新商品Data
type SpuUpdateData struct {
	ProductId    uint64 `json:"product_id"`     //小商店内部商品ID
	OutProductId string `json:"out_product_id"` //商家自定义商品ID
	UpdateTime   string `json:"update_time"`    //更新时间
}

//上架商品
type SpuUpRequest struct {
	APIRequest
	ProductId    uint64 `json:"product_id"`     //小商店内部商品ID
	OutProductId string `json:"out_product_id"` //商家自定义商品ID
}

//上架商品
type SpuUpResponse struct {
	APIResponse
}

//下架商品
type SpuDownRequest struct {
	APIRequest
	ProductId    uint64 `json:"product_id"`     //小商店内部商品ID
	OutProductId string `json:"out_product_id"` //商家自定义商品ID
}

//下架商品
type SpuDownResponse struct {
	APIResponse
}
