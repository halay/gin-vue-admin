package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Consultation = new(CN)

type CN struct{}

func (a *CN) CreateConsultation(c *gin.Context) {
	ctx := c.Request.Context()
	var info model.AppConsultation
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serviceConsultation.CreateConsultation(ctx, &info); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (a *CN) DeleteConsultation(c *gin.Context) {
	ctx := c.Request.Context()
	ID := c.Query("ID")
	if err := serviceConsultation.DeleteConsultation(ctx, ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (a *CN) DeleteConsultationByIds(c *gin.Context) {
	ctx := c.Request.Context()
	IDs := c.QueryArray("IDs[]")
	if err := serviceConsultation.DeleteConsultationByIds(ctx, IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

func (a *CN) UpdateConsultation(c *gin.Context) {
	ctx := c.Request.Context()
	var info model.AppConsultation
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serviceConsultation.UpdateConsultation(ctx, info); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *CN) FindConsultation(c *gin.Context) {
	ctx := c.Request.Context()
	ID := c.Query("ID")
	re, err := serviceConsultation.GetConsultation(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(re, c)
}

func (a *CN) GetConsultationList(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo request.ConsultationSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceConsultation.GetConsultationInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}

func (a *CN) GetConsultationPublic(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo request.ConsultationSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceConsultation.GetConsultationPublicList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}

func (a *CN) FindConsultationPublic(c *gin.Context) {
	ctx := c.Request.Context()
	ID := c.Query("ID")
	re, err := serviceConsultation.GetConsultationPublicDetail(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(re, c)
}
