package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var UserPointsLog = new(UPL)

type UPL struct {}

// Init 初始化 用户积分流水 路由信息
func (r *UPL) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("UPL").Use(middleware.OperationRecord())
		group.POST("createUserPointsLog", apiUserPointsLog.CreateUserPointsLog)   // 新建用户积分流水
		group.DELETE("deleteUserPointsLog", apiUserPointsLog.DeleteUserPointsLog) // 删除用户积分流水
		group.DELETE("deleteUserPointsLogByIds", apiUserPointsLog.DeleteUserPointsLogByIds) // 批量删除用户积分流水
		group.PUT("updateUserPointsLog", apiUserPointsLog.UpdateUserPointsLog)    // 更新用户积分流水
	}
	{
	    group := private.Group("UPL")
		group.GET("findUserPointsLog", apiUserPointsLog.FindUserPointsLog)        // 根据ID获取用户积分流水
		group.GET("getUserPointsLogList", apiUserPointsLog.GetUserPointsLogList)  // 获取用户积分流水列表
	}
	{
	    group := public.Group("UPL")
	    group.GET("getUserPointsLogDataSource", apiUserPointsLog.GetUserPointsLogDataSource)  // 获取用户积分流水数据源
	    group.GET("getUserPointsLogPublic", apiUserPointsLog.GetUserPointsLogPublic)  // 用户积分流水开放接口
	}
}
