package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/api"
	"github.com/gin-gonic/gin"
)

type DownlinePurchaseRecordRouter struct{}

func (s *DownlinePurchaseRecordRouter) InitDownlinePurchaseRecordRouter(Router *gin.RouterGroup) {
	dprRouter := Router.Group("dpr")
	dprApi := api.Api.DownlinePurchaseRecord
	{
		dprRouter.GET("getDownlinePurchaseRecordList", dprApi.GetDownlinePurchaseRecordList)     // 后台/商户查看
		dprRouter.GET("getMyDownlinePurchaseRecordList", dprApi.GetMyDownlinePurchaseRecordList) // App用户查看
	}
}
