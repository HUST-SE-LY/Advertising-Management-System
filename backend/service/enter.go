package service

import (
	"backend/service/admin"
	"backend/service/company"
	"backend/service/manage"
)

type ServiceGroup struct {
	AdminServiceGroup   admin.AdminServiceGroup
	CompanyServiceGroup company.CompanyServiceGroup
	ManageServiceGroup  manage.ManageServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
