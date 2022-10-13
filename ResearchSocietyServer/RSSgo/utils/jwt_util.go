package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

type JwtUtil struct {
}

var viperUtil ViperUtil

func (JwtUtil) ReleaseToken(userID uint) string {
	jwtConfig := viperUtil.Read("config.yml", "jwt")
	secret := jwtConfig["secret"].(string)
	expire := jwtConfig["expire"].(int)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(expire),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hkc.ink",
			Subject:   "Authorization",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	authorization, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("token error:" + err.Error())
		return "error"
	}

	return authorization
}

func (JwtUtil) ParseToken(authorization string) (*jwt.Token, *Claims, error) {
	jwtConfig := viperUtil.Read("config.yml", "jwt")
	secret := jwtConfig["secret"].(string)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(authorization, claims, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	return token, claims, err
}
