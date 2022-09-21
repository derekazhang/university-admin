package mini

import "university/service"

type MiniGroup struct {
	MiniCourseApi
}

var miniCourseService = service.ServiceGroupApp.MiniServiceGroup.MiniCourseService
