package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var PointsSettings = new(PTS)

type PTS struct {}

// CreatePointsSettings 创建平台积分基础设置与奖励规则
// @Tags PointsSettings
// @Summary 创建平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PointsSettings true "创建平台积分基础设置与奖励规则"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PTS/createPointsSettings [post]
func (a *PTS) CreatePointsSettings(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.PointsSettings
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = servicePointsSettings.CreatePointsSettings(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePointsSettings 删除平台积分基础设置与奖励规则
// @Tags PointsSettings
// @Summary 删除平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PointsSettings true "删除平台积分基础设置与奖励规则"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PTS/deletePointsSettings [delete]
func (a *PTS) DeletePointsSettings(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := servicePointsSettings.DeletePointsSettings(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeletePointsSettingsByIds 批量删除平台积分基础设置与奖励规则
// @Tags PointsSettings
// @Summary 批量删除平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PTS/deletePointsSettingsByIds [delete]
func (a *PTS) DeletePointsSettingsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := servicePointsSettings.DeletePointsSettingsByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdatePointsSettings 更新平台积分基础设置与奖励规则
// @Tags PointsSettings
// @Summary 更新平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PointsSettings true "更新平台积分基础设置与奖励规则"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PTS/updatePointsSettings [put]
func (a *PTS) UpdatePointsSettings(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.PointsSettings
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = servicePointsSettings.UpdatePointsSettings(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindPointsSettings 用id查询平台积分基础设置与奖励规则
// @Tags PointsSettings
// @Summary 用id查询平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询平台积分基础设置与奖励规则"
// @Success 200 {object} response.Response{data=model.PointsSettings,msg=string} "查询成功"
// @Router /PTS/findPointsSettings [get]
func (a *PTS) FindPointsSettings(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rePTS, err := servicePointsSettings.GetPointsSettings(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(rePTS, c)
}
// GetPointsSettingsList 分页获取平台积分基础设置与奖励规则列表
// @Tags PointsSettings
// @Summary 分页获取平台积分基础设置与奖励规则列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PointsSettingsSearch true "分页获取平台积分基础设置与奖励规则列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PTS/getPointsSettingsList [get]
func (a *PTS) GetPointsSettingsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.PointsSettingsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := servicePointsSettings.GetPointsSettingsInfoList(ctx,pageInfo)
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
// GetPointsSettingsPublic 不需要鉴权的平台积分基础设置与奖励规则接口
// @Tags PointsSettings
// @Summary 不需要鉴权的平台积分基础设置与奖励规则接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PTS/getPointsSettingsPublic [get]
func (a *PTS) GetPointsSettingsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    servicePointsSettings.GetPointsSettingsPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的平台积分基础设置与奖励规则接口信息"}, "获取成功", c)
}
