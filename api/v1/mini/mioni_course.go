package mini

import (
	"university/models/common/response"

	"github.com/gin-gonic/gin"
)

type MiniCourseApi struct {
}

func (mini *MiniCourseApi) CourseList(c *gin.Context) {
	response.OkWithData("coures list", c)
}
