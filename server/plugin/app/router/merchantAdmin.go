package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MerchantAdmin = new(MADM)

type MADM struct {}

// Init 初始化 后台用户与商户关联映射 路由信息
func (r *MADM) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("MADM").Use(middleware.OperationRecord())
		group.POST("createMerchantAdmin", apiMerchantAdmin.CreateMerchantAdmin)   // 新建后台用户与商户关联映射
		group.DELETE("deleteMerchantAdmin", apiMerchantAdmin.DeleteMerchantAdmin) // 删除后台用户与商户关联映射
		group.DELETE("deleteMerchantAdminByIds", apiMerchantAdmin.DeleteMerchantAdminByIds) // 批量删除后台用户与商户关联映射
		group.PUT("updateMerchantAdmin", apiMerchantAdmin.UpdateMerchantAdmin)    // 更新后台用户与商户关联映射
	}
	{
	    group := private.Group("MADM")
		group.GET("findMerchantAdmin", apiMerchantAdmin.FindMerchantAdmin)        // 根据ID获取后台用户与商户关联映射
		group.GET("getMerchantAdminList", apiMerchantAdmin.GetMerchantAdminList)  // 获取后台用户与商户关联映射列表
	}
	{
	    group := public.Group("MADM")
	    group.GET("getMerchantAdminDataSource", apiMerchantAdmin.GetMerchantAdminDataSource)  // 获取后台用户与商户关联映射数据源
	    group.GET("getMerchantAdminPublic", apiMerchantAdmin.GetMerchantAdminPublic)  // 后台用户与商户关联映射开放接口
	}
}
