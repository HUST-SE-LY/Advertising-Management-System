package admin

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type AdminRouterGroup struct {
}

func (a *AdminRouterGroup) Init(router *gin.RouterGroup) {
	adminRouter := router.Group("admin").Use(middleware.AdminJWTAuth())
	adminLoginRouter := router.Group("admin") // Without jwt check
	var manageAdminApi = api.ApiGroupApp.ManageApiGroup.ManageAdminApi
	var manageCompanyApi = api.ApiGroupApp.ManageApiGroup.ManageCompanyApi
	{
		adminRouter.POST("register", manageAdminApi.CreateAdmin)
		// TODO(Determine the name of the route)
		adminRouter.GET("company_review", manageCompanyApi.ReadAllCompaniesToBeReviewed)
	}

	{
		adminLoginRouter.POST("login", manageAdminApi.AdminLogin)
	}
}
