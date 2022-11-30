package app

import (
	"electronic-gallery/global"
	"electronic-gallery/internal/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// GetJWTKey jwtKey要转成 []byte才能用
func GetJWTKey() []byte {
	return []byte(global.JwtSetting.Key)
}

func ReleaseToken(user model.User) (string, error) {
	// 设置 token 的有效时间 7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "oceanlearn..tech",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetJWTKey())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	// 从token解析claims
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return GetJWTKey(), nil
	})

	return token, claims, err
}
