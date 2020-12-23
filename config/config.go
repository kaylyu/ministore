package config

import "time"

/*
小商店配置
*/
type Config struct {
	AccessToken string
	Timeout     time.Duration //请求超时设置
	LogPrefix   string        //日志前缀
}
