package router

import (
	"university/router/manage"
	"university/router/mini"
)

type RouterGroup struct {
	Mini   mini.MiniRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
