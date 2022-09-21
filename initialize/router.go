package initialize

import (
	"net/http"
	"university/global"
	"university/middleware"
	"university/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	//管理后台路由
	manageRouter := router.RouterGroupApp.Manage
	ManageGroup := Router.Group("manage")
	PublicGroup := Router.Group("")

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		//管理后台路由初始化
		manageRouter.InitManageCourseRouter(ManageGroup)
		// manageRouter.InitManageAdminUserRouter(ManageGroup)
		// manageRouter.InitManageGoodsCategoryRouter(ManageGroup)
		// manageRouter.InitManageGoodsInfoRouter(ManageGroup)
		// manageRouter.InitManageCarouselRouter(ManageGroup)
		// manageRouter.InitManageIndexConfigRouter(ManageGroup)
		// manageRouter.InitManageOrderRouter(ManageGroup)
	}
	//小程序路由
	miniRouter := router.RouterGroupApp.Mini.MiniCourseRouter
	MallGroup := Router.Group("mini")
	{
		// 小程序路由初始化
		miniRouter.InitMiniCourseRouter(MallGroup)
		// miniRouter.InitMallGoodsInfoIndexRouter(MallGroup)
		// miniRouter.InitMallGoodsCategoryIndexRouter(MallGroup)
		// miniRouter.InitMallUserRouter(MallGroup)
		// miniRouter.InitMallUserAddressRouter(MallGroup)
		// miniRouter.InitMallShopCartRouter(MallGroup)
		// miniRouter.InitMallOrderRouter(MallGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
