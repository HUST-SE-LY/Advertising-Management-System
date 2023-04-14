package api

import "backend/api/manage"

type ApiGroup struct {
	ManageApiGroup manage.ManageApiGroup
}

var ApiGroupApp = new(ApiGroup)
