package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var ShareholderProfit = new(spr)

type spr struct {}

// Init 初始化 股东分润 路由信息
func (r *spr) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("spr").Use(middleware.OperationRecord())
		group.POST("createShareholderProfit", apiShareholderProfit.CreateShareholderProfit)   // 新建股东分润
		group.DELETE("deleteShareholderProfit", apiShareholderProfit.DeleteShareholderProfit) // 删除股东分润
		group.DELETE("deleteShareholderProfitByIds", apiShareholderProfit.DeleteShareholderProfitByIds) // 批量删除股东分润
		group.PUT("updateShareholderProfit", apiShareholderProfit.UpdateShareholderProfit)    // 更新股东分润
	}
	{
	    group := private.Group("spr")
		group.GET("findShareholderProfit", apiShareholderProfit.FindShareholderProfit)        // 根据ID获取股东分润
		group.GET("getShareholderProfitList", apiShareholderProfit.GetShareholderProfitList)  // 获取股东分润列表
	}
	{
	    group := public.Group("spr")
	    group.GET("getShareholderProfitPublic", apiShareholderProfit.GetShareholderProfitPublic)  // 股东分润开放接口
	}
}
