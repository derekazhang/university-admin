package services

import (
	"university/services/manage"
	"university/services/mini"
)

type ServiceGroup struct {
	ManageServiceGroup manage.ManageServiceGroup
	MiniServiceGroup   mini.MiniServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
