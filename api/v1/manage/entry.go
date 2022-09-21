package manage

import "university/service"

type ManageGroup struct {
	ManageAdminUserApi
	ManageCourseApi
}

var manageAdminUserService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserService
var manageCourseService = service.ServiceGroupApp.ManageServiceGroup.ManageCourseService
