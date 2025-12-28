package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Settlement = new(settlement)

type settlement struct {}

// Init 初始化 结算管理 路由信息
func (r *settlement) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("settlement").Use(middleware.OperationRecord())
		group.POST("createSettlement", apiSettlement.CreateSettlement)   // 新建结算管理
		group.DELETE("deleteSettlement", apiSettlement.DeleteSettlement) // 删除结算管理
		group.DELETE("deleteSettlementByIds", apiSettlement.DeleteSettlementByIds) // 批量删除结算管理
		group.PUT("updateSettlement", apiSettlement.UpdateSettlement)    // 更新结算管理
	}
	{
	    group := private.Group("settlement")
		group.GET("findSettlement", apiSettlement.FindSettlement)        // 根据ID获取结算管理
		group.GET("getSettlementList", apiSettlement.GetSettlementList)  // 获取结算管理列表
	}
	{
	    group := public.Group("settlement")
	    group.GET("getSettlementPublic", apiSettlement.GetSettlementPublic)  // 结算管理开放接口
	}
}
