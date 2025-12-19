package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MerchantPointsSettings = new(MPTS)

type MPTS struct {}

// Init 初始化 商户积分营销配置 路由信息
func (r *MPTS) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("MPTS").Use(middleware.OperationRecord())
		group.POST("createMerchantPointsSettings", apiMerchantPointsSettings.CreateMerchantPointsSettings)   // 新建商户积分营销配置
		group.DELETE("deleteMerchantPointsSettings", apiMerchantPointsSettings.DeleteMerchantPointsSettings) // 删除商户积分营销配置
		group.DELETE("deleteMerchantPointsSettingsByIds", apiMerchantPointsSettings.DeleteMerchantPointsSettingsByIds) // 批量删除商户积分营销配置
		group.PUT("updateMerchantPointsSettings", apiMerchantPointsSettings.UpdateMerchantPointsSettings)    // 更新商户积分营销配置
	}
	{
	    group := private.Group("MPTS")
		group.GET("findMerchantPointsSettings", apiMerchantPointsSettings.FindMerchantPointsSettings)        // 根据ID获取商户积分营销配置
		group.GET("getMerchantPointsSettingsList", apiMerchantPointsSettings.GetMerchantPointsSettingsList)  // 获取商户积分营销配置列表
	}
	{
	    group := public.Group("MPTS")
	    group.GET("getMerchantPointsSettingsPublic", apiMerchantPointsSettings.GetMerchantPointsSettingsPublic)  // 商户积分营销配置开放接口
	}
}
