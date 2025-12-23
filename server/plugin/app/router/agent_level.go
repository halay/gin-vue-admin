package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var AgentLevel = new(AL)

type AL struct {}

// Init 初始化 营销方案设置 路由信息
func (r *AL) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("AL").Use(middleware.OperationRecord())
		group.POST("createAgentLevel", apiAgentLevel.CreateAgentLevel)   // 新建营销方案设置
		group.DELETE("deleteAgentLevel", apiAgentLevel.DeleteAgentLevel) // 删除营销方案设置
		group.DELETE("deleteAgentLevelByIds", apiAgentLevel.DeleteAgentLevelByIds) // 批量删除营销方案设置
		group.PUT("updateAgentLevel", apiAgentLevel.UpdateAgentLevel)    // 更新营销方案设置
	}
	{
	    group := private.Group("AL")
		group.GET("findAgentLevel", apiAgentLevel.FindAgentLevel)        // 根据ID获取营销方案设置
		group.GET("getAgentLevelList", apiAgentLevel.GetAgentLevelList)  // 获取营销方案设置列表
	}
	{
	    group := public.Group("AL")
	    group.GET("getAgentLevelDataSource", apiAgentLevel.GetAgentLevelDataSource)  // 获取营销方案设置数据源
	    group.GET("getAgentLevelPublic", apiAgentLevel.GetAgentLevelPublic)  // 营销方案设置开放接口
	}
}
