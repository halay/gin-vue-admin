package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var OrderItem = new(ORDI)

type ORDI struct{}

// Init 初始化 订单明细 路由信息
func (r *ORDI) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("ORDI").Use(middleware.OperationRecord())
		group.POST("createOrderItem", apiOrderItem.CreateOrderItem)             // 新建订单明细
		group.DELETE("deleteOrderItem", apiOrderItem.DeleteOrderItem)           // 删除订单明细
		group.DELETE("deleteOrderItemByIds", apiOrderItem.DeleteOrderItemByIds) // 批量删除订单明细
		group.PUT("updateOrderItem", apiOrderItem.UpdateOrderItem)              // 更新订单明细
	}
	{
		group := private.Group("ORDI")
		group.GET("findOrderItem", apiOrderItem.FindOrderItem)       // 根据ID获取订单明细
		group.GET("getOrderItemList", apiOrderItem.GetOrderItemList) // 获取订单明细列表
	}
	{
		group := public.Group("ORDI")
		group.GET("getOrderItemDataSource", apiOrderItem.GetOrderItemDataSource) // 获取订单明细数据源
		group.GET("getOrderItemPublic", apiOrderItem.GetOrderItemPublic)         // 订单明细开放接口
	}
}
