package company

import "backend/service"

type CompanyApiGroup struct {
	CompanyAccountApi
	CompanyAdvertisementApi
}

var companyService = service.ServiceGroupApp.CompanyServiceGroup
