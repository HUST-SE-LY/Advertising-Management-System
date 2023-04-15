package admin

import (
	"backend/api"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

type ManageAdminRouter struct {
}

func (r *ManageAdminRouter) InitManageAdminUserRouter(router *gin.RouterGroup) {
	adminRouter := router.Group("admin").Use(middleware.AdminJWTAuth())
	adminLoginRouter := router.Group("admin") // Without jwt check
	var manageAdminApi = api.ApiGroupApp.ManageApiGroup.ManageAdminApi
	{
		adminRouter.POST("register", manageAdminApi.CreateAdmin)
	}
	{
		adminLoginRouter.POST("login", manageAdminApi.AdminLogin)
	}
}
