package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

type JwtData struct {
	Subject  string
	Secret   string
	Id       string
	Audience string
	Expires  int64
}

func CreateToken(uid int, ttl int64, secret string) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        fmt.Sprintf("%d", uid),
			Subject:   "",
			Audience:  "",
			ExpiresAt: time.Now().Unix() + ttl,
			Issuer:    "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	//tokenStr, err := token.SignedString([]byte("946356eb-204d-4be3-8eb0-32720a403814"))
	if err != nil {
		return "", fmt.Errorf(`create token err:%v`, err)
	}
	return tokenStr, err
}

func VerifyToken(tokenString, secret string) (ret *CustomClaims, err error) {
	tokenValue, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		err = errors.New(err.Error())
		return
	}
	claims, ok := tokenValue.Claims.(*CustomClaims)
	if !ok {
		err = errors.New("token is invalid")
		return
	}
	ret = claims
	return
}

func GetUid(c *gin.Context) int {
	if val, ex := c.Get("uid"); ex {
		return val.(int)
	}
	return 0
}
