package mini

import (
	v1 "university/api/v1"

	"github.com/gin-gonic/gin"
)

type MiniCourseRouter struct {
}

func (m *MiniCourseRouter) InitMiniCourseRouter(Router *gin.RouterGroup) {
	miniCourseRouter := Router.Group("v1")
	var manageCourseApi = v1.ApiGroupApp.MiniApiGroup.MiniCourseApi
	{
		miniCourseRouter.GET("course", manageCourseApi.CourseList)
	}

}
