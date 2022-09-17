package routers

import (
	"apis/src/controllers/course_controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisteClient(r *gin.Engine) {
	clientRouter := r.Group("client")
	{
		clientRouter.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "Hello World",
			})
		})
	}

	clientRouter.GET("crouse", course_controller.GetCourseList)
	clientRouter.POST("crouse", course_controller.CreateCourse)

}
