package store

import (
	"fmt"
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiGetStoreInfo               = "/product/store/get_info"
	apiImgUpload                  = "/product/img/upload"
	apiRegisterCheckAuditStatus   = "/product/register/check_audit_status"
	apiRegisterShop               = "/product/register/register_shop"
	apiRegisterSubmitMerchantInfo = "/product/register/submit_merchantinfo"
	apiRegisterSubmitBasicInfo    = "/product/register/submit_basicinfo"
)

//Store
type Store struct {
	core *component.Core
}

/*
创建店铺实例
*/
func NewStore(core *component.Core) *Store {
	return &Store{core}
}

/*
获取基本信息
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/store/get_store_info.html
*/
func (s *Store) GetStoreInfo(request *models.StoreGetInfoRequest) (response models.StoreGetInfoResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiGetStoreInfo, request, &response)
	return
}

/*
上传图片
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/uploadimg.html
*/
func (s *Store) ImgUpload(request *models.StoreImgUploadRequest) (response models.StoreImgUploadResponse, body string, err error) {
	body, err = s.core.Upload(fmt.Sprintf("%s?&height=%d&width=%d", apiImgUpload, request.Height, request.Width), request.Media, &response)
	return
}

/*
异步状态查询
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/check_audit_status.html
*/
func (s *Store) RegisterCheckAuditStatus(request *models.StoreRegisterCheckAuditStatusRequest, componentAccessToken ...string) (response models.StoreRegisterCheckAuditStatusResponse, body string, err error) {
	//记录
	accessToken := request.GetAccessToken()
	//当调用完注册接口，查询注册好的appid时，使用component access token。当查询到appid后，需要需要查询审核状态以及调用其它接口时，使用authorizer access token
	if len(componentAccessToken) > 0 && len(componentAccessToken[0]) > 0 {
		request.SetAccessToken(componentAccessToken[0])
	}
	body, err = s.core.HttpPostJson(apiRegisterCheckAuditStatus, request, &response)
	//执行完毕后还原access token
	request.SetAccessToken(accessToken)
	return
}

/*
注册小商店账号
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/register_shop.html
*/
func (s *Store) StoreRegister(request *models.StoreRegisterRequest, componentAccessToken string) (response models.StoreRegisterResponse, body string, err error) {
	//设置
	request.SetComponentAccessToken(componentAccessToken)
	body, err = s.core.HttpPostJson(apiRegisterShop, request, &response)
	return
}

/*
提交支付资质
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/submit_merchantinfo.html
*/
func (s *Store) StoreRegisterSubmitMerchantInfo(request *models.StoreRegisterSubmitMerchantInfoRequest) (response models.StoreRegisterSubmitMerchantInfoResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiRegisterSubmitMerchantInfo, request, &response)
	return
}

/*
提交小商店基础信息
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/submit_basicinfo.html
*/
func (s *Store) StoreRegisterSubmitBasicInfo(request *models.StoreRegisterSubmitBasicInfoRequest) (response models.StoreRegisterSubmitBasicInfoResponse, body string, err error) {
	body, err = s.core.HttpPostJson(apiRegisterSubmitBasicInfo, request, &response)
	return
}
