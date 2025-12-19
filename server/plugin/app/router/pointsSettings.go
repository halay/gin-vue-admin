package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var PointsSettings = new(PTS)

type PTS struct {}

// Init 初始化 平台积分基础设置与奖励规则 路由信息
func (r *PTS) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("PTS").Use(middleware.OperationRecord())
		group.POST("createPointsSettings", apiPointsSettings.CreatePointsSettings)   // 新建平台积分基础设置与奖励规则
		group.DELETE("deletePointsSettings", apiPointsSettings.DeletePointsSettings) // 删除平台积分基础设置与奖励规则
		group.DELETE("deletePointsSettingsByIds", apiPointsSettings.DeletePointsSettingsByIds) // 批量删除平台积分基础设置与奖励规则
		group.PUT("updatePointsSettings", apiPointsSettings.UpdatePointsSettings)    // 更新平台积分基础设置与奖励规则
	}
	{
	    group := private.Group("PTS")
		group.GET("findPointsSettings", apiPointsSettings.FindPointsSettings)        // 根据ID获取平台积分基础设置与奖励规则
		group.GET("getPointsSettingsList", apiPointsSettings.GetPointsSettingsList)  // 获取平台积分基础设置与奖励规则列表
	}
	{
	    group := public.Group("PTS")
	    group.GET("getPointsSettingsPublic", apiPointsSettings.GetPointsSettingsPublic)  // 平台积分基础设置与奖励规则开放接口
	}
}
