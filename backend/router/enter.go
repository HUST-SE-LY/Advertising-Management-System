package router

import "backend/router/admin"

type RouterGroup struct {
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
