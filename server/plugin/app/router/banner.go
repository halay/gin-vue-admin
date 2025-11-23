package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Banner = new(banner)

type banner struct{}

// Init 初始化 Banner图片 路由信息
func (r *banner) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("banner").Use(middleware.OperationRecord())
		group.POST("createBanner", apiBanner.CreateBanner)             // 新建Banner图片
		group.DELETE("deleteBanner", apiBanner.DeleteBanner)           // 删除Banner图片
		group.DELETE("deleteBannerByIds", apiBanner.DeleteBannerByIds) // 批量删除Banner图片
		group.PUT("updateBanner", apiBanner.UpdateBanner)              // 更新Banner图片
	}
	{
		group := private.Group("banner")
		group.GET("findBanner", apiBanner.FindBanner)       // 根据ID获取Banner图片
		group.GET("getBannerList", apiBanner.GetBannerList) // 获取Banner图片列表
	}
	{
		group := public.Group("banner")
		group.GET("getBannerPublic", apiBanner.GetBannerPublic) // Banner图片开放接口
	}
}
