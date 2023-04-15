package middleware

import (
	"backend/utils/gin_ext"
	jwt2 "backend/utils/jwt"
	"backend/utils/status"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
)

func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		st := status.Success
		body, err := io.ReadAll(c.Request.Body)
		token := string(body)
		if err != nil {
			st = status.Error
		} else if len(token) == 0 {
			st = status.InvalidParams
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
