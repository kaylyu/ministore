package models

//获取订单列表
type OrderGetListRequest struct {
	APIRequest
	StartCreateTime string      `json:"start_create_time"` //开始时间 2020-03-25 12:05:25
	EndCreateTime   string      `json:"end_create_time"`   //结束时间 2020-04-25 12:05:25
	StartUpdateTime string      `json:"start_update_time"` //开始时间 2020-03-25 12:05:25
	EndUpdateTime   string      `json:"end_update_time"`   //结束时间 2020-04-25 12:05:25
	Page            uint64      `json:"page"`              //页码
	PageSize        uint64      `json:"page_size"`         //每页个数
	Status          OrderStatus `json:"status"`            //订单状态
}
type OrderGetListResponse struct {
	APIResponse
	Data     []*OrderData `json:"orders"`
	TotalNum uint64       `json:"total_num"` //订单总数
}

type OrderData struct {
	OrderId         uint64           `json:"order_id"`    //订单ID
	Openid          string           `json:"openid"`      //用户的openid，用于物流助手接口
	TotalNum        string           `json:"total_num"`   //订单总数
	CreateTime      string           `json:"create_time"` //创建时间
	UpdateTime      string           `json:"update_time"` //更新时间
	OrderDetail     *OrderDetail     `json:"order_detail"`
	AftersaleDetail *AftersaleDetail `json:"aftersale_detail"`
	ExtInfo         *ExtInfo         `json:"ext_info"`
}

type OrderDetail struct {
	ProductInfos []*ProductInfo `json:"product_infos"`
	PayInfo      *PayInfo       `json:"pay_info"`
	PriceInfo    *PriceInfo     `json:"price_info"`
	DeliveryInfo *DeliveryInfo  `json:"delivery_info"`
}

type ProductInfo struct {
	ProductId             uint64     `json:"product_id"`               //小商店内部商品ID
	SkuId                 uint64     `json:"sku_id"`                   //小商店内部skuID
	Title                 string     `json:"title"`                    //标题
	ThumbImg              string     `json:"thumb_img"`                //sku小图
	SkuCnt                uint64     `json:"sku_cnt"`                  //sku数量
	OnAftersaleSkuCnt     uint64     `json:"on_aftersale_sku_cnt"`     //正在售后/退款流程中的sku数量
	FinishAftersaleSkuCnt uint64     `json:"finish_aftersale_sku_cnt"` //完成售后/退款的sku数量
	SalePrice             uint64     `json:"sale_price"`               //售卖价格,以分为单位
	SkuAttrs              []*SkuAttr `json:"sku_attrs"`
}
type PayInfo struct {
	PayMethod     string `json:"pay_method"`     //支付方式（目前只有"微信支付"）
	PrepayId      string `json:"prepay_id"`      //预支付ID
	TransactionId string `json:"transaction_id"` //支付订单号
	PrepayTime    string `json:"prepay_time"`    //预付款时间
	PayTime       string `json:"pay_time"`       //付款时间
}
type PriceInfo struct {
	ProductPrice    uint64 `json:"product_price"`    //	商品金额（单位：分）
	OrderPrice      uint64 `json:"order_price"`      //订单金额（单位：分）
	Freight         uint64 `json:"freight"`          //运费（单位：分）
	DiscountedPrice uint64 `json:"discounted_price"` //优惠金额（单位：分）
	IsDiscounted    bool   `json:"is_discounted"`    //是否有优惠（false：无优惠/true：有优惠）
}
type DeliveryInfo struct {
	DeliveryMethod      string                 `json:"delivery_method"` //快递方式（目前只有"快递"）
	DeliveryTime        string                 `json:"delivery_time"`   //发货时间
	DeliveryProductInfo []*DeliveryProductInfo `json:"delivery_product_info"`
	DeliveryAddressInfo *DeliveryAddressInfo   `json:"address_info"`
	ExpressFee          []*ExpressFee          `json:"express_fee"`
	InsuranceInfo       *InsuranceInfo         `json:"insurance_info"`
}
type DeliveryProductInfo struct {
	WaybillId  string `json:"waybill_id"`  //快递单号
	DeliveryId string `json:"delivery_id"` //快递公司编号
}
type DeliveryAddressInfo struct {
	UserName     string `json:"user_name"`     //收货人姓名
	PostalCode   string `json:"postal_code"`   //邮编
	ProvinceName string `json:"province_name"` //国标收货地址第一级地址
	CityName     string `json:"city_name"`     //国标收货地址第二级地址
	CountryName  string `json:"county_name"`   //国标收货地址第三级地址
	DetailInfo   string `json:"detail_info"`   //详细收货地址信息
	NationalCode string `json:"national_code"` //收货地址国家码
	TelNumber    string `json:"tel_number"`    //收货人手机号码
}
type ExpressFee struct {
	Result         uint64 `json:"result"`
	ShippingMethod string `json:"shipping_method"`
}
type InsuranceInfo struct {
	Type           string `json:"type"`
	InsurancePrice uint64 `json:"insurance_price"`
}

type AftersaleDetail struct {
	AftersaleOrderList  []*AftersaleOrderList `json:"aftersale_order_list"`   //售后列表
	OnAftersaleOrderCnt uint64                `json:"on_aftersale_order_cnt"` //正在售后流程的售后单数
}
type AftersaleOrderList struct {
	AftersaleOrderId uint64 `json:"aftersale_order_id"` //售后单ID（售后接口开放后可用于查询）
}

type ExtInfo struct {
	CustomerNotes string `json:"customer_notes"` //用户备注
	MerchantNotes string `json:"merchant_notes"` //商家备注
}

//获取订单详情
type OrderDetailRequest struct {
	APIRequest
	OrderId uint64 `json:"order_id"` //订单ID
}
type OrderDetailResponse struct {
	APIResponse
	Data *OrderData `json:"order"`
}

//搜索订单
type OrderSearchRequest struct {
	APIRequest
	StartPayTimeTime      string      `json:"start_pay_time,omitempty"` //开始时间 2020-03-25 12:05:25
	EndPayTimeTime        string      `json:"end_pay_time,omitempty"`   //结束时间 2020-04-25 12:05:25
	Title                 string      `json:"title,omitempty"`          //商品标题关键词
	SkuCode               string      `json:"sku_code,omitempty"`       //	商品编码
	UserName              string      `json:"user_name,omitempty"`      //收件人
	TelNumber             string      `json:"tel_number,omitempty"`     //收件人电话
	OnAftersaleOrderExist uint        `json:"on_aftersale_order_exist"` //不填该参数:全部订单 0:没有正在售后的订单, 1:正在售后单数量大于等于1的订单
	Page                  uint64      `json:"page"`                     //页码
	PageSize              uint64      `json:"page_size"`                //每页个数
	Status                OrderStatus `json:"status,omitempty"`         //订单状态
}
type OrderSearchResponse struct {
	APIResponse
	Data     []*OrderData `json:"orders"`
	TotalNum uint64       `json:"total_num"` //订单总数
}
