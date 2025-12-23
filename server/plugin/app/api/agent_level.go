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

var AgentLevel = new(AL)

type AL struct{}

// CreateAgentLevel 创建营销方案设置
// @Tags AgentLevel
// @Summary 创建营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentLevel true "创建营销方案设置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /AL/createAgentLevel [post]
func (a *AL) CreateAgentLevel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.AgentLevel
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建营销方案", c)
		return
	}
	info.MerchantId = int64Ptr(int64(*mid))
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAgentLevel.CreateAgentLevel(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAgentLevel 删除营销方案设置
// @Tags AgentLevel
// @Summary 删除营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentLevel true "删除营销方案设置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /AL/deleteAgentLevel [delete]
func (a *AL) DeleteAgentLevel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建营销方案", c)
		return
	}
	err := serviceAgentLevel.DeleteAgentLevel(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAgentLevelByIds 批量删除营销方案设置
// @Tags AgentLevel
// @Summary 批量删除营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /AL/deleteAgentLevelByIds [delete]
func (a *AL) DeleteAgentLevelByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建营销方案", c)
		return
	}
	err := serviceAgentLevel.DeleteAgentLevelByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAgentLevel 更新营销方案设置
// @Tags AgentLevel
// @Summary 更新营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentLevel true "更新营销方案设置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /AL/updateAgentLevel [put]
func (a *AL) UpdateAgentLevel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.AgentLevel
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建营销方案", c)
		return
	}
	info.MerchantId = int64Ptr(int64(*mid))
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAgentLevel.UpdateAgentLevel(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAgentLevel 用id查询营销方案设置
// @Tags AgentLevel
// @Summary 用id查询营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询营销方案设置"
// @Success 200 {object} response.Response{data=model.AgentLevel,msg=string} "查询成功"
// @Router /AL/findAgentLevel [get]
func (a *AL) FindAgentLevel(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建营销方案", c)
		return
	}
	reAL, err := serviceAgentLevel.GetAgentLevel(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reAL, c)
}

// GetAgentLevelList 分页获取营销方案设置列表
// @Tags AgentLevel
// @Summary 分页获取营销方案设置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AgentLevelSearch true "分页获取营销方案设置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /AL/getAgentLevelList [get]
func (a *AL) GetAgentLevelList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.AgentLevelSearch
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建营销方案", c)
		return
	}
	pageInfo.MerchantId = int64Ptr(int64(*mid))
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceAgentLevel.GetAgentLevelInfoList(ctx, pageInfo)
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

// GetAgentLevelDataSource 获取AgentLevel的数据源
// @Tags AgentLevel
// @Summary 获取AgentLevel的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /AL/getAgentLevelDataSource [get]
func (a *AL) GetAgentLevelDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := serviceAgentLevel.GetAgentLevelDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetAgentLevelPublic 不需要鉴权的营销方案设置接口
// @Tags AgentLevel
// @Summary 不需要鉴权的营销方案设置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /AL/getAgentLevelPublic [get]
func (a *AL) GetAgentLevelPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceAgentLevel.GetAgentLevelPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的营销方案设置接口信息"}, "获取成功", c)
}
func int64Ptr(v int64) *int64 { return &v }
