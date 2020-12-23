package models

//商品信息
type SkuData struct {
	ProductId    uint64     `json:"product_id"`     //小商店内部商品ID
	OutProductId string     `json:"out_product_id"` //商家自定义商品ID
	OutSkuId     string     `json:"out_sku_id"`     //sku_id
	SkuId        uint64     `json:"sku_id"`         //sku_id
	ThumbImg     string     `json:"thumb_img"`      //sku小图
	SalePrice    uint64     `json:"sale_price"`     //售卖价格,以分为单位
	MarketPrice  uint64     `json:"market_price"`   //	市场价格,以分为单位
	StockNum     uint64     `json:"stock_num"`      //库存
	Barcode      string     `json:"barcode"`        //条形码
	SkuCode      string     `json:"sku_code"`       //商品编码
	SkuAttrs     []*SkuAttr `json:"sku_attr"`       //属性
}

//添加SKU
type SkuAddRequest struct {
	APIRequest
	SkuData
}
type SkuAddResponse struct {
	APIResponse
	Data *SkuAddResponseData `json:"data"`
}
type SkuAddResponseData struct {
	SkuId      uint64 `json:"sku_id"`      //skuID
	CreateTime string `json:"create_time"` //创建时间
}

//批量添加SKU
type SkuBatchAddRequest struct {
	APIRequest
	Skus []*SkuData `json:"skus"` //请求数据
}
type SkuBatchAddResponse struct {
	APIResponse
	Data []*SkuBatchAddData `json:"data"`
}
type SkuBatchAddData struct {
	SkuId      uint64 `json:"sku_id"`      //skuID
	OutSkuId   string `json:"out_sku_id"`  //商家自定义商品ID
	CreateTime string `json:"create_time"` //创建时间
}

//删除SKU
type SkuDelRequest struct {
	APIRequest
	ProductId    uint64 `json:"product_id"`     //小商店内部商品ID，与out_product_id二选一
	OutProductId string `json:"out_product_id"` //商家自定义商品ID，与product_id二选一
	SkuId        uint64 `json:"sku_id"`         //out_sku_id、sku_id 二选一
	OutSkuId     string `json:"out_sku_id"`     //out_sku_id、sku_id 二选一
}
type SkuDelResponse struct {
	APIResponse
}

//获取SKU信息
type SkuGetRequest struct {
	APIRequest
	SkuId         uint64        `json:"sku_id"`          //out_sku_id、sku_id 二选一
	OutSkuId      string        `json:"out_sku_id"`      //out_sku_id、sku_id 二选一
	NeedEditSku   NeedEditSku   `json:"need_edit_sku"`   // 默认0:获取线上数据, 1:获取草稿数据
	NeedRealStock NeedRealStock `json:"need_real_stock"` // 默认0:获取草稿库存, 1:获取线上真实库存
}
type SkuGetResponse struct {
	APIResponse
	Data *SkuGetResponseData `json:"data"`
}
type SkuGetResponseData struct {
	SkuData
	Status SkuStatus `json:"status"` //sku状态
}

//批量获取SKU信息
type SkuGetListRequest struct {
	APIRequest
	ProductId     uint64        `json:"product_id"`      //小商店内部商品ID
	NeedEditSku   NeedEditSku   `json:"need_edit_sku"`   // 默认0:获取线上数据, 1:获取草稿数据
	NeedRealStock NeedRealStock `json:"need_real_stock"` // 默认0:获取草稿库存, 1:获取线上真实库存
}
type SkuGetListResponse struct {
	APIResponse
	Data []*SkuData `json:"skus"`
}

//更新SKU信息
type SkuUpdateRequest struct {
	APIRequest
	SkuData
}
type SkuUpdateResponse struct {
	APIResponse
	Data *SkuUpdateData `json:"data"`
}
type SkuUpdateData struct {
	SkuId      uint64 `json:"sku_id"`      //skuID
	UpdateTime string `json:"update_time"` //更新时间
}

//更新SKU价格
type SkuUpdatePriceRequest struct {
	APIRequest
	SkuData
}
type SkuUpdatePriceResponse struct {
	APIResponse
	Data *SkuUpdatePriceData `json:"data"`
}
type SkuUpdatePriceData struct {
	SkuId      uint64 `json:"sku_id"`      //skuID
	UpdateTime string `json:"update_time"` //更新时间
}

//更新SKU库存
type SkuUpdateStockRequest struct {
	APIRequest
	SkuData
}
type SkuUpdateStockResponse struct {
	APIResponse
	Data *SkuUpdatePriceData `json:"data"`
}
type SkuUpdateStockData struct {
	UpdateTime string `json:"update_time"` //更新时间
}

//获取SKU库存
type SkuGetStockRequest struct {
	APIRequest
	SkuData
}
type SkuGetStockResponse struct {
	APIResponse
	Data *SkuUpdatePriceData `json:"data"`
}
type SkuGetStockData struct {
	StockNum uint64 `json:"stock_num"` //库存量
}
