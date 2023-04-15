package api

import (
	"backend/api/company"
	"backend/api/manage"
)

type ApiGroup struct {
	ManageApiGroup  manage.ManageApiGroup
	CompanyApiGroup company.CompanyApiGroup
}

var ApiGroupApp = new(ApiGroup)
