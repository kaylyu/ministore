package coupon

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiCreateCoupon       = "/product/coupon/create"
	apiUpdateCoupon       = "/product/coupon/update"
	apiUpdateStatusCoupon = "/product/coupon/update_status"
	apiGetCoupon          = "/product/coupon/get"
	apiGetListCoupon      = "/product/coupon/get_list"
	apiPushCoupon         = "/product/coupon/push"
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
func (c *Coupon) CreateCoupon(request *models.CouponCreateRequest) (response models.CouponCreateResponse, body string, err error) {
	body, err = c.core.HttpPostJson(apiCreateCoupon, request, &response)
	return
}

/*
更新优惠券
https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent/API/coupon/update_coupon.html
*/
func (c *Coupon) UpdateCoupon(request *models.CouponUpdateRequest) (response models.CouponUpdateResponse, body string, err error) {
	body, err = c.core.HttpPostJson(apiUpdateCoupon, request, &response)
	return
}

/*
更新优惠券状态
https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent/API/coupon/update_coupon_status.html
*/
func (c *Coupon) UpdateStatusCoupon(request *models.CouponUpdateStatusRequest) (response models.CouponUpdateStatusResponse, body string, err error) {
	body, err = c.core.HttpPostJson(apiUpdateStatusCoupon, request, &response)
	return
}

/*
获取优惠券
https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent/API/coupon/update_coupon_status.html
*/
func (c *Coupon) GetCoupon(request *models.CouponGetRequest) (response models.CouponGetResponse, body string, err error) {
	body, err = c.core.HttpPostJson(apiGetCoupon, request, &response)
	return
}

/*
获取优惠券列表
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/get_coupon.html
*/
func (c *Coupon) GetCouponList(request *models.CouponGetListRequest) (response models.CouponGetListResponse, body string, err error) {
	body, err = c.core.HttpPostJson(apiGetListCoupon, request, &response)
	return
}

/*
发放优惠券
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/coupon/push_coupon.html
*/
func (c *Coupon) PushCoupon(request *models.CouponPusRequest) (response models.CouponPusResponse, body string, err error) {
	body, err = c.core.HttpPostJson(apiPushCoupon, request, &response)
	return
}
