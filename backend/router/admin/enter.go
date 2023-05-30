package admin

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)

type AdminRouterGroup struct {
}

func (a *AdminRouterGroup) Init(router *gin.RouterGroup) {
	adminRouter := router.Group("admin") // Without jwt check
	var adminAccountApi = api.ApiGroupApp.AdminApiGroup.AdminAccountApi
	{
		adminRouter.POST("login", adminAccountApi.AdminLogin)
		adminRouter.GET("check_login", adminAccountApi.CheckLogin)
	}
}
