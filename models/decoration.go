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

//检查装修服务使用状态
type DecorationCheckStatusRequest struct {
	APIRequest
	ServiceId uint64 `json:"service_id"`
}
type DecorationCheckStatusResponse struct {
	APIResponse
	Status uint64 `json:"status"` //1表示已启用，2表示未启用
}
