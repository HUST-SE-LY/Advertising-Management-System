package service

import "backend/service/manage"

type ServiceGroup struct {
	ManageServiceGroup manage.ManageServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
