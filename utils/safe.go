package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))
var ADMINTOKEN = "ab0d6a774481e9d3ad478f429be7ece8"

// 生成随机盐
func NewSalt() string {
	return strconv.Itoa(seed.Int())
}

// md5加密
func MD5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}

// 创建密码
func CreatePwd(sourcePwd, salt string) string {
	return MD5(sourcePwd + salt)
}

// 检验密码
func CheckPwd(pwd, salt, dest string) bool {
	return MD5(pwd+salt) == dest
}

//base64 RFC 4648 解密,标准解密
func Base64Decode(origonStr string) string {
	var dest []byte
	dest, err := base64.StdEncoding.DecodeString(origonStr)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(dest)
}

//base64 RFC 4648 加密,标准加密
func Base64Encode(origonStr string) string {
	return base64.StdEncoding.EncodeToString([]byte(origonStr))
}
