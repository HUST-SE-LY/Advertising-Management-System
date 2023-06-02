package company

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyRouterGroup struct {
}

func (c *CompanyRouterGroup) Init(router *gin.RouterGroup) {
	companyRouterWithoutJwt := router.Group("company")
	companyRouter := router.Group("company").Use(middleware.CompanyJwtAuth())
	companyAccountApi := api.ApiGroupApp.CompanyApiGroup.CompanyAccountApi
	companyAdvertisementApi := api.ApiGroupApp.CompanyApiGroup.CompanyAdvertisementApi
	{
		companyRouterWithoutJwt.POST("register", companyAccountApi.CompanyRegister)
		companyRouterWithoutJwt.POST("login", companyAccountApi.CompanyLogin)
		companyRouter.POST("recharge", companyAccountApi.CompanyRecharge)
		companyRouter.POST("update-info", companyAccountApi.CompanyUpdateInfo)
		companyRouter.POST("update-pwd", companyAccountApi.CompanyUpdatePwd)
		companyRouter.POST("logout", companyAccountApi.CompanyLogout)
		companyRouter.GET("cancel", companyAccountApi.CompanyCancel)
		companyRouter.PUT("advt/upload", companyAdvertisementApi.CompanyUploadAdvertisement)
		companyRouter.GET("get-info", companyAccountApi.CompanyGetInfo)
		companyRouter.GET("get-records", companyAdvertisementApi.GetRecord)
		companyRouter.GET("get-applications", companyAccountApi.CompanyGetAllApplications)
	}
}
