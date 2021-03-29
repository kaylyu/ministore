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

//发起压测
type DecorationPressureTestRequest struct {
	APIRequest
}
type DecorationPressureTestResponse struct {
	APIResponse
	PressureId uint64 `json:"pressure_id"` //压测id
}

//查询压测状态
type DecorationPressureTestStatusRequest struct {
	APIRequest
}
type DecorationPressureTestStatusResponse struct {
	APIResponse
	Status uint64 `json:"status"` //1：排队中；2：压测中；3：压测失败 4：压测结束
}

//获取压测报告
type DecorationGetPressureTestReportRequest struct {
	APIRequest
	PressureId uint64 `json:"pressure_id"` //压测id
}
type DecorationGetPressureTestReportResponse struct {
	APIResponse
	BlankpagePencent uint64         `json:"blankpage_pencent"` //打开页面的白页率
	AverTimeCost     uint64         `json:"aver_time_cost"`    //打开页面的平均耗时, 单位ms
	MaxTimeCost      uint64         `json:"max_time_cost"`     //打开页面的最大耗时, 单位ms
	TotalLaunchCnt   uint64         `json:"total_launch_cnt"`  //共打开了多少次该页
	TotalRequestCnt  uint64         `json:"total_request_cnt"` //共发起多少次网络请求
	Appid            string         `json:"appid"`             //小程序id
	RunTime          uint64         `json:"run_time"`          //压测时间, 单位s
	PressureId       uint64         `json:"pressure_id"`       //压测ID
	NetworkList      []*NetworkList `json:"network_list"`      //各网络请求的统计
}
type NetworkList struct {
	Path            string `json:"path"`              //网络请求path
	AverTimeCost    uint64 `json:"aver_time_cost"`    //该path的平均耗时, 单位ms
	MaxTimeCost     uint64 `json:"max_time_cost"`     //该path的最大耗时，单位ms
	TotalRequestCnt uint64 `json:"total_request_cnt"` //共发起多少次网络请求
	SuccPercent     uint64 `json:"succ_percent"`      //请求该path的成功率
}
