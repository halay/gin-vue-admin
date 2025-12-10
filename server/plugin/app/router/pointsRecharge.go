package router

import (
	"github.com/gin-gonic/gin"
)

var PointsRecharge = new(PR)

type PR struct{}

func (r *PR) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("pr")
		group.POST("createRechargeOrder", apiPointsRecharge.CreateRechargeOrder)
	}
	{
		group := public.Group("pr")
		group.POST("webhook", apiPointsRecharge.PayCallback)
	}
}
