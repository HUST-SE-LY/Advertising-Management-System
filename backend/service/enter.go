package service

import (
	"backend/service/company"
	"backend/service/manage"
)

type ServiceGroup struct {
	ManageServiceGroup  manage.ManageServiceGroup
	CompanyServiceGroup company.CompanyServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
