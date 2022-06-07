package middleware

import (
	"firebackend/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 证书签名密钥
var jwtKey = []byte("firecapital")

/// 定义生成token的方法
func GenerateToken(u model.LoginUser) (string, error) {
	// 定义过期时间,7天后过期
	expirationTime := time.Now().Add(3 * 24 * time.Hour)
	claims := &model.MyClaims{
		UserId:   u.ID,
		Username: u.UserName,
		Password: u.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 发布时间
			Subject:   "token",               // 主题
			Issuer:    "fire",                // 发布者
		},
	}
	// 注意单词别写错了
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
