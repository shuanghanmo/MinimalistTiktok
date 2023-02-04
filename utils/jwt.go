package utils

import (
	gojwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtKey []byte

func init() {
	jwtKey = []byte(os.Getenv("douyin_simple"))
}

type Claims struct {
	Uid int64
	gojwt.StandardClaims
}

// Award 生成Token
func Award(uid int64) (string, error) {
	// 过期时间设置为3小时
	expireTime := time.Now().Add(3 * time.Hour)
	claims := &Claims{
		Uid: uid,
		StandardClaims: gojwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	// 生成token
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*gojwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := gojwt.ParseWithClaims(tokenStr, claims, func(t *gojwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, err
}
