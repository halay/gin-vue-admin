package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/service"

var (
	Api                         = new(api)
	serviceAppUsers             = service.Service.AppUsers
	serviceMerchants            = service.Service.Merchants
	serviceBanner               = service.Service.Banner
	serviceMerchantAnnouncement = service.Service.MerchantAnnouncement
	serviceMerchantAdmin        = service.Service.MerchantAdmin
)

type api struct {
	AppUsers             appUsers
	Merchants            mc
	YApi                 yApi
	Banner               banner
	MerchantAnnouncement MA
	MerchantAdmin        MADM
}
