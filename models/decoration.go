package models

//装修开启/停用
type DecorationSwitchRequest struct {
	APIRequest
	SwitchStatus SwitchStatus `json:"switch_status"` //1代表启用，2代表停用
}
type DecorationSwitchResponse struct {
	APIResponse
}

//装修部署体验版
type DecorationExperienceRequest struct {
	APIRequest
	ServiceId uint64 `json:"service_id"`
}
type DecorationExperienceResponse struct {
	APIResponse
}
