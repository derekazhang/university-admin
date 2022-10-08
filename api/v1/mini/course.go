package mini

import (
	"university/models/common/response"

	"github.com/gin-gonic/gin"
)

type Course struct {
}

func (mini *Course) CourseList(c *gin.Context) {
	response.OkWithData("course list", c)
}
