package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Node = new(NODE)

type NODE struct {}

// Init 初始化 节点策略 路由信息
func (r *NODE) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("NODE").Use(middleware.OperationRecord())
		group.POST("createNode", apiNode.CreateNode)   // 新建节点策略
		group.DELETE("deleteNode", apiNode.DeleteNode) // 删除节点策略
		group.DELETE("deleteNodeByIds", apiNode.DeleteNodeByIds) // 批量删除节点策略
		group.PUT("updateNode", apiNode.UpdateNode)    // 更新节点策略
	}
	{
	    group := private.Group("NODE")
		group.GET("findNode", apiNode.FindNode)        // 根据ID获取节点策略
		group.GET("getNodeList", apiNode.GetNodeList)  // 获取节点策略列表
	}
	{
	    group := public.Group("NODE")
	    group.GET("getNodePublic", apiNode.GetNodePublic)  // 节点策略开放接口
	}
}
