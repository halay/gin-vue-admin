package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type DownlinePurchaseRecordApi struct{}

// GetDownlinePurchaseRecordList 分页获取下级购买记录列表 (后台/商户)
// @Tags DownlinePurchaseRecord
// @Summary 分页获取下级购买记录列表 (后台/商户)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.DownlinePurchaseRecordSearch true "分页获取下级购买记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dpr/getDownlinePurchaseRecordList [get]
func (a *DownlinePurchaseRecordApi) GetDownlinePurchaseRecordList(c *gin.Context) {
	var pageInfo request.DownlinePurchaseRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 自动注入商户ID (如果是商户管理员)
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("未登录或无权限", c)
		return
	}
	ctx := c.Request.Context()
	if userID != 1 {
		mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
		if errMid == nil && mid != nil {
			// 如果绑定了商户，强制查询该商户的数据
			merchantID := uint(*mid)
			pageInfo.MerchantId = &merchantID
		}
	}
	list, total, err := serviceDownlinePurchaseRecord.GetDownlinePurchaseRecordList(ctx, pageInfo)
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

// GetMyDownlinePurchaseRecordList 分页获取我的下级购买记录列表 (APP端)
// @Tags DownlinePurchaseRecord
// @Summary 分页获取我的下级购买记录列表 (APP端)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.DownlinePurchaseRecordSearch true "分页获取我的下级购买记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dpr/getMyDownlinePurchaseRecordList [get]
func (a *DownlinePurchaseRecordApi) GetMyDownlinePurchaseRecordList(c *gin.Context) {
	var pageInfo request.DownlinePurchaseRecordSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)
	ctx := c.Request.Context()

	list, total, err := serviceDownlinePurchaseRecord.GetMyDownlinePurchaseRecordList(ctx, userID, pageInfo)
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
