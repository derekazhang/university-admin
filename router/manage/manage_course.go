package manage

import (
	v1 "university/api/v1"

	"github.com/gin-gonic/gin"
)

type ManageCourseRouter struct {
}

func (m *ManageCourseRouter) InitManageCourseRouter(Router *gin.RouterGroup) {
	manageCourseRouter := Router.Group("v1")
	var manageCourseApi = v1.ApiGroupApp.ManageApiGroup.ManageCourseApi
	{
		manageCourseRouter.GET("course", manageCourseApi.CreateCourse)
	}

}
