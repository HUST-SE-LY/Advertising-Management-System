package manage

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type ManageRouterGroup struct {
}

func (a *ManageRouterGroup) Init(router *gin.RouterGroup) {
	adminRouter := router.Group("manage").Use(middleware.AdminJWTAuth())
	var manageAdminApi = api.ApiGroupApp.ManageApiGroup.ManageAdminApi
	var manageCompanyApi = api.ApiGroupApp.ManageApiGroup.ManageCompanyApi
	{
		adminRouter.POST("create-admin", manageAdminApi.CreateAdmin)
		// TODO(Determine the name of the route)
		adminRouter.GET("company-review", manageCompanyApi.ReadAllCompaniesToBeReviewed)
		adminRouter.POST("allow-company-register", manageCompanyApi.AllowRegistrationForCompanies)
	}

}
