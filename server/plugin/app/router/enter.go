package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/api"

var (
	Router                    = new(router)
	apiAppUsers               = api.Api.AppUsers
	apiMerchants              = api.Api.Merchants
	apiYApi                   = api.Api.YApi
	apiBanner                 = api.Api.Banner
	apiMerchantAnnouncement   = api.Api.MerchantAnnouncement
	apiMerchantAdmin          = api.Api.MerchantAdmin
	apiProductCategory        = api.Api.ProductCategory
	apiProduct                = api.Api.Product
	apiProductSku             = api.Api.ProductSku
	apiProductSpec            = api.Api.ProductSpec
	apiProductSpecOption      = api.Api.ProductSpecOption
	apiOrder                  = api.Api.Order
	apiOrderItem              = api.Api.OrderItem
	apiUserPointsAccount      = api.Api.UserPointsAccount
	apiUserPointsLog          = api.Api.UserPointsLog
	apiSearch                 = api.Api.Search
	apiConsultation           = api.Api.Consultation
	apiPointsCfg              = api.Api.PointsCfg
	apiPointsRecharge         = api.Api.PointsRecharge
	apiAppRelease             = api.Api.AppRelease
	apiMerchantCategory       = api.Api.MerchantCategory
	apiMembershipLevel        = api.Api.MembershipLevel
	apiNode                   = api.Api.Node
	apiPointsSettings         = api.Api.PointsSettings
	apiMerchantPointsSettings = api.Api.MerchantPointsSettings
	apiUserAddress            = api.Api.UserAddress
	apiAgentLevel             = api.Api.AgentLevel
	apiAgentTransactionDetail = api.Api.AgentTransactionDetail
	apiSettlement             = api.Api.Settlement
)

type router struct {
	AppUsers               appUsers
	Merchants              mc
	YApi                   yApi
	Banner                 banner
	MerchantAnnouncement   MA
	MerchantAdmin          MADM
	ProductCategory        PC
	Product                P
	ProductSku             SKU
	ProductSpec            PS
	ProductSpecOption      PSO
	Order                  ORD
	OrderItem              ORDI
	UserPointsAccount      UPA
	UserPointsLog          UPL
	Search                 S
	Consultation           CN
	AppRelease             AR
	MerchantCategory       MCAT
	MembershipLevel        ML
	Node                   NODE
	PointsCfg              PCFG
	PointsRecharge         PR
	PointsSettings         PTS
	MerchantPointsSettings MPTS
	UserAddress            UA
	AgentLevel             AL
	DownlinePurchaseRecord DownlinePurchaseRecordRouter
	AgentTransactionDetail atd
	Settlement             settlement
}
