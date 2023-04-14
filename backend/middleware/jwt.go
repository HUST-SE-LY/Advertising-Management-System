package jwt

import (
	"backend/pkg/constants/status"
	"backend/pkg/response"
	"backend/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		code := status.SUCCESS
		body, err := io.ReadAll(context.Request.Body)
		token := string(body)
		if err != nil {
			code = status.ERROR
		} else if len(token) == 0 {
			code = status.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err {
				case jwt.ErrTokenExpired:
					code = status.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = status.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != status.SUCCESS {
			context.JSON(http.StatusUnauthorized, response.Response(code, "", nil))
			context.Abort()
			return
		}

		context.Next()
	}
}
