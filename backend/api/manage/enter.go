package manage

import "backend/service"

type ManageApiGroup struct {
	ManageAdminApi
}

var adminService = service.ServiceGroupApp.ManageServiceGroup
