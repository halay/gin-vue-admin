package api

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/service"

var (
	Api             = new(api)
	serviceAppUsers = service.Service.AppUsers
)

type api struct{ AppUsers appUsers }
