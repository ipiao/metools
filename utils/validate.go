package utils

import "regexp"

// 验证是否是手机号
func IsPhone(s string) bool {
	regExp := `^0?(13|14|15|18)[0-9]{9}$`
	reg := regexp.MustCompile(regExp)
	return reg.MatchString(s)
}

// 验证是否是邮箱
func IsEmail(s string) bool {
	regExp := `^\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}$`
	reg := regexp.MustCompile(regExp)
	return reg.MatchString(s)
}
