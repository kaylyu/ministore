package coupon

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiCreateCoupon = "/product/coupon/create"
	apiGetCoupon    = "/product/coupon/get_list"
	apiPushCoupon   = "/product/coupon/push"
)

//商品
type Coupon struct {
	core *component.Core
}

/*
创建优惠券实例
*/
func NewCoupon(core *component.Core) *Coupon {
	return &Coupon{core}
}

/*
创建优惠券
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/create_coupon.html
*/
func (c *Coupon) CreateCoupon(request *models.CouponCreateRequest) (response models.CouponCreateResponse, err error) {
	err = c.core.HttpPostJson(apiCreateCoupon, request, &response)
	return
}

/*
获取优惠券列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/get_coupon.html
*/
func (c *Coupon) GetCouponList(request *models.CouponGetListRequest) (response models.CouponGetListResponse, err error) {
	err = c.core.HttpPostJson(apiGetCoupon, request, &response)
	return
}

/*
发放优惠券
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/push_coupon.html
*/
func (c *Coupon) PushCoupon(request *models.CouponPusRequest) (response models.CouponPusResponse, err error) {
	err = c.core.HttpPostJson(apiGetCoupon, request, &response)
	return
}
