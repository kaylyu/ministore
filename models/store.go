package models

//获取基本信息
type StoreGetInfoRequest struct {
	APIRequest
}
type StoreGetInfoResponse struct {
	APIResponse
	Data *StoreGetInfoData `json:"data"`
}
type StoreGetInfoData struct {
	StoreName string    `json:"store_name"` //商店名称
	Logo      string    `json:"logo"`       //商店头像
	Type      StoreType `json:"type"`       //1: 企业店, 2: 个人店
}

//上传图片
type StoreImgUploadRequest struct {
	APIRequest
	Height uint64 `json:"height"`
	Width  uint64 `json:"width"`
	Media  string `json:"media"`
}
type StoreImgUploadResponse struct {
	APIResponse
	Data *ImgUploadData `json:"pic_file"`
}
type ImgUploadData struct {
	MediaId    string `json:"media_id"`     //media_id
	PayMediaId string `json:"pay_media_id"` //支付media_id
}

//异步状态查询
type StoreRegisterCheckAuditStatusRequest struct {
	APIRequest
	WxName string `json:"wx_name"` //微信号
}
type StoreRegisterCheckAuditStatusResponse struct {
	APIResponse
	Data *StoreRegisterCheckAuditStatusData `json:"data"`
}
type StoreRegisterCheckAuditStatusData struct {
	RegisterStatus     StoreRegisterStatus `json:"register_status"`      //注册状态 0:成功 1:已发送协议还未签约 2: 未发送协议或协议已过期，需发送协议，当register_status为0时以下字段有意义
	MerchantInfoStatus StoreStatus         `json:"merchant_info_status"` //商家信息状态, 具体含义查看状态枚举值
	AcctVerifyStatus   StoreStatus         `json:"acct_verify_status"`   //账户验证状态, 具体含义查看状态枚举值
	BasicInfoStatus    StoreStatus         `json:"basic_info_status"`    //基础信息状态, 具体含义查看状态枚举值
	PaySignStatus      StoreStatus         `json:"pay_sign_status"`      //支付签约状态, 具体含义查看状态枚举值
	AuditRejectReasons string              `json:"audit_reject_reasons"` //基础信息驳回原因
	LegalValidationUrl string              `json:"legal_validation_url"` //法人验证链接
	PayAuditDetail     []*PayAuditDetail   `json:"pay_audit_detail"`
	RegisteredAppid    string              `json:"registered_appid"` //注册的appid
}
type PayAuditDetail struct {
	ParamName    string `json:"param_name"`
	RejectReason string `json:"reject_reason"`
}

//注册小商店账号
type StoreRegisterRequest struct {
	APIRequest
	WxName           string                        `json:"wx_name"`                 //微信号
	IdCardName       string                        `json:"id_card_name"`            //身份证姓名
	IdCardNumber     string                        `json:"id_card_number"`          //身份证号
	ChannelId        string                        `json:"channel_id,omitempty"`    //	渠道号
	ApiOpenstoreType StoreRegisterApiOpenStoreType `json:"api_openstore_type"`      //1-整店打包，2-组件开放
	AuthPageUrl      string                        `json:"auth_page_url,omitempty"` //授权页url
}
type StoreRegisterResponse struct {
	APIResponse
}

//提交支付资质
type StoreRegisterSubmitMerchantInfoRequest struct {
	APIRequest
	Appid                  string                   `json:"appid"`                            //openid
	SubjectType            StoreRegisterSubjectType `json:"subject_type"`                     //主体类型
	BusiLicense            *BusiLicense             `json:"busi_license"`                     //营业执照/登记证书信息
	OrganizationCodeInfo   *OrganizationCodeInfo    `json:"organization_code_info,omitempty"` //组织机构代码证信息
	IdCardInfo             *IdCardInfo              `json:"id_card_info"`                     //身份证
	SuperAdministratorInfo *SuperAdministratorInfo  `json:"super_administrator_info"`         //超级管理员信息 请填写店铺的超级管理员信息。超级管理员需在开户后进行签约， 并可接收日常重要管理信息和进行资金操作，请确定其为商户法定代表人或负责人。
	MerchantShortname      string                   `json:"merchant_shortname"`               //户简称 UTF-8格式，中文占3个字节，即最多16个汉字长度。 将在支付完成页向买家展示，需与商家的实际售卖商品相符 。示例值：腾讯
	SpecialQualification   *SpecialQualification    `json:"special_qualification,omitempty"`  //特殊资质
	SupplementaryMaterial  *SupplementaryMaterial   `json:"supplementary_material,omitempty"` //补充材料
	SupplementaryDesc      string                   `json:"supplementary_desc,omitempty"`     //补充描述
}
type StoreRegisterSubmitMerchantInfoResponse struct {
	APIResponse
}
type BusiLicense struct {
	LicenseType         uint64   `json:"license_type"`         //营业执照类型 1:三证合一 2: 普通营业执照
	PicFile             *PicFile `json:"pic_file"`             //营业执照图片
	RegistrationNum     string   `json:"registration_num"`     //请填写营业执照上的注册号/统一社会信用代码， 须为15位数字或18位数字大写字母。示例值：123456789012345678 特殊规则：长度最小15个字节
	MerchantName        string   `json:"merchant_name"`        //1、请填写营业执照/登记证书的商家名称，2~110个字符，支持括号 2、个体工商户/党政、机关及事业单位，不能以“公司”结尾。 3、个体工商户，若营业执照上商户名称为空或为“无”，请填写"个体户+经营者姓名"， 如“个体户张三” 。示例值：腾讯科技有限公司
	LegalRepresentative string   `json:"legal_representative"` //	请填写证件的经营者/法定代表人姓名。示例值：张三
	RegisteredAddrs     string   `json:"registered_addrs"`     //注册地址
	StartDate           string   `json:"start_date"`           //注册日期，格式：2014-01-01
	EndDate             string   `json:"end_date"`             //结束有效期，格式：2014-01-01 1、若证件有效期为长期，请填写：长期。 2、结束时间需大于开始时间。3、有效期必须大于60天，即结束时间距当前时间需超过60天
}
type PicFile struct {
	MediaId    string `json:"media_id"`     //media_id
	PayMediaId string `json:"pay_media_id"` //支付media_id
}
type OrganizationCodeInfo struct {
	PicFile          *PicFile `json:"pic_file"`          //组织结构代码证图片
	OrganizationCode string   `json:"organization_code"` //1、请填写组织机构代码证上的组织机构代码。 2、可填写9或10位 数字\字母\连字符。示例值：12345679-A
	StartDate        string   `json:"start_date"`        //注册日期，格式：2014-01-01
	EndDate          string   `json:"end_date"`          //结束有效期，格式：2014-01-01 1、若证件有效期为长期，请填写：长期。2、结束时间需大于开始时间。3、有效期必须大于60天，即结束时间距当前时间需超过60天。
}
type IdCardInfo struct {
	PortraitPicFile *PicFile `json:"portrait_pic_file"` //人像面照片，格式同pic_file
	NationPicFile   *PicFile `json:"nation_pic_file"`   //国徽面照片，格式同pic_file
	IdCardName      string   `json:"id_card_name"`      //请填写经营者/法定代表人对应身份证的姓名，2~30个中文字符、英文字符、符号。
	IdCardNumber    string   `json:"id_card_number"`    //请填写经营者/法定代表人对应身份证的号码
	StartDate       string   `json:"start_date"`        //注册日期，格式：2014-01-01
	EndDate         string   `json:"end_date"`          //结束有效期，格式：2014-01-01 1、若证件有效期为长期，请填写：长期。 2、结束时间需大于开始时间。3、有效期必须大于60天，即结束时间距当前时间需超过60天
}
type SuperAdministratorInfo struct {
	Type         string `json:"type"`           //个体工商户/企业/党政、机关及事业单位/其他组织，可选择：65-法人/经营者、66- 负责人。 （负责人：经商户授权办理微信支付业务的人员，授权范围包括但不限于签约，入驻过程需完成账户验证）。示例值：65
	Name         string `json:"name"`           //1、若管理员类型为“法人”，则该姓名需与法人身份证姓名一致。 2、若管理员类型为“经办人”，则可填写实际经办人的姓名。
	IdCardNumber string `json:"id_card_number"` //	1、若管理员类型为法人，则该身份证号码需与法人身份证号码一致。若管理员类型为经办人， 则可填写实际经办人的身份证号码。 2、可传身份证、来往内地通行证、来往大陆通行证、护照等证件号码。3、超级管理员签约时，校验微信号绑定的银行卡实名信息，是否与该证件号码一致。
	Phone        string `json:"phone"`          //请填写管理员的手机号，11位数字， 用于接收微信支付的重要管理信息及日常操作验证码
	Mail         string `json:"mail"`           //1、用于接收微信支付的开户邮件及日常业务通知。 2、需要带@，遵循邮箱格式校验 。
}
type SpecialQualification struct {
	PicFileList []*ImgUploadData `json:"pic_file_list"` //图片media id数组，每项的格式同pic_file
}
type SupplementaryMaterial struct {
	PicFileList []*ImgUploadData `json:"pic_file_list"` //图片media id数组，每项的格式同pic_file
}

//提交小商店基础信息
type StoreRegisterSubmitBasicInfoRequest struct {
	APIRequest
	Appid      string      `json:"appid"`       //openid
	NameInfo   *NameInfo   `json:"name_info"`   //名称信息
	ReturnInfo *ReturnInfo `json:"return_info"` //退货地址
}
type StoreRegisterSubmitBasicInfoResponse struct {
	APIResponse
}
type NameInfo struct {
	NickName         string   `json:"nickname"`                     //1）小程序名称可以由中文、数字、英文、空格及部分特殊符号(“空格”、“-”、“+”、“&”、“.”)组成。长度在4-30个字符之间，一个中文字等于2个字符。 2）公众号、小程序在微信公众平台上的名称是唯一的，且属于同一主体下，可以重名。3）不得与不同主体的公众号名称重名。
	Abbr             string   `json:"abbr"`                         //1）小程序简称可以从小程序名称中按顺序截取字符创建。长度在4-10个字符之间，一个中文字等于2个字符。 2）小程序简称在微信公众平台是不唯一的，可以重名。但对于仿冒、侵权等恶意情况，平台仍会做出相关处罚。开发者也可通过侵权投诉维护自己的正当权益。3）小程序简称设置后，将在客户端任务栏向用户展示。开发者可以凭借此功能，更好地实现产品品牌价值和展示。目前暂不支持名称的其他功能。
	Introduction     string   `json:"introduction"`                 //请确认介绍内容不含国家相关法律法规禁止内容,介绍字数为4-120个字符，一个中文占2个字符。一个月内可申请5次修改。请提供可支持命名的材料
	NamingOtherStuff []string `json:"naming_other_stuff,omitempty"` //补充材料，传media id数组，当返回210047时必填
}
type ReturnInfo struct {
	AddressInfo    *AddressInfo `json:"address_info"`    //退货地址
	Mail           string       `json:"mail"`            //邮箱
	CompanyAddress *AddressInfo `json:"company_address"` //公司地址，结构同address_info
}
type AddressInfo struct {
	UserName     string `json:"user_name"`               //收货人姓名
	PostalCode   uint64 `json:"postal_code,omitempty"`   //邮政编码
	ProvinceName string `json:"province_name"`           //省份，格式：广东省 北京市
	CityName     string `json:"city_name"`               //城市，格式：广州市
	CountryName  string `json:"country_name"`            //区，格式：海珠区
	DetailInfo   string `json:"detail_info,omitempty"`   //详细地址
	NationalCode uint64 `json:"national_code,omitempty"` //国家码
	TelNumber    string `json:"tel_number"`              //电话号
}
