package order

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiGetOrderList   = "/product/order/get_list"
	apiGetOrderDetail = "/product/order/get"
	apiGetOrderSearch = "/product/order/search"
)

//商品
type Order struct {
	core *component.Core
}

/*
创建订单实例
*/
func NewOrder(core *component.Core) *Order {
	return &Order{core}
}

/*
获取订单列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/get_order_list.html
*/
func (o *Order) GetOrderList(request *models.OrderGetListRequest) (response models.OrderGetListResponse, body string, err error) {
	err = o.core.HttpPostJson(apiGetOrderList, request, &response)
	return
}

/*
获取订单详情
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/get_order_detail.html
*/
func (o *Order) GetOrderDetail(request *models.OrderDetailRequest) (response models.OrderDetailResponse, body string, err error) {
	err = o.core.HttpPostJson(apiGetOrderDetail, request, &response)
	return
}

/*
搜索订单
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/order/search_order.html
*/
func (o *Order) GetOrderSearch(request *models.OrderSearchRequest) (response models.OrderSearchResponse, body string, err error) {
	err = o.core.HttpPostJson(apiGetOrderSearch, request, &response)
	return
}
