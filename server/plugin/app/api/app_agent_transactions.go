package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var AppAgentTransactions = new(appAgentTransactions)

type appAgentTransactions struct{}

// CreateAppAgentTransactions 创建appAgentTransactions表
// @Tags AppAgentTransactions
// @Summary 创建appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppAgentTransactions true "创建appAgentTransactions表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /appAgentTransactions/createAppAgentTransactions [post]
func (a *appAgentTransactions) CreateAppAgentTransactions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.AppAgentTransactions
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAppAgentTransactions.CreateAppAgentTransactions(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAppAgentTransactions 删除appAgentTransactions表
// @Tags AppAgentTransactions
// @Summary 删除appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppAgentTransactions true "删除appAgentTransactions表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /appAgentTransactions/deleteAppAgentTransactions [delete]
func (a *appAgentTransactions) DeleteAppAgentTransactions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := serviceAppAgentTransactions.DeleteAppAgentTransactions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAppAgentTransactionsByIds 批量删除appAgentTransactions表
// @Tags AppAgentTransactions
// @Summary 批量删除appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /appAgentTransactions/deleteAppAgentTransactionsByIds [delete]
func (a *appAgentTransactions) DeleteAppAgentTransactionsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := serviceAppAgentTransactions.DeleteAppAgentTransactionsByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAppAgentTransactions 更新appAgentTransactions表
// @Tags AppAgentTransactions
// @Summary 更新appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppAgentTransactions true "更新appAgentTransactions表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /appAgentTransactions/updateAppAgentTransactions [put]
func (a *appAgentTransactions) UpdateAppAgentTransactions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.AppAgentTransactions
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAppAgentTransactions.UpdateAppAgentTransactions(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAppAgentTransactions 用id查询appAgentTransactions表
// @Tags AppAgentTransactions
// @Summary 用id查询appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询appAgentTransactions表"
// @Success 200 {object} response.Response{data=model.AppAgentTransactions,msg=string} "查询成功"
// @Router /appAgentTransactions/findAppAgentTransactions [get]
func (a *appAgentTransactions) FindAppAgentTransactions(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reappAgentTransactions, err := serviceAppAgentTransactions.GetAppAgentTransactions(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reappAgentTransactions, c)
}

// GetAppAgentTransactionsList 分页获取appAgentTransactions表列表
// @Tags AppAgentTransactions
// @Summary 分页获取appAgentTransactions表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AppAgentTransactionsSearch true "分页获取appAgentTransactions表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /appAgentTransactions/getAppAgentTransactionsList [get]
func (a *appAgentTransactions) GetAppAgentTransactionsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.AppAgentTransactionsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 自动注入商户ID (如果是商户管理员)
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid == nil && mid != nil {
		// 如果绑定了商户，强制查询该商户的数据
		merchantID := uint(*mid)
		pageInfo.MerchantId = &merchantID
	}
	list, total, err := serviceAppAgentTransactions.GetAppAgentTransactionsInfoList(ctx, pageInfo)
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

// GetMyAgentTransactionsList APP用户查询自己的分润明细
// @Tags AppAgentTransactions
// @Summary APP用户查询自己的分润明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /appAgentTransactions/getMyAgentTransactionsList [get]
func (a *appAgentTransactions) GetMyAgentTransactionsList(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo request.AppAgentTransactionsSearch
	_ = c.ShouldBindQuery(&pageInfo)

	// 获取当前登录用户ID
	userID := utils.GetUserID(c)
	uid := uint(userID)
	pageInfo.BeneficiaryId = &uid // 强制查询受益人为自己的记录

	list, total, err := serviceAppAgentTransactions.GetAppAgentTransactionsInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	
	// 补充商户信息 (list 是 []model.AppAgentTransactions 指针类型)
	// 需要预加载 Merchant 或者手动查询
	// 简单起见，这里遍历查询一下，或者让 Service 层支持 Preload
	// 为了不修改 Service 的通用逻辑，这里手动补全
	
	type MyTransactionResponse struct {
		model.AppAgentTransactions
		MerchantName string `json:"merchantName"`
	}
	
	var resultList []MyTransactionResponse
	for _, item := range list {
		resp := MyTransactionResponse{
			AppAgentTransactions: item,
		}
		if item.MerchantId != nil {
			var merchant model.Merchants
			if err := global.GVA_DB.WithContext(ctx).Select("merchant_name").First(&merchant, *item.MerchantId).Error; err == nil {
				if merchant.MerchantName != nil {
					resp.MerchantName = *merchant.MerchantName
				}
			}
		}
		resultList = append(resultList, resp)
	}

	response.OkWithDetailed(response.PageResult{
		List:     resultList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
