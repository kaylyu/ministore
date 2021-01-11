package models

//计费类型
type ValuationType uint

const (
	ValuationTypeCase   ValuationType = iota + 1 //按件
	ValuationTypeWeight                          //按重量
)

//服务类型
type ServiceType uint

const (
	ServiceTypeGoods     ServiceType = 2 //商品
	ServiceTypeOrder     ServiceType = 3 //订单
	ServiceTypeLogistics ServiceType = 4 //物流
)

//订单状态
type ServiceOrderStatus uint

const (
	ServiceOrderStatusTakeEffect ServiceOrderStatus = iota + 1 //生效中
	ServiceOrderStatusExpire                                   //已过期
	ServiceOrderStatusBlock                                    //已停用
	ServiceOrderStatusUpaid                                    //未付款
)

//商品线上状态
type SpuStatus uint

const (
	SpuStatusInitial         SpuStatus = 0  //初始值
	SpuStatusPutaway         SpuStatus = 5  //上架
	SpuStatusTrash           SpuStatus = 6  //回收站
	SpuStatusLogicallyDelete SpuStatus = 9  //逻辑删除
	SpuStatusSelfDown        SpuStatus = 11 //自主下架
	SpuStatusSellOutDown     SpuStatus = 12 //售磬下架
	SpuStatusIllegalDown     SpuStatus = 13 //违规下架/风控系统下架
)

//商品草稿状态
type SpuEditStatus uint

const (
	SpuEditStatusInitial      SpuEditStatus = iota //初始值
	SpuEditStatusEdit                              //编辑中
	SpuEditStatusVerifying                         //审核中
	SpuEditStatusVerifyFailed                      //审核失败
	SpuEditStatusVerifyPass                        //审核成功
)

//默认0:获取线上数据, 1:获取草稿数据
type NeedEditSpu uint

const (
	NeedEditSpuOnline NeedEditSpu = iota //0:获取线上数据
	NeedEditSpuDraft                     //1:获取草稿数据
)

// 默认0:获取线上数据, 1:获取草稿数据
type NeedEditSku uint

const (
	NeedEditSkuOnline NeedEditSku = iota //0:获取线上数据
	NeedEditSkuDraft                     //1:获取草稿数据
)

// 默认0:获取草稿库存, 1:获取线上真实库存
type NeedRealStock uint

const (
	NeedRealStockDraft  NeedRealStock = iota //0:获取草稿库存
	NeedRealStockOnline                      //1:获取线上真实库存
)

// SKU状态
type SkuStatus uint

const (
	SkuStatusPutaway SkuStatus = 5  //上架中
	SkuStatusSoftDel SkuStatus = 21 //假删除
)

//订单类型
type OrderStatus uint

const (
	OrderStatusUnpaid         OrderStatus = 10  //待付款
	OrderStatusSendPending    OrderStatus = 20  //待发货
	OrderStatusReceivePending OrderStatus = 30  //待收货
	OrderStatusComplete       OrderStatus = 100 //完成
	OrderStatusCompleteCancel OrderStatus = 200 //全部商品售后之后，订单取消
	OrderStatusTimeoutCancel  OrderStatus = 250 //用户主动取消或待付款超时取消
)

//优惠券类型
type CouponType uint

const (
	CouponTypeCondition     CouponType = 1   //1	商品条件折券, discount_condition.product_ids, discount_condition.product_cnt, discount_info.discount_num 必填
	CouponTypeFull          CouponType = 2   //2	商品满减券, discount_condition.product_ids, discount_condition.product_price, discount_info.discount_fee 必填
	CouponTypeComm          CouponType = 3   //3	商品统一折扣券, discount_condition.product_ids, discount_info.discount_num必填
	CouponTypeDirect        CouponType = 4   //4	商品直减券, 如果小于可用的商品中的最小价格会提醒(没有商品时超过50w提醒）, discount_condition.product_ids, discount_fee 必填
	CouponTypeShopCondition CouponType = 101 //101	店铺条件折扣券, discount_condition.product_cnt, discount_info.discount_num必填
	CouponTypeShopFull      CouponType = 102 //102	店铺满减券, discount_condition.product_price, discount_info.discount_fee 必填
	CouponTypeShopComm      CouponType = 103 //103	店铺统一折扣券, discount_info.discount_num 必填
	CouponTypeShopDirect    CouponType = 104 //104	店铺直减券, 如果小于可用的商品中的最小价格会提醒(没有商品时超过50w提醒）, discount_fee 必填
)

//优惠券状态
type CouponStatus uint

const (
	CouponStatusEditing        CouponStatus = 1   //未生效，编辑中
	CouponStatusValid          CouponStatus = 2   //生效
	CouponStatusExpired        CouponStatus = 3   //已过期
	CouponStatusCancellation   CouponStatus = 4   //已作废
	CouponStatusDel            CouponStatus = 5   //删除
	CouponStatusEffectOf       CouponStatus = 100 //生效中
	CouponStatusAlreadyExpired CouponStatus = 101 //已过期
	CouponStatusUsed           CouponStatus = 102 //已经使用
)

//1: 企业店, 2: 个人店
type StoreType uint

const (
	StoreTypeEnterprise StoreType = iota //1: 企业店
	StoreTypePerson                      // 2: 个人店
)

//状态枚举值
type StoreStatus uint

const (
	StoreStatusNonOpr       StoreStatus = iota + 1 //1: 当前不可操作
	StoreStatusNormal                              // 2: 正常，当前可操作
	StoreStatusVerifying                           // 3: 已提交，审核中
	StoreStatusVerifyFailed                        // 4: 审核失败，被驳回
	StoreStatusVerifySuc                           // 5: 成功
	StoreStatusFreeze                              // 6: 冻结
)

//注册状态
type StoreRegisterStatus uint

const (
	StoreRegisterStatusSuc            StoreRegisterStatus = iota //0:成功
	StoreRegisterStatusNonSign                                   // 1:已发送协议还未签约
	StoreRegisterStatusUnSendProtocol                            // 2: 未发送协议或协议已过期，需发送协议
)

//注册小商店类型 1-整店打包，2-组件开放
type StoreRegisterApiOpenStoreType uint

const (
	StoreRegisterApiOpenStoreTypeShopPackage StoreRegisterApiOpenStoreType = iota + 1 // 1-整店打包
	StoreRegisterApiOpenStoreTypeComponent                                            //2-组件开放
)

//主体类型
type StoreRegisterSubjectType string

const (
	StoreRegisterSubjectTypeEnterprise StoreRegisterSubjectType = "2" //"2"：企业，营业执照上的主体类型一般为有限公司、有限责任公司。
	StoreRegisterSubjectTypePerson     StoreRegisterSubjectType = "4" //"4"：个体工商户，营业执照上的主体类型一般为个体户、个体工商户、个体经营。
)

//装修状态1代表启用，2代表停用
type SwitchStatus uint

const (
	SwitchStatusEnable  SwitchStatus = 1 //1代表启用
	SwitchStatusDisable SwitchStatus = 2 //2代表停用
)
