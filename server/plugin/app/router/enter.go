package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/api"

var (
	Router                  = new(router)
	apiAppUsers             = api.Api.AppUsers
	apiMerchants            = api.Api.Merchants
	apiYApi                 = api.Api.YApi
	apiBanner               = api.Api.Banner
	apiMerchantAnnouncement = api.Api.MerchantAnnouncement
	apiMerchantAdmin        = api.Api.MerchantAdmin
	apiProductCategory      = api.Api.ProductCategory
	apiProduct              = api.Api.Product
	apiProductSku           = api.Api.ProductSku
	apiProductSpec          = api.Api.ProductSpec
	apiProductSpecOption    = api.Api.ProductSpecOption
	apiOrder                = api.Api.Order
	apiOrderItem            = api.Api.OrderItem
	apiUserPointsAccount    = api.Api.UserPointsAccount
	apiUserPointsLog        = api.Api.UserPointsLog
)

type router struct {
	AppUsers             appUsers
	Merchants            mc
	YApi                 yApi
	Banner               banner
	MerchantAnnouncement MA
	MerchantAdmin        MADM
	ProductCategory      PC
	Product              P
	ProductSku           SKU
	ProductSpec          PS
	ProductSpecOption    PSO
	Order                ORD
	OrderItem            ORDI
	UserPointsAccount    UPA
	UserPointsLog        UPL
}
