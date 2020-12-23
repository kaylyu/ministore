package models

//获取快递公司列表
type DeliveryGetCompanyRequest struct {
	APIRequest
}
type DeliveryGetCompanyResponse struct {
	APIResponse
	Data []*CompanyList `json:"company_list"`
}

type CompanyList struct {
	DeliveryId   string `json:"delivery_id"`   //快递公司id
	DeliveryName string `json:"delivery_name"` //快递公司名称
}

//订单发货
type DeliverySendRequest struct {
	APIRequest
	OrderId      uint64                 `json:"order_id"`      //订单id
	DeliveryList []*DeliveryProductInfo `json:"delivery_list"` //快递列表
}
type DeliverySendResponse struct {
	APIResponse
}
