package delivery

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiGetDeliveryCompanyList = "/product/delivery/get_company_list"
	apiSendDelivery           = "/product/delivery/send"
)

//商品
type Delivery struct {
	core *component.Core
}

/*
创建物流实例
*/
func NewDelivery(core *component.Core) *Delivery {
	return &Delivery{core}
}

/*
获取快递公司列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/delivery/get_delivery_company_list.html
*/
func (d *Delivery) GetDeliveryCompanyList(request *models.DeliveryGetCompanyRequest) (response models.DeliveryGetCompanyResponse, err error) {
	err = d.core.HttpPostJson(apiGetDeliveryCompanyList, request, &response)
	return
}

/*
订单发货
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/delivery/send_delivery.html
*/
func (d *Delivery) SendDelivery(request *models.DeliverySendRequest) (response models.DeliverySendResponse, err error) {
	err = d.core.HttpPostJson(apiSendDelivery, request, &response)
	return
}
