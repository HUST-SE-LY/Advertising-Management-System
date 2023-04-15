package api

import (
	"backend/api/admin"
	"backend/api/company"
	"backend/api/manage"
)

type ApiGroup struct {
	AdminApiGroup   admin.AdminApiGroup
	ManageApiGroup  manage.ManageApiGroup
	CompanyApiGroup company.CompanyApiGroup
}

var ApiGroupApp = new(ApiGroup)
