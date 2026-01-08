package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var ShareholderProfit = new(spr)

type spr struct{}

// CreateShareholderProfit 创建股东分润
// @Tags ShareholderProfit
// @Summary 创建股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ShareholderProfit true "创建股东分润"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /spr/createShareholderProfit [post]
func (a *spr) CreateShareholderProfit(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.ShareholderProfit
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建", c)
		return
	}
	err = serviceShareholderProfit.CreateShareholderProfit(ctx, &info, *mid)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteShareholderProfit 删除股东分润
// @Tags ShareholderProfit
// @Summary 删除股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ShareholderProfit true "删除股东分润"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /spr/deleteShareholderProfit [delete]
func (a *spr) DeleteShareholderProfit(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法删除", c)
		return
	}
	err := serviceShareholderProfit.DeleteShareholderProfit(ctx, ID, *mid)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteShareholderProfitByIds 批量删除股东分润
// @Tags ShareholderProfit
// @Summary 批量删除股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /spr/deleteShareholderProfitByIds [delete]
func (a *spr) DeleteShareholderProfitByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法批量删除", c)
		return
	}
	err := serviceShareholderProfit.DeleteShareholderProfitByIds(ctx, IDs, *mid)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateShareholderProfit 更新股东分润
// @Tags ShareholderProfit
// @Summary 更新股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ShareholderProfit true "更新股东分润"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /spr/updateShareholderProfit [put]
func (a *spr) UpdateShareholderProfit(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.ShareholderProfit
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法更新", c)
		return
	}
	err = serviceShareholderProfit.UpdateShareholderProfit(ctx, info, *mid)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindShareholderProfit 用id查询股东分润
// @Tags ShareholderProfit
// @Summary 用id查询股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询股东分润"
// @Success 200 {object} response.Response{data=model.ShareholderProfit,msg=string} "查询成功"
// @Router /spr/findShareholderProfit [get]
func (a *spr) FindShareholderProfit(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法查询", c)
		return
	}
	respr, err := serviceShareholderProfit.GetShareholderProfit(ctx, ID, *mid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(respr, c)
}

// GetShareholderProfitList 分页获取股东分润列表
// @Tags ShareholderProfit
// @Summary 分页获取股东分润列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.ShareholderProfitSearch true "分页获取股东分润列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /spr/getShareholderProfitList [get]
func (a *spr) GetShareholderProfitList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.ShareholderProfitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法获取订单列表", c)
		return
	}
	list, total, err := serviceShareholderProfit.GetShareholderProfitInfoList(ctx, pageInfo, *mid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetShareholderProfitPublic 不需要鉴权的股东分润接口
// @Tags ShareholderProfit
// @Summary 不需要鉴权的股东分润接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /spr/getShareholderProfitPublic [get]
func (a *spr) GetShareholderProfitPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceShareholderProfit.GetShareholderProfitPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的股东分润接口信息"}, "获取成功", c)
}
