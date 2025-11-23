package router

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
)

var Merchants = new(mc)

type mc struct{}

// Init 初始化 商家信息 路由信息
func (r *mc) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("mc").Use(middleware.OperationRecord())
		group.POST("createMerchants", apiMerchants.CreateMerchants)             // 新建商家信息
		group.DELETE("deleteMerchants", apiMerchants.DeleteMerchants)           // 删除商家信息
		group.DELETE("deleteMerchantsByIds", apiMerchants.DeleteMerchantsByIds) // 批量删除商家信息
		group.PUT("updateMerchants", apiMerchants.UpdateMerchants)              // 更新商家信息
	}
	{
		group := private.Group("mc")
		group.GET("findMerchants", apiMerchants.FindMerchants)       // 根据ID获取商家信息
		group.GET("getMerchantsList", apiMerchants.GetMerchantsList) // 获取商家信息列表
	}
	{
		group := public.Group("mc")
		group.GET("getMerchantsPublic", apiMerchants.GetMerchantsPublic) // 商家信息开放接口
		group.GET("getMerchantsData", apiMerchants.FindMerchants)        // 根据ID获取商家信息
		group.GET("getMerchantsListData", apiMerchants.GetMerchantsList) // 获取商家信息列表
		group.GET("getRecommendedMerchants", apiMerchants.GetRecommendedMerchants) // 推荐商家列表
		group.GET("getMerchantCategoryList", apiMerchants.GetMerchantCategoryList) // 商家分类列表
		group.GET("getMerchantsByCategory", apiMerchants.GetMerchantsByCategory) // 按分类获取商家
	}
}
