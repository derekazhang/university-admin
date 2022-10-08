package manage

import (
	v1 "university/api/v1"
	"university/middleware"

	"github.com/gin-gonic/gin"
)

type Course struct {
}

func (m *Course) InitManageCourseRouter(Router *gin.RouterGroup) {
	manageCourseRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var manageCourseApi = v1.Manage.Course
	{
		manageCourseRouter.GET("course", manageCourseApi.CreateCourse)
	}
}
