package manage

import "backend/service"

type ManageApiGroup struct {
	ManageAdminApi
	ManageCompanyApi
	ManageAdvertisementApi
	ManageSpaceApi
}

var adminService = service.ServiceGroupApp.ManageServiceGroup
