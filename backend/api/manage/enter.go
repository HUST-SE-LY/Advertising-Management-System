package manage

import "backend/service"

type ManageApiGroup struct {
	ManageAdminApi
	ManageCompanyApi
	ManageAdvertisementApi
}

var adminService = service.ServiceGroupApp.ManageServiceGroup
