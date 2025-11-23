package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var AppRelease = new(AR)

type AR struct {}

// Init 初始化 App客户端版本升级 路由信息
func (r *AR) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("AR").Use(middleware.OperationRecord())
		group.POST("createAppRelease", apiAppRelease.CreateAppRelease)   // 新建App客户端版本升级
		group.DELETE("deleteAppRelease", apiAppRelease.DeleteAppRelease) // 删除App客户端版本升级
		group.DELETE("deleteAppReleaseByIds", apiAppRelease.DeleteAppReleaseByIds) // 批量删除App客户端版本升级
		group.PUT("updateAppRelease", apiAppRelease.UpdateAppRelease)    // 更新App客户端版本升级
	}
	{
	    group := private.Group("AR")
		group.GET("findAppRelease", apiAppRelease.FindAppRelease)        // 根据ID获取App客户端版本升级
		group.GET("getAppReleaseList", apiAppRelease.GetAppReleaseList)  // 获取App客户端版本升级列表
	}
	{
	    group := public.Group("AR")
	    group.GET("getAppReleasePublic", apiAppRelease.GetAppReleasePublic)  // App客户端版本升级开放接口
	}
}
