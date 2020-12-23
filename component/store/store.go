package store

import (
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
func (s *Store) GetStoreInfo(request *models.StoreGetInfoRequest) (response models.StoreGetInfoResponse, err error) {
	err = s.core.HttpPostJson(apiGetStoreInfo, request, &response)
	return
}

/*
上传图片
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/uploadimg.html
*/
//func (s *Store) ImgUpload(request *models.StoreImgUploadRequest) (response models.StoreImgUploadResponse, err error) {
//	err = s.core.HttpPostJson(apiImgUpload, request, &response)
//	return
//}

/*
异步状态查询
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/check_audit_status.html
*/
func (s *Store) RegisterCheckAuditStatus(request *models.StoreRegisterCheckAuditStatusRequest) (response models.StoreRegisterCheckAuditStatusResponse, err error) {
	err = s.core.HttpPostJson(apiRegisterCheckAuditStatus, request, &response)
	return
}

/*
注册小商店账号
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/register_shop.html
*/
func (s *Store) StoreRegister(request *models.StoreRegisterRequest) (response models.StoreRegisterResponse, err error) {
	err = s.core.HttpPostJson(apiRegisterShop, request, &response)
	return
}

/*
提交支付资质
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/submit_merchantinfo.html
*/
func (s *Store) StoreRegisterSubmitMerchantInfo(request *models.StoreRegisterSubmitMerchantInfoRequest) (response models.StoreRegisterSubmitMerchantInfoResponse, err error) {
	err = s.core.HttpPostJson(apiRegisterSubmitMerchantInfo, request, &response)
	return
}

/*
提交小商店基础信息
https://developers.weixin.qq.com/doc/ministore/minishopopencomponent/API/register/submit_basicinfo.html
*/
func (s *Store) StoreRegisterSubmitBasicInfo(request *models.StoreRegisterSubmitBasicInfoRequest) (response models.StoreRegisterSubmitBasicInfoResponse, err error) {
	err = s.core.HttpPostJson(apiRegisterSubmitBasicInfo, request, &response)
	return
}
