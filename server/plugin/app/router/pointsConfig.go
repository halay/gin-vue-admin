package router

import (
    "github.com/flipped-aurora/gin-vue-admin/server/middleware"
    "github.com/gin-gonic/gin"
)

var PointsCfg = new(PCFG)

type PCFG struct{}

func (r *PCFG) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
    {
        group := private.Group("pcfg").Use(middleware.OperationRecord())
        group.POST("createPointsConfig", apiPointsCfg.CreatePointsConfig)
        group.DELETE("deletePointsConfig", apiPointsCfg.DeletePointsConfig)
        group.DELETE("deletePointsConfigByIds", apiPointsCfg.DeletePointsConfigByIds)
        group.PUT("updatePointsConfig", apiPointsCfg.UpdatePointsConfig)
    }
    {
        group := private.Group("pcfg")
        group.GET("findPointsConfig", apiPointsCfg.FindPointsConfig)
        group.GET("getPointsConfigList", apiPointsCfg.GetPointsConfigList)
    }
    {
        group := public.Group("pcfg")
        group.GET("getPointsConfigPublic", apiPointsCfg.GetPointsConfigPublic)
    }
}

