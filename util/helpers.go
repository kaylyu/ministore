package util

import (
	"encoding/json"
	"strings"
	"time"
)

//模拟三元操作
func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}

//格式化时间
func Format(t *time.Time) (v string) {
	if t != nil {
		return t.Format("2006-01-02 15:04:05")
	}
	return
}

//时间字符串转换成时间对象
func StrToTime(s string) *time.Time {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", s, loc)

	return &t
}

//转JSON
func JsonEncode(v interface{}) string {
	if v == nil {
		return "{}"
	}

	j, e := json.Marshal(v)
	if e != nil || strings.Compare(string(j), "null") == 0 {
		return "{}"
	}

	return string(j)
}

//转JSON
func JsonDecode(s string, out interface{}) error {
	e := json.Unmarshal([]byte(s), out)

	if e != nil {
		return e
	}

	return nil
}
