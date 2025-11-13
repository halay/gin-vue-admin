package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Order = new(ORD)

type ORD struct {}

// Init 初始化 订单 路由信息
func (r *ORD) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
    {
        group := private.Group("ORD").Use(middleware.OperationRecord())
        group.POST("createOrder", apiOrder.CreateOrder)   // 新建订单
        group.DELETE("deleteOrder", apiOrder.DeleteOrder) // 删除订单
        group.DELETE("deleteOrderByIds", apiOrder.DeleteOrderByIds) // 批量删除订单
        group.PUT("updateOrder", apiOrder.UpdateOrder)    // 更新订单
    }
    {
        group := private.Group("ORD")
        group.GET("findOrder", apiOrder.FindOrder)        // 根据ID获取订单
        group.GET("getOrderList", apiOrder.GetOrderList)  // 获取订单列表
    }
    {
        group := public.Group("ORD")
        group.GET("getOrderDataSource", apiOrder.GetOrderDataSource)  // 获取订单数据源
        group.GET("getOrderPublic", apiOrder.GetOrderPublic)  // 订单开放接口
        group.POST("createOrderByPoints", apiOrder.CreateOrderByPoints) // C端创建积分订单
        group.POST("payOrderByPoints", apiOrder.PayOrderByPoints) // C端积分支付
    }
}
