package jwt

import (
	"github.com/form3tech-oss/jwt-go"

	"time"
)

type JWTClaims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

// 密钥和过期时间
var secret = []byte("minitok-secret")

const ExpireDuration = 300 * time.Second

// GenToken 生成 JWT token
func GenToken(userid int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(ExpireDuration)
	// 创建
	claims := JWTClaims{
		ID:       userid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	// 用指定加密方法创建 claim 对象并用密钥加密
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return token, err
}

// ParseToken 解析 token 得到 JWTClaims 对象
func ParseToken(token string) (*JWTClaims, error) {
	// 传入密钥解析 token, 得到 JWT Web Token 对象
	jwtWebToken, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || jwtWebToken == nil || jwtWebToken.Valid != true {
		return nil, err
	}

	// 返回私有声明
	if claims, ok := jwtWebToken.Claims.(*JWTClaims); ok {
		return claims, nil
	}

	return nil, err
}
