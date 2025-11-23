package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var UserPointsAccount = new(UPA)

type UPA struct{}

// Init 初始化 用户积分账户 路由信息
func (r *UPA) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("UPA").Use(middleware.OperationRecord())
		group.POST("createUserPointsAccount", apiUserPointsAccount.CreateUserPointsAccount)             // 新建用户积分账户
		group.DELETE("deleteUserPointsAccount", apiUserPointsAccount.DeleteUserPointsAccount)           // 删除用户积分账户
		group.DELETE("deleteUserPointsAccountByIds", apiUserPointsAccount.DeleteUserPointsAccountByIds) // 批量删除用户积分账户
		group.PUT("updateUserPointsAccount", apiUserPointsAccount.UpdateUserPointsAccount)              // 更新用户积分账户
	}
	{
		group := private.Group("UPA")
		group.GET("findUserPointsAccount", apiUserPointsAccount.FindUserPointsAccount)       // 根据ID获取用户积分账户
		group.GET("getUserPointsAccountList", apiUserPointsAccount.GetUserPointsAccountList) // 获取用户积分账户列表
	}
	{
		group := public.Group("UPA")
		group.GET("getUserPointsAccountDataSource", apiUserPointsAccount.GetUserPointsAccountDataSource) // 获取用户积分账户数据源
		group.GET("getUserPointsAccountPublic", apiUserPointsAccount.GetUserPointsAccountPublic)         // 用户积分账户开放接口
	}
	{
		group := private.Group("UPA")
		group.GET("getMyPointsBalance", apiUserPointsAccount.GetMyPointsBalance) // 我的积分余额
	}
}
