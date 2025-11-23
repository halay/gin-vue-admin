package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MerchantCategory = new(MCAT)

type MCAT struct{}

// Init 初始化 商户分类 路由信息
func (r *MCAT) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("MCAT").Use(middleware.OperationRecord())
		group.POST("createMerchantCategory", apiMerchantCategory.CreateMerchantCategory)             // 新建商户分类
		group.DELETE("deleteMerchantCategory", apiMerchantCategory.DeleteMerchantCategory)           // 删除商户分类
		group.DELETE("deleteMerchantCategoryByIds", apiMerchantCategory.DeleteMerchantCategoryByIds) // 批量删除商户分类
		group.PUT("updateMerchantCategory", apiMerchantCategory.UpdateMerchantCategory)              // 更新商户分类
	}
	{
		group := private.Group("MCAT")
		group.GET("findMerchantCategory", apiMerchantCategory.FindMerchantCategory)       // 根据ID获取商户分类
		group.GET("getMerchantCategoryList", apiMerchantCategory.GetMerchantCategoryList) // 获取商户分类列表
	}
	{
		group := public.Group("MCAT")
		group.GET("getMerchantCategoryPublic", apiMerchantCategory.GetMerchantCategoryPublic) // 商户分类开放接口
	}
}
