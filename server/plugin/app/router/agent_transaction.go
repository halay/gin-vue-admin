package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var AgentTransactionDetail = new(atd)

type atd struct {}

// Init 初始化 代理交易明细 路由信息
func (r *atd) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("atd").Use(middleware.OperationRecord())
		group.POST("createAgentTransactionDetail", apiAgentTransactionDetail.CreateAgentTransactionDetail)   // 新建代理交易明细
		group.DELETE("deleteAgentTransactionDetail", apiAgentTransactionDetail.DeleteAgentTransactionDetail) // 删除代理交易明细
		group.DELETE("deleteAgentTransactionDetailByIds", apiAgentTransactionDetail.DeleteAgentTransactionDetailByIds) // 批量删除代理交易明细
		group.PUT("updateAgentTransactionDetail", apiAgentTransactionDetail.UpdateAgentTransactionDetail)    // 更新代理交易明细
	}
	{
	    group := private.Group("atd")
		group.GET("findAgentTransactionDetail", apiAgentTransactionDetail.FindAgentTransactionDetail)        // 根据ID获取代理交易明细
		group.GET("getAgentTransactionDetailList", apiAgentTransactionDetail.GetAgentTransactionDetailList)  // 获取代理交易明细列表
	}
	{
	    group := public.Group("atd")
	    group.GET("getAgentTransactionDetailPublic", apiAgentTransactionDetail.GetAgentTransactionDetailPublic)  // 代理交易明细开放接口
	}
}
