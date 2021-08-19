package util

import (
	"time"

	"github.com/Coder-stars/gin-example/pkg/setting"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	UserId   string `json:"userid"`
	jwt.StandardClaims
}

func GenerateToken(username string, userid string) (string, error) {
	nowTime := time.Now()
	expiredTime := nowTime.Add(3 * time.Hour)

	Claims := Claims{
		username,
		userid,
		jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    "gin-example",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
