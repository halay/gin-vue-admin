package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MerchantAnnouncement = new(MA)

type MA struct{}

// Init 初始化 商户公告信息 路由信息
func (r *MA) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("MA").Use(middleware.OperationRecord())
		group.POST("createMerchantAnnouncement", apiMerchantAnnouncement.CreateMerchantAnnouncement)             // 新建商户公告信息
		group.DELETE("deleteMerchantAnnouncement", apiMerchantAnnouncement.DeleteMerchantAnnouncement)           // 删除商户公告信息
		group.DELETE("deleteMerchantAnnouncementByIds", apiMerchantAnnouncement.DeleteMerchantAnnouncementByIds) // 批量删除商户公告信息
		group.PUT("updateMerchantAnnouncement", apiMerchantAnnouncement.UpdateMerchantAnnouncement)              // 更新商户公告信息
	}
	{
		group := private.Group("MA")
		group.GET("findMerchantAnnouncement", apiMerchantAnnouncement.FindMerchantAnnouncement)       // 根据ID获取商户公告信息
		group.GET("getMerchantAnnouncementList", apiMerchantAnnouncement.GetMerchantAnnouncementList) // 获取商户公告信息列表
	}
	{
		group := public.Group("MA")
		group.GET("getMerchantAnnouncementDataSource", apiMerchantAnnouncement.GetMerchantAnnouncementDataSource) // 获取商户公告信息数据源
		group.GET("getMerchantAnnouncementPublic", apiMerchantAnnouncement.GetMerchantAnnouncementPublic)         // 商户公告信息开放接口
	}
}
