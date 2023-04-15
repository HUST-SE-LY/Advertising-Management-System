package router

import (
	"backend/router/admin"
	"backend/router/company"
)

type RouterGroup struct {
	Admin   admin.AdminRouterGroup
	Company company.CompanyRouterGroup
}

var RouterGroupApp = new(RouterGroup)
