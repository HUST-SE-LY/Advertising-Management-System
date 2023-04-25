package router

import (
	"backend/router/admin"
	"backend/router/company"
	"backend/router/manage"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	Admin   admin.AdminRouterGroup
	Company company.CompanyRouterGroup
	Manage  manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)

func (r RouterGroup) Init(rootRouterGroup *gin.RouterGroup) {
	r.Admin.Init(rootRouterGroup)
	r.Company.Init(rootRouterGroup)
	r.Manage.Init(rootRouterGroup)
}
