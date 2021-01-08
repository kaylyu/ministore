package config

import "time"

/*
小商店配置
*/
type Config struct {
	AccessToken          string        //小店token
	ComponentAccessToken string        //第三方平台token
	Timeout              time.Duration //请求超时设置
	LogPrefix            string        //日志前缀
	Debug                bool          //是否开启日志
}
