package decoration

import (
	"github.com/kaylyu/ministore/component"
	"github.com/kaylyu/ministore/models"
)

const (
	apiSwitch     = "/product/decoration/decoration_service"
	apiExperience = "/product/decoration/decoration_service_experience"
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
