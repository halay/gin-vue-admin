package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	appUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var UserPointsLog = new(UPL)

type UPL struct{}

// CreateUserPointsLog 创建用户积分流水
// @Tags UserPointsLog
// @Summary 创建用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsLog true "创建用户积分流水"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /UPL/createUserPointsLog [post]
func (a *UPL) CreateUserPointsLog(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.UserPointsLog
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceUserPointsLog.CreateUserPointsLog(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserPointsLog 删除用户积分流水
// @Tags UserPointsLog
// @Summary 删除用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsLog true "删除用户积分流水"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /UPL/deleteUserPointsLog [delete]
func (a *UPL) DeleteUserPointsLog(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceUserPointsLog.DeleteUserPointsLog(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserPointsLogByIds 批量删除用户积分流水
// @Tags UserPointsLog
// @Summary 批量删除用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /UPL/deleteUserPointsLogByIds [delete]
func (a *UPL) DeleteUserPointsLogByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceUserPointsLog.DeleteUserPointsLogByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserPointsLog 更新用户积分流水
// @Tags UserPointsLog
// @Summary 更新用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsLog true "更新用户积分流水"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /UPL/updateUserPointsLog [put]
func (a *UPL) UpdateUserPointsLog(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.UserPointsLog
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceUserPointsLog.UpdateUserPointsLog(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserPointsLog 用id查询用户积分流水
// @Tags UserPointsLog
// @Summary 用id查询用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户积分流水"
// @Success 200 {object} response.Response{data=model.UserPointsLog,msg=string} "查询成功"
// @Router /UPL/findUserPointsLog [get]
func (a *UPL) FindUserPointsLog(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reUPL, err := serviceUserPointsLog.GetUserPointsLog(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reUPL, c)
}

// GetUserPointsLogList 分页获取用户积分流水列表
// @Tags UserPointsLog
// @Summary 分页获取用户积分流水列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.UserPointsLogSearch true "分页获取用户积分流水列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /UPL/getUserPointsLogList [get]
func (a *UPL) GetUserPointsLogList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.UserPointsLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := appUtils.GetUserID(c)
	mid, _ := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	pageInfo.MerchantID = mid
	list, total, err := serviceUserPointsLog.GetUserPointsLogInfoList(ctx, pageInfo)
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

// GetUserPointsLogDataSource 获取UserPointsLog的数据源
// @Tags UserPointsLog
// @Summary 获取UserPointsLog的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /UPL/getUserPointsLogDataSource [get]
func (a *UPL) GetUserPointsLogDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := serviceUserPointsLog.GetUserPointsLogDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetUserPointsLogPublic 不需要鉴权的用户积分流水接口
// @Tags UserPointsLog
// @Summary 不需要鉴权的用户积分流水接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /UPL/getUserPointsLogPublic [get]
func (a *UPL) GetUserPointsLogPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceUserPointsLog.GetUserPointsLogPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的用户积分流水接口信息"}, "获取成功", c)
}

// GetMyPointsLogs 获取当前登录用户积分流水
// @Tags UserPointsLog
// @Summary 获取当前登录用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.UserPointsLogSearch true "分页参数"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /UPL/getMyPointsLogs [get]
func (a *UPL) GetMyPointsLogs(c *gin.Context) {
	ctx := c.Request.Context()
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	var pageInfo request.UserPointsLogSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uidInt := int(uid)
	pageInfo.UserID = &uidInt
	list, total, err := serviceUserPointsLog.GetUserPointsLogInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}
