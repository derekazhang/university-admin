package course_controller

import (
	"apis/common/jwt"
	"apis/common/resp"
	"apis/src/models/course"
	"apis/src/services/course_service"

	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	id := c.Query("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	if idInt64 <= 0 {
		resp.RespParamErr(c)
		return
	}
	req := course.Course{}
	err := c.BindJSON(&req)
	if err != nil {
		resp.RespParamErr(c)
		return
	}
	req.Uid = jwt.GetUid(c)
	if req.Cover == "" || req.Name == "" || req.Brief == "" {
		resp.RespParamErr(c)
		return
	}
	req.ID = idInt64
	err = course_service.CreateCourse(req)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c)
}

func GetCourseList(ctx *gin.Context) {
	page := ctx.Query("page")
	offset := ctx.Query("offset")
	pageInt, _ := strconv.Atoi(page)
	offsetInt, _ := strconv.Atoi(offset)
	ret, err := course_service.GetCourseList(pageInt, offsetInt)
	if err != nil {
		resp.RespInternalErr(ctx, err.Error())
		return
	}
	resp.RespOk(ctx, ret)
}

func GetUserCourseList(ctx *gin.Context) {

}
