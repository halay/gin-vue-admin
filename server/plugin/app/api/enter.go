package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/service"

var (
	Api                           = new(api)
	serviceAppUsers               = service.Service.AppUsers
	serviceMerchants              = service.Service.Merchants
	serviceBanner                 = service.Service.Banner
	serviceMerchantAnnouncement   = service.Service.MerchantAnnouncement
	serviceMerchantAdmin          = service.Service.MerchantAdmin
	serviceProductCategory        = service.Service.ProductCategory
	serviceProduct                = service.Service.Product
	serviceProductSku             = service.Service.ProductSku
	serviceProductSpec            = service.Service.ProductSpec
	serviceProductSpecOption      = service.Service.ProductSpecOption
	serviceOrder                  = service.Service.Order
	serviceOrderItem              = service.Service.OrderItem
	serviceUserPointsAccount      = service.Service.UserPointsAccount
	serviceUserPointsLog          = service.Service.UserPointsLog
	serviceSearch                 = service.Service.Search
	serviceConsultation           = service.Service.Consultation
	serviceAppRelease             = service.Service.AppRelease
	serviceMerchantCategory       = service.Service.MerchantCategory
	serviceMembershipLevel        = service.Service.MembershipLevel
	serviceNode                   = service.Service.Node
	servicePointsCfg              = service.Service.PointsCfg
	servicePointsRecharge         = service.Service.PointsRecharge
	servicePointsSettings         = service.Service.PointsSettings
	serviceMerchantPointsSettings = service.Service.MerchantPointsSettings
	serviceUserAddress            = service.Service.UserAddress
)

type api struct {
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
	Search                 SEARCH
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
}
