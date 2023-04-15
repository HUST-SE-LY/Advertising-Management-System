package company

import "backend/service"

type CompanyApiGroup struct {
	CompanyAccountApi
}

var companyService = service.ServiceGroupApp.CompanyServiceGroup
