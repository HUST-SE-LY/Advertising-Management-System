package initialize

import (
	"backend/global"
	"backend/middleware"
	"backend/router"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {

	r.Use(gin.Logger())
	r.Use(middleware.Cors())

	rootRouterGroup := r.Group("")

	router.RouterGroupApp.Init(rootRouterGroup)

	r.Static("/public", *global.GVA_FILE_SETTING)

}
