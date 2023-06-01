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
	adminRouterWithJwt := router.Group("admin").Use(middleware.AdminJWTAuth())
	var adminAccountApi = api.ApiGroupApp.AdminApiGroup.AdminAccountApi
	{
		adminRouter.POST("login", adminAccountApi.AdminLogin)
		adminRouter.POST("logout", adminAccountApi.AdminLogout)
		adminRouterWithJwt.GET("check_login", adminAccountApi.CheckLogin)
	}
}
