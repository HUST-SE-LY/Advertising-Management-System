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

	router.RouterGroupApp.Init(rootRouterGroup)

	r.Static("public", "./public")

}
