package router

import (
    "github.com/flipped-aurora/gin-vue-admin/server/middleware"
    "github.com/gin-gonic/gin"
)

var Consultation = new(CN)

type CN struct{}

func (r *CN) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
    {
        group := private.Group("CN").Use(middleware.OperationRecord())
        group.POST("createConsultation", apiConsultation.CreateConsultation)
        group.DELETE("deleteConsultation", apiConsultation.DeleteConsultation)
        group.DELETE("deleteConsultationByIds", apiConsultation.DeleteConsultationByIds)
        group.PUT("updateConsultation", apiConsultation.UpdateConsultation)
    }
    {
        group := private.Group("CN")
        group.GET("findConsultation", apiConsultation.FindConsultation)
        group.GET("getConsultationList", apiConsultation.GetConsultationList)
    }
    {
        group := public.Group("CN")
        group.GET("getConsultationPublic", apiConsultation.GetConsultationPublic)
        group.GET("findConsultationPublic", apiConsultation.FindConsultationPublic)
    }
}