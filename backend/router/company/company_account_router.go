package company

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)

type CompanyAccountRouter struct {
}

func (c *CompanyAccountRouter) InitCompanyAccountRouter(router *gin.RouterGroup) {
	//companyRouter := router.Group("company").Use(middleware.CompanyJwtAuth())
	companyRouterWithoutJwt := router.Group("company")
	var companyAccountApi = api.ApiGroupApp.CompanyApiGroup.CompanyAccountApi
	{
		companyRouterWithoutJwt.POST("register", companyAccountApi.RegisterCompany)
		companyRouterWithoutJwt.POST("login", companyAccountApi.CompanyLogin)
	}
}
