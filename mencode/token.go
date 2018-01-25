package mencode

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ipiao/metools/mutils"
)

var (
	// VerifyKey 签名公钥
	VerifyKey []byte
	// SignKey 签名私钥
	SignKey []byte
)

// MyClaims 自定义
type MyClaims struct {
	Info map[string]interface{}
	jwt.StandardClaims
}

// Int64 获取int值
func (claim *MyClaims) Int64(key string) (int64, error) {
	r, ok := claim.Info[key]
	if !ok {
		return 0, fmt.Errorf("key named %s does not exists", key)
	}
	res, err := mutils.IntfaceToInt64(r)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Bool 获取bool值
func (claim *MyClaims) Bool(key string) (bool, error) {
	r, ok := claim.Info[key]
	if !ok {
		return false, fmt.Errorf("key named %s does not exists", key)
	}
	res, ok2 := r.(bool)
	if !ok2 {
		return false, fmt.Errorf("cannot parse interface %s to bool", key)
	}
	return res, nil
}

// String 获取string值
func (claim *MyClaims) String(key string) (string, error) {
	r, ok := claim.Info[key]
	if !ok {
		return "", fmt.Errorf("key named %s does not exists", key)
	}
	res, ok2 := r.(string)
	if !ok2 {
		return "", fmt.Errorf("cannot parse interface %s to string", key)
	}
	return res, nil
}

// LoadKeysForAccessToken 读取签名
func LoadKeysForAccessToken(privFile string, pubFile string) {
	var err error
	SignKey, err = ioutil.ReadFile(privFile)

	if err != nil {
		log.Fatal("Fatal error: failed to read priv key ")
		return
	}
	VerifyKey, err = ioutil.ReadFile(pubFile)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

// GenerateNewAccessToken 生成token
func GenerateNewAccessToken(info map[string]interface{}, duration int64) (string, error) {
	claims := MyClaims{
		info,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(time.Second)*duration,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(SignKey) // sign the token
	return tokenString, err
}

// ParseAccessToken 解析token
func ParseAccessToken(tokenString string) (*MyClaims, error) {

	var claims = new(MyClaims)
	var ok bool
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SignKey, nil
	})
	if err != nil {
		return claims, err
	}
	if claims, ok = token.Claims.(*MyClaims); !ok || !token.Valid {
		return claims, fmt.Errorf("token错误：%s", err.Error())
	}
	return claims, err
}
