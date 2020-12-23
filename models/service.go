package models

//登录验证
type CheckAuthRequest struct {
	APIRequest
	Code string `json:"code"` //跳转码(小商店服务市场跳转到第三方url里面会带上code)
}

//登录验证
type CheckAuthResponse struct {
	APIResponse
	Data *CheckAuth `json:"data"`
}

type CheckAuth struct {
	Appid     string `json:"appid"`      //小程序ID
	ServiceId uint64 `json:"service_id"` //服务ID
}

//获取用户购买的在有效期内的服务列表
type GetListRequest struct {
	APIRequest
}

//获取用户购买的在有效期内的服务列表
type GetListResponse struct {
	APIResponse
	Data []*GetServiceList `json:"service_list"`
}

type GetServiceList struct {
	ServiceId   uint64      `json:"service_id"`   //服务ID
	ServiceName string      `json:"service_name"` //服务名称
	ExpireTime  string      `json:"expire_time"`  //过期时间
	ServiceType ServiceType `json:"service_type"` //服务类型
}

//获取用户购买的在有效期内的服务列表
type GetServiceOrderListRequest struct {
	APIRequest
	StartCreateTime string `json:"start_create_time"` //开始时间 2020-03-25 12:05:25
	EndCreateTime   string `json:"end_create_time"`   //结束时间 2020-04-25 12:05:25
	Page            uint64 `json:"page"`              //页码
	PageSize        uint64 `json:"page_size"`         //每页个数
}

//获取用户购买的在有效期内的服务列表
type GetServiceOrderListListResponse struct {
	APIResponse
	Data []*ServiceOrderList `json:"service_order_list"`
}

type ServiceOrderList struct {
	ServiceOrderId  uint64             `json:"service_order_id"` //服务订单ID
	ServiceId       uint64             `json:"service_id"`       //服务ID
	ServiceName     string             `json:"service_name"`     //服务名称
	CreateTime      string             `json:"create_time"`      //创建时间
	ExpireTime      string             `json:"expire_time"`      //过期时间
	ServiceType     ServiceType        `json:"service_type"`     //服务类型
	SpecificationId string             `json:"specification_id"` //规格英文名称
	TotalPrice      uint64             `json:"total_price"`      //订单总价格
	Status          ServiceOrderStatus `json:"status"`           //订单状态
}
