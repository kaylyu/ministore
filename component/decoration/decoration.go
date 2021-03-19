package decoration

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiSwitch      = "/product/decoration/decoration_service"
	apiExperience  = "/product/decoration/decoration_service_experience"
	apiCheckStatus = "/product/decoration/wxaproduct_decoration_check_status"
)

//装修
type Decoration struct {
	core *component.Core
}

/*
创建装修实例
*/
func NewDecoration(core *component.Core) *Decoration {
	return &Decoration{core}
}

/*
启用或停用装修
https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/value_added_services/decoration.html
*/
func (d *Decoration) DecorationSwitch(request *models.DecorationSwitchRequest) (response models.DecorationSwitchResponse, body string, err error) {
	body, err = d.core.HttpPostJson(apiSwitch, request, &response)
	return
}

/*
部署体验版
https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/value_added_services/decoration.html
*/
func (d *Decoration) DecorationExperience(request *models.DecorationExperienceRequest) (response models.DecorationExperienceResponse, body string, err error) {
	body, err = d.core.HttpPostJson(apiExperience, request, &response)
	return
}

/*
检查装修服务使用状态
https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent/API/service/decoration/decoration_check_status.html
*/
func (d *Decoration) DecorationCheckStatus(request *models.DecorationCheckStatusRequest) (response models.DecorationCheckStatusResponse, body string, err error) {
	body, err = d.core.HttpPostJson(apiCheckStatus, request, &response)
	return
}
