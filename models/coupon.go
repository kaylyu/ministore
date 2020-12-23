package models

//优惠券信息
type CouponInfo struct {
	Name          string         `json:"name"`          //优惠券名称
	DiscountInfo  *DiscountInfo  `json:"discount_info"` //商品
	CouponExtInfo *CouponExtInfo `json:"ext_info"`      //优惠券
	PromoteInfo   *PromoteInfo   `json:"promote_info"`  //推广
	ReceiveInfo   *ReceiveInfo   `json:"receive_info"`  //领用
	ValidInfo     *ValidInfo     `json:"valid_info"`    //有效
}
type DiscountInfo struct {
	DiscountCondition *DiscountCondition `json:"discount_condition"`     //商品折扣条件
	DiscountFee       uint64             `json:"discount_fee"`           //满减金额
	DiscountNum       uint64             `json:"discount_num,omitempty"` //打折商品数量，满减券需填写
}
type DiscountCondition struct {
	ProductCnt   uint     `json:"product_cnt,omitempty"`   //商品折扣券打折金额
	ProductIds   []uint64 `json:"product_ids,omitempty"`   //商品id，商品折扣券需填写
	ProductPrice uint64   `json:"product_price,omitempty"` //商品价格，满减券需填写
}
type CouponExtInfo struct {
	InvalidTime   string `json:"invalid_time"`              //优惠券失效时间
	JumpProductId uint64 `json:"jump_product_id,omitempty"` //商品折扣券领取后跳转的商品id
	Notes         string `json:"notes,omitempty"`           //备注信息
	ValidTime     string `json:"valid_time"`                //优惠券有效时间
}
type PromoteInfo struct {
	CustomizeChannel string `json:"customize_channel"` //用户自定义推广渠道
	PromoteType      string `json:"promote_type"`      //	推广类型
}
type ReceiveInfo struct {
	EndTime           string `json:"end_time"`             //优惠券领用结束时间
	LimitNumOnePerson uint64 `json:"limit_num_one_person"` //是否限制一人使用
	StartTime         string `json:"start_time"`           //优惠券领用开始时间
	TotalNum          uint64 `json:"total_num"`            //优惠券领用总数
}
type ValidInfo struct {
	EndTime     uint64 `json:"end_time,omitempty"`   //优惠券有效期结束时间，若填了start必填
	StartTime   uint64 `json:"start_time,omitempty"` //优惠券有效期开始时间，和valid_day_num二选一
	ValidDayNum uint64 `json:"valid_day_num"`        //优惠券有效期天数，和start_time二选一
	ValidType   string `json:"valid_type"`           //优惠券有效期类型
}
type StockInfo struct {
	IssuedNum  uint64 `json:"issued_num"`
	ReceiveNum uint64 `json:"receive_num"`
	UsedNum    uint64 `json:"used_num"`
}

//创建优惠券
type CouponCreateRequest struct {
	APIRequest
	Type       CouponType  `json:"type"`
	CouponInfo *CouponInfo `json:"coupon_info"`
}
type CouponCreateResponse struct {
	APIResponse
	Data []*CouponCreateData `json:"data"`
}
type CouponCreateData struct {
	CouponId uint64 `json:"coupon_id"` //优惠券ID
}

//获取优惠券列表
type CouponGetListRequest struct {
	APIRequest
	StartCrateTime string       `json:"start_crate_time"` //开始时间 2020-03-25 12:05:25
	EndCreateTime  string       `json:"end_create_time"`  ////结束时间 2020-04-25 12:05:25
	Page           uint64       `json:"page"`             //页码
	PageSize       uint64       `json:"page_size"`        //每页个数
	Status         CouponStatus `json:"status,omitempty"` //优惠券状态
}
type CouponGetListResponse struct {
	APIResponse
	Data []*CouponGetListData `json:"coupons"`
}
type CouponGetListData struct {
	CouponId   uint64       `json:"coupon_id"`   //优惠券ID
	Type       CouponType   `json:"type"`        //优惠券类型
	Status     CouponStatus `json:"status"`      //优惠券状态
	CreateTime string       `json:"create_time"` //优惠券创建时间
	UpdateTime string       `json:"update_time"` //优惠券更新时间
	CouponInfo *CouponInfo  `json:"coupon_info"` //优惠券
	StockInfo  *StockInfo   `json:"stock_info"`  //库存
}

//发放优惠券
type CouponPusRequest struct {
	APIRequest
	Openid   string `json:"openid"`    //用户openid
	CouponId string `json:"coupon_id"` //优惠券ID，可从获取优惠券中获得
}
type CouponPusResponse struct {
	APIResponse
}
