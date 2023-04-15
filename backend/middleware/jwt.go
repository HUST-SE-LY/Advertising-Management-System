package middleware

import (
	"backend/global"
	"backend/utils/gin_ext"
	jwt2 "backend/utils/jwt"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const bearLength = len(global.JWT_TOKEN_PREFIX)

func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		st := status.Success
		token := c.Request.Header.Get("Authorization")
		if len(token) < bearLength {
			st = status.ErrorAuthCheckTokenInvalid
		} else {
			_, err := jwt2.ParseToken(token)
			if err != nil {
				switch err {
				case jwt.ErrTokenExpired:
					st = status.ErrorAuthCheckTokenTimeout
				default:
					st = status.ErrorAuthCheckTokenFail
				}
			}

		}

		if st != status.Success {
			c.JSON(http.StatusUnauthorized, gin_ext.Response(st, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
