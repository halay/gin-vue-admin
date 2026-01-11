package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var AppAgentTransactions = new(appAgentTransactions)

type appAgentTransactions struct {}

// Init 初始化 appAgentTransactions表 路由信息
func (r *appAgentTransactions) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("appAgentTransactions").Use(middleware.OperationRecord())
		group.POST("createAppAgentTransactions", apiAppAgentTransactions.CreateAppAgentTransactions)   // 新建appAgentTransactions表
		group.DELETE("deleteAppAgentTransactions", apiAppAgentTransactions.DeleteAppAgentTransactions) // 删除appAgentTransactions表
		group.DELETE("deleteAppAgentTransactionsByIds", apiAppAgentTransactions.DeleteAppAgentTransactionsByIds) // 批量删除appAgentTransactions表
		group.PUT("updateAppAgentTransactions", apiAppAgentTransactions.UpdateAppAgentTransactions)    // 更新appAgentTransactions表
	}
	{
	    group := private.Group("appAgentTransactions")
		group.GET("findAppAgentTransactions", apiAppAgentTransactions.FindAppAgentTransactions)        // 根据ID获取appAgentTransactions表
		group.GET("getAppAgentTransactionsList", apiAppAgentTransactions.GetAppAgentTransactionsList)  // 获取appAgentTransactions表列表
		group.GET("getMyAgentTransactionsList", apiAppAgentTransactions.GetMyAgentTransactionsList)  // APP用户获取自己的分润明细
	}
	{
	    group := public.Group("appAgentTransactions")
		_ = group
	}
}
