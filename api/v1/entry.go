package v1

import (
	"university/api/v1/manage"
	"university/api/v1/mini"
)

type ApiGroup struct {
	ManageApiGroup manage.ManageGroup
	MiniApiGroup   mini.MiniGroup
}

var ApiGroupApp = new(ApiGroup)
