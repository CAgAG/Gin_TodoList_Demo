package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	logging "github.com/sirupsen/logrus"
	"time"
)

type TokenData struct {
	Data  interface{} `json:"data"`
	Token string      `json:"token"`
}

var jwtStr = "TodoList_demo"
var jwtSecret = []byte(jwtStr)

type Claims struct {
	Id uint `json:"id"`
	jwt.RegisteredClaims
}

// 签发用户 token
func GenerateToken(id uint) (string, error) {
	claims := Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 有效时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
			Issuer:    "cag",                                              // 签发人
			Subject:   "title",                                            // 主题
			ID:        "1",                                                // JWT ID用于标识该JWT
			Audience:  []string{"someone"},                                // 用户
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 加噪声
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		logging.Info("JWT 生成失败")
		return "", err
	}
	return token, nil
}

// 验证用户 token
func CheckToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		logging.Info("无法验证 JWT")
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, errors.New("验证失败")
}
