package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var AppDealer = new(ADL)

type ADL struct {}

// Init 初始化 经销商管理 路由信息
func (r *ADL) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("ADL").Use(middleware.OperationRecord())
		group.POST("createAppDealer", apiAppDealer.CreateAppDealer)   // 新建经销商管理
		group.DELETE("deleteAppDealer", apiAppDealer.DeleteAppDealer) // 删除经销商管理
		group.DELETE("deleteAppDealerByIds", apiAppDealer.DeleteAppDealerByIds) // 批量删除经销商管理
		group.PUT("updateAppDealer", apiAppDealer.UpdateAppDealer)    // 更新经销商管理
	}
	{
	    group := private.Group("ADL")
		group.GET("findAppDealer", apiAppDealer.FindAppDealer)        // 根据ID获取经销商管理
		group.GET("getAppDealerList", apiAppDealer.GetAppDealerList)  // 获取经销商管理列表
	}
	{
	    group := public.Group("ADL")
	    group.GET("getAppDealerDataSource", apiAppDealer.GetAppDealerDataSource)  // 获取经销商管理数据源
	    group.GET("getAppDealerPublic", apiAppDealer.GetAppDealerPublic)  // 经销商管理开放接口
	}
}
