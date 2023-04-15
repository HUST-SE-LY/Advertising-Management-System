package company

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)

type CompanyRouterGroup struct {
}

func (c *CompanyRouterGroup) Init(router *gin.RouterGroup) {
	companyRouterWithoutJwt := router.Group("company")
	var companyAccountApi = api.ApiGroupApp.CompanyApiGroup.CompanyAccountApi
	{
		companyRouterWithoutJwt.POST("register", companyAccountApi.RegisterCompany)
		companyRouterWithoutJwt.POST("login", companyAccountApi.CompanyLogin)
	}
}
