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
)

type router struct {
	AppUsers             appUsers
	Merchants            mc
	YApi                 yApi
	Banner               banner
	MerchantAnnouncement MA
	MerchantAdmin        MADM
}
