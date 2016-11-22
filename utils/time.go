package utils

import "time"

// 解析时间
func ParseTime(t string) time.Time {
	var res, err = time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	if err != nil {
		res, _ = time.ParseInLocation("2006-01-02", t, time.Local)
	}
	return res
}
