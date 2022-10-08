package mini

import (
	"github.com/gin-gonic/gin"
	v1 "university/api/v1"
	"university/controllers"
)

type Course struct {
}

func (m *Course) InitMiniCourseRouter(Router *gin.RouterGroup) {
	miniCourseRouter := Router.Group("v1")
	var miniCourseApi = v1.Mini.Course
	{
		miniCourseRouter.GET("course", miniCourseApi.CourseList)
	}
	miniCommonRouter := Router.Group("common")
	{
		miniCommonRouter.GET("auth", controllers.MiniProgram.Auth.APISNSSession)
	}

}
