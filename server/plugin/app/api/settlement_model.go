package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var Settlement = new(settlement)

type settlement struct {}

// CreateSettlement 创建结算管理
// @Tags Settlement
// @Summary 创建结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Settlement true "创建结算管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /settlement/createSettlement [post]
func (a *settlement) CreateSettlement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Settlement
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSettlement.CreateSettlement(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSettlement 删除结算管理
// @Tags Settlement
// @Summary 删除结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Settlement true "删除结算管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /settlement/deleteSettlement [delete]
func (a *settlement) DeleteSettlement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceSettlement.DeleteSettlement(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteSettlementByIds 批量删除结算管理
// @Tags Settlement
// @Summary 批量删除结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /settlement/deleteSettlementByIds [delete]
func (a *settlement) DeleteSettlementByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceSettlement.DeleteSettlementByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateSettlement 更新结算管理
// @Tags Settlement
// @Summary 更新结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Settlement true "更新结算管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /settlement/updateSettlement [put]
func (a *settlement) UpdateSettlement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Settlement
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceSettlement.UpdateSettlement(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindSettlement 用id查询结算管理
// @Tags Settlement
// @Summary 用id查询结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询结算管理"
// @Success 200 {object} response.Response{data=model.Settlement,msg=string} "查询成功"
// @Router /settlement/findSettlement [get]
func (a *settlement) FindSettlement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	resettlement, err := serviceSettlement.GetSettlement(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(resettlement, c)
}
// GetSettlementList 分页获取结算管理列表
// @Tags Settlement
// @Summary 分页获取结算管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.SettlementSearch true "分页获取结算管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /settlement/getSettlementList [get]
func (a *settlement) GetSettlementList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.SettlementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceSettlement.GetSettlementInfoList(ctx,pageInfo)
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
// GetSettlementPublic 不需要鉴权的结算管理接口
// @Tags Settlement
// @Summary 不需要鉴权的结算管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /settlement/getSettlementPublic [get]
func (a *settlement) GetSettlementPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceSettlement.GetSettlementPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的结算管理接口信息"}, "获取成功", c)
}
