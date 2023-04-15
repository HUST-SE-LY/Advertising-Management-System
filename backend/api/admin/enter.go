package admin

import "backend/service"

type AdminApiGroup struct {
	AdminAccountApi
}

var adminService = service.ServiceGroupApp.ManageServiceGroup
