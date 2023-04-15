package initialize

import (
	"backend/middleware"
	"backend/router"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {

	r.Use(gin.Logger())
	r.Use(middleware.Cors())

	rootRouterGroup := r.Group("")

	adminRouterGroup := router.RouterGroupApp.Admin
	companyRouterGroup := router.RouterGroupApp.Company
	{
		adminRouterGroup.Init(rootRouterGroup)
		companyRouterGroup.InitCompanyAccountRouter(rootRouterGroup)
	}

}
