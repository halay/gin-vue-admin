package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/router"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	public.Use()
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	router.Router.AppUsers.Init(public, private)
	router.Router.Merchants.Init(public, private)
	router.Router.YApi.Init(public, private)
	router.Router.Banner.Init(public, private)
	router.Router.MerchantAnnouncement.Init(public, private)
	router.Router.MerchantAdmin.Init(public, private)
	router.Router.ProductCategory.Init(public, private)
	router.Router.Product.Init(public, private)
	router.Router.ProductSku.Init(public, private)
	router.Router.ProductSpec.Init(public, private)
	router.Router.ProductSpecOption.Init(public, private)
	router.Router.Order.Init(public, private)
	router.Router.OrderItem.Init(public, private)
	router.Router.UserPointsAccount.Init(public, private)
	router.Router.UserPointsLog.Init(public, private)
	router.Router.Search.Init(public, private)
	router.Router.Consultation.Init(public, private)
	router.Router.PointsCfg.Init(public, private)
	router.Router.PointsRecharge.Init(public, private)
	router.Router.AppRelease.Init(public, private)
	router.Router.MerchantCategory.Init(public, private)
	router.Router.MembershipLevel.Init(public, private)
	router.Router.Node.Init(public, private)
	router.Router.PointsSettings.Init(public, private)
	router.Router.MerchantPointsSettings.Init(public, private)
	router.Router.UserAddress.Init(public, private)
	router.Router.AgentLevel.Init(public, private)
	router.Router.DownlinePurchaseRecord.InitDownlinePurchaseRecordRouter(private)
}
