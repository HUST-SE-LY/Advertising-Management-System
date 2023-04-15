package manage

import "backend/service"

type ManageApiGroup struct {
	ManageAdminApi
	ManageCompanyApi
}

var adminService = service.ServiceGroupApp.ManageServiceGroup
var companyService = service.ServiceGroupApp.CompanyServiceGroup
