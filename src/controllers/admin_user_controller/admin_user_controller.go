package admin_user_controller

import (
	"apis/common/jwt"
	"apis/common/resp"
	"apis/src/models/admin_user"
	"apis/src/services/admin_user_service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	account := c.Query("account")
	pwd := c.Query("pwd")
	if account == "" || pwd == "" {
		resp.RespGeneralErr(c, "账号密码错误")
		return
	}
	var (
		err error
		ret admin_user.AdminUserFull
	)
	ret.AdminUser, ret.Token, err = admin_user_service.Login(account, pwd)
	if err != nil {
		resp.RespGeneralErr(c, "账号密码错误")
		return
	}
	resp.RespOk(c, ret)
}

func Logout(c *gin.Context) {
	err := admin_user_service.Logout(jwt.GetUid(c))
	if err != nil {
		resp.RespGeneralErr(c, err.Error())
		return
	}
	resp.RespOk(c)
}
