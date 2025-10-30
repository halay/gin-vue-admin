package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/api"

var (
	Router      = new(router)
	apiAppUsers = api.Api.AppUsers
)

type router struct{ AppUsers appUsers }
