package jwt

import (
	"backend/global"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var JwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.MapClaims
}

const accessTokenExpiration = 7 * 24 * time.Hour

func GenerateToken(username, password, role string) (string, error) {
	currentTime := time.Now()
	expiredTime := currentTime.Add(accessTokenExpiration)

	claims := Claims{
		username,
		password,
		jwt.MapClaims{
			"exp":  expiredTime.UnixMilli(),
			"role": role,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	jwtToken := "Bearer " + token
	return jwtToken, err
}

func ParseToken(token string) (*Claims, error) {
	token = strings.Replace(token, global.JWT_TOKEN_PREFIX, "", -1)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}

	return nil, err
}
