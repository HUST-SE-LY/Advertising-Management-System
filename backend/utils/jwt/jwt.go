package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.MapClaims
}

const accessTokenExpiration = 3 * time.Hour

func GenerateToken(username, password string) (string, error) {
	currentTime := time.Now()
	expiredTime := currentTime.Add(accessTokenExpiration)

	claims := Claims{
		EncodedMD5(username),
		EncodedMD5(password),
		jwt.MapClaims{
			"exp": expiredTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}

	return nil, err
}
