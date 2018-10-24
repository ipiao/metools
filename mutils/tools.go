package mutils

import (
	"regexp"
	"strings"
)

// SnakeName 驼峰转蛇形
func SnakeName(base string) string {
	var r = make([]rune, 0, len(base))
	var b = []rune(base)
	for i := 0; i < len(b); i++ {
		if i > 0 && b[i] >= 'A' && b[i] <= 'Z' {
			r = append(r, '_', b[i]+32)
			continue
		}
		if i == 0 && b[i] >= 'A' && b[i] <= 'Z' {
			r = append(r, b[i]+32)
			continue
		}
		r = append(r, b[i])
	}
	return string(r)
}

// TransFieldName 转换字段名称
// 测试性能不够
func TransFieldName(name string) string {
	return strings.ToLower(regexp.MustCompile(`\B[A-Z]`).ReplaceAllString(name, "_$0"))
}
