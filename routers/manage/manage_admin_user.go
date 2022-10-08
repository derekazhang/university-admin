package manage

import (
	v1 "university/api/v1"

	"github.com/gin-gonic/gin"
)

type AdminUser struct {
}

func (m *AdminUser) InitManageAdminUserRouter(Router *gin.RouterGroup) {
	manageAdminUserRouter := Router.Group("v1")

	var manageAdminUserApi = v1.Manage.AdminUser
	{
		manageAdminUserRouter.POST("admin/login", manageAdminUserApi.AdminLogin)
		manageAdminUserRouter.POST("admin/new", manageAdminUserApi.CreateAdminUser)
	}

}
