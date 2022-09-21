package service

import (
	"university/service/manage"
	"university/service/mini"
)

type ServiceGroup struct {
	ManageServiceGroup manage.ManageServiceGroup
	MiniServiceGroup   mini.MiniServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
