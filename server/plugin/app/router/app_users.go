package router

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
)

var AppUsers = new(appUsers)

type appUsers struct{}

// Init 初始化 appUsers表 路由信息
func (r *appUsers) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("appUsers").Use(middleware.OperationRecord())
		group.POST("createAppUsers", apiAppUsers.CreateAppUsers)             // 新建appUsers表
		group.DELETE("deleteAppUsers", apiAppUsers.DeleteAppUsers)           // 删除appUsers表
		group.DELETE("deleteAppUsersByIds", apiAppUsers.DeleteAppUsersByIds) // 批量删除appUsers表
		group.PUT("updateAppUsers", apiAppUsers.UpdateAppUsers)              // 更新appUsers表
	}
	{
		group := private.Group("appUsers")
		group.GET("findAppUsers", apiAppUsers.FindAppUsers)       // 根据ID获取appUsers表
		group.GET("getAppUsersList", apiAppUsers.GetAppUsersList) // 获取appUsers表列表
		group.GET("getUserInfo", apiAppUsers.GetUserInfo)         // 获取当前登录用户信息
		group.GET("getDashboard", apiAppUsers.GetDashboard)       // 获取APP用户首页/个人中心聚合数据
		group.GET("getInviteCount", apiAppUsers.GetInviteCount)   // 获取邀请统计
		group.GET("getSubUsers", apiAppUsers.GetSubUsers)         // 获取下级用户列表
		group.POST("logout", apiAppUsers.Logout)                  //退出登录
	}
	{
		group := public.Group("appUsers")
		group.POST("sendAppUserEmail", apiAppUsers.SendVerificationCode) // App用户管理开放发送邮件接口
		group.POST("register", apiAppUsers.Register)                     // 用户注册
		group.POST("login", apiAppUsers.Login)                           // 用户登录
		group.POST("resetPassword", apiAppUsers.ResetPassword)           // 重置密码
	}
}
