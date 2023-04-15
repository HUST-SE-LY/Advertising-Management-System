package middleware

import (
	"backend/global"
	"backend/service"
	"backend/utils/gin_ext"
	jwt2 "backend/utils/jwt"
	"backend/utils/status"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const bearLength = len(global.JWT_TOKEN_PREFIX)

var adminAccountService = service.ServiceGroupApp.AdminServiceGroup.AdminAccountService
var companyAccountService = service.ServiceGroupApp.CompanyServiceGroup.CompanyAccountService

func extractAndCheckToken(c *gin.Context) (token string, claims *jwt2.Claims, err error) {
	token = c.Request.Header.Get("Authorization")
	if len(token) < bearLength {
		err = &status.ErrorAuthCheckTokenInvalid
	} else {
		claims, err = jwt2.ParseToken(token)
		if err != nil {
			switch err {
			case jwt.ErrTokenExpired:
				err = &status.ErrorAuthCheckTokenExpired
			default:
				err = &status.ErrorAuthCheckTokenFail
			}
		}
	}
	return
}

func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _, err := extractAndCheckToken(c)
		if err != nil {
			if errors.Is(err, status.ErrorAuthCheckTokenExpired) {
				err := adminAccountService.DeleteAdminToken(token)
				if err != nil {
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin_ext.Response(err, nil))
			c.Abort()
			return
		}

		if _, err := adminAccountService.FindAdminToken(token); err != nil {
			err = status.ErrorAuthCheckTokenNotFound
			c.JSON(http.StatusUnauthorized, gin_ext.Response(err, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}

func CompanyJwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _, err := extractAndCheckToken(c)
		if err != nil {
			if errors.Is(err, status.ErrorAuthCheckTokenExpired) {
				err := companyAccountService.DeleteCompanyToken(token)
				if err != nil {
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin_ext.Response(err, nil))
			c.Abort()
			return
		}

		if _, err := companyAccountService.FindCompanyToken(token); err != nil {
			err = status.ErrorAuthCheckTokenNotFound
			c.JSON(http.StatusUnauthorized, gin_ext.Response(err, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
