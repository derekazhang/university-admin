package manage

import (
	"university/global"
	"university/models/common/response"
	"university/models/manage"
	"university/models/manage/request"

	"university/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageAdminUserApi struct {
}

func (m *ManageAdminUserApi) CreateAdminUser(c *gin.Context) {
	var params request.ManageAdminParam
	_ = c.ShouldBindJSON(&params)

	if err := utils.Verify(params, utils.AdminUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	manageAdminUser := manage.ManageAdminUser{
		LoginUserName: params.LoginUserName,
		NickName:      params.NickName,
		LoginPassword: utils.MD5V([]byte(params.LoginPassword)),
	}
	if err := manageAdminUserService.CreateManageAdminUser(manageAdminUser); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

}
