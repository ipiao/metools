package datetool

import (
	"fmt"
	"time"
)

// ParseTime 解析时间
func ParseTime(s string) (time.Time, error) {
	var layout = ""
	l := len(s)
	switch l {
	case 19:
		s1 := s[4:5]
		layout = fmt.Sprintf("2006%s01%s02 15:04:05", s1, s1)
	case 10:
		s1 := s[4:5]
		layout = fmt.Sprintf("2006%s01%s02", s1, s1)
	case 8:
		layout = "15:04:05"
	case 7:
		s1 := s[4:5]
		layout = fmt.Sprintf("2006%s01", s1)
	}
	if layout == "" {
		return time.Time{}, fmt.Errorf("不支持的时间格式:%s", s)
	}
	v, err := time.ParseInLocation(layout, s, time.Local)
	return v, err
}
