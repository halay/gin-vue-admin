package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/api"

var (
	Router       = new(router)
	apiAppUsers  = api.Api.AppUsers
	apiMerchants = api.Api.Merchants
	apiYApi      = api.Api.YApi
)

type router struct {
	AppUsers  appUsers
	Merchants mc
	YApi      yApi
}
