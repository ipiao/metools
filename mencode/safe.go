package mencode

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// NewSalt 生成随机盐
func NewSalt() string {
	var seed = rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(seed.Int())
}

// MD5 md5加密
func MD5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}

// Base64Decode base64 RFC 4648 解密,标准解密
func Base64Decode(origonStr string) string {
	var dest []byte
	dest, err := base64.StdEncoding.DecodeString(origonStr)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(dest)
}

// Base64Encode base64 RFC 4648 加密,标准加密
func Base64Encode(origonStr string) string {
	return base64.StdEncoding.EncodeToString([]byte(origonStr))
}
