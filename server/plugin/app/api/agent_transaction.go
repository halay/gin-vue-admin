package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AgentTransactionApi struct{}

// GetAgentTransactionList 分页获取代理分润记录列表
// @Tags AgentTransaction
// @Summary 分页获取代理分润记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AgentTransactionSearch true "分页获取代理分润记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /agentTransaction/getAgentTransactionList [get]
func (a *AgentTransactionApi) GetAgentTransactionList(c *gin.Context) {
	var pageInfo request.AgentTransactionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 自动注入商户ID (如果是商户管理员)
	userID := utils.GetUserID(c)
	ctx := c.Request.Context()
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid == nil && mid != nil {
		// 如果绑定了商户，强制查询该商户的数据
		merchantID := uint(*mid)
		pageInfo.MerchantId = &merchantID
	}

	list, total, err := serviceAgentTransaction.GetAgentTransactionList(pageInfo)
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
