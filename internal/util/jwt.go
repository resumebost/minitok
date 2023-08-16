package util

import (
	"github.com/form3tech-oss/jwt-go"
	"time"
)

type MyClaims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

var secret = []byte("mini-tok324")

const ExpireDuration = 300 * time.Second

func GenerateToken(id int64, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(ExpireDuration)
	claims := MyClaims{
		ID:       id,
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return token, err
}

func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

//TODO 中间件？
