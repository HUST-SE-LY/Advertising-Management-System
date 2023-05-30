package admin

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type AdminRouterGroup struct {
}

func (a *AdminRouterGroup) Init(router *gin.RouterGroup) {
	adminRouter := router.Group("admin") // Without jwt check
	adminRouterWithJwt := router.Group("admin").Use(middleware.CompanyJwtAuth())
	var adminAccountApi = api.ApiGroupApp.AdminApiGroup.AdminAccountApi
	{
		adminRouter.POST("login", adminAccountApi.AdminLogin)
		adminRouterWithJwt.GET("check_login", adminAccountApi.CheckLogin)
	}
}
