package manage

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type ManageRouterGroup struct {
}

func (a *ManageRouterGroup) Init(router *gin.RouterGroup) {
	manageRouter := router.Group("manage").Use(middleware.AdminJWTAuth())
	var manageAdminApi = api.ApiGroupApp.ManageApiGroup.ManageAdminApi
	var manageCompanyApi = api.ApiGroupApp.ManageApiGroup.ManageCompanyApi
	{
		manageRouter.POST("create-admin", manageAdminApi.CreateAdmin)
		manageRouter.GET("company", manageCompanyApi.GetAllCompanies)
		// TODO(Determine the name of the route)
		manageRouter.GET("company-review", manageCompanyApi.GetAllCompaniesToBeReviewed)
		manageRouter.POST("allow-company-register", manageCompanyApi.AllowRegistrationForCompanies)
	}

}
