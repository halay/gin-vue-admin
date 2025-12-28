package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var AgentTransactionDetail = new(atd)

type atd struct {}

// CreateAgentTransactionDetail 创建代理交易明细
// @Tags AgentTransactionDetail
// @Summary 创建代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentTransactionDetail true "创建代理交易明细"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /atd/createAgentTransactionDetail [post]
func (a *atd) CreateAgentTransactionDetail(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.AgentTransactionDetail
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAgentTransactionDetail.CreateAgentTransactionDetail(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAgentTransactionDetail 删除代理交易明细
// @Tags AgentTransactionDetail
// @Summary 删除代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentTransactionDetail true "删除代理交易明细"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /atd/deleteAgentTransactionDetail [delete]
func (a *atd) DeleteAgentTransactionDetail(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceAgentTransactionDetail.DeleteAgentTransactionDetail(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteAgentTransactionDetailByIds 批量删除代理交易明细
// @Tags AgentTransactionDetail
// @Summary 批量删除代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /atd/deleteAgentTransactionDetailByIds [delete]
func (a *atd) DeleteAgentTransactionDetailByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceAgentTransactionDetail.DeleteAgentTransactionDetailByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateAgentTransactionDetail 更新代理交易明细
// @Tags AgentTransactionDetail
// @Summary 更新代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentTransactionDetail true "更新代理交易明细"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /atd/updateAgentTransactionDetail [put]
func (a *atd) UpdateAgentTransactionDetail(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.AgentTransactionDetail
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAgentTransactionDetail.UpdateAgentTransactionDetail(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindAgentTransactionDetail 用id查询代理交易明细
// @Tags AgentTransactionDetail
// @Summary 用id查询代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询代理交易明细"
// @Success 200 {object} response.Response{data=model.AgentTransactionDetail,msg=string} "查询成功"
// @Router /atd/findAgentTransactionDetail [get]
func (a *atd) FindAgentTransactionDetail(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reatd, err := serviceAgentTransactionDetail.GetAgentTransactionDetail(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reatd, c)
}
// GetAgentTransactionDetailList 分页获取代理交易明细列表
// @Tags AgentTransactionDetail
// @Summary 分页获取代理交易明细列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AgentTransactionDetailSearch true "分页获取代理交易明细列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /atd/getAgentTransactionDetailList [get]
func (a *atd) GetAgentTransactionDetailList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.AgentTransactionDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceAgentTransactionDetail.GetAgentTransactionDetailInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetAgentTransactionDetailPublic 不需要鉴权的代理交易明细接口
// @Tags AgentTransactionDetail
// @Summary 不需要鉴权的代理交易明细接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /atd/getAgentTransactionDetailPublic [get]
func (a *atd) GetAgentTransactionDetailPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceAgentTransactionDetail.GetAgentTransactionDetailPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的代理交易明细接口信息"}, "获取成功", c)
}
