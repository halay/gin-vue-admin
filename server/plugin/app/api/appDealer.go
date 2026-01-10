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

var AppDealer = new(ADL)

type ADL struct {}

// CreateAppDealer 创建经销商管理
// @Tags AppDealer
// @Summary 创建经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppDealer true "创建经销商管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ADL/createAppDealer [post]
func (a *ADL) CreateAppDealer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.AppDealer
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

    // 自动绑定商户ID
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法创建", c)
        return
    }
    mid64 := int64(*mid)
    info.MerchantId = &mid64

	err = serviceAppDealer.CreateAppDealer(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAppDealer 删除经销商管理
// @Tags AppDealer
// @Summary 删除经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppDealer true "删除经销商管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ADL/deleteAppDealer [delete]
func (a *ADL) DeleteAppDealer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceAppDealer.DeleteAppDealer(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteAppDealerByIds 批量删除经销商管理
// @Tags AppDealer
// @Summary 批量删除经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ADL/deleteAppDealerByIds [delete]
func (a *ADL) DeleteAppDealerByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceAppDealer.DeleteAppDealerByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateAppDealer 更新经销商管理
// @Tags AppDealer
// @Summary 更新经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppDealer true "更新经销商管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ADL/updateAppDealer [put]
func (a *ADL) UpdateAppDealer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.AppDealer
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAppDealer.UpdateAppDealer(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindAppDealer 用id查询经销商管理
// @Tags AppDealer
// @Summary 用id查询经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询经销商管理"
// @Success 200 {object} response.Response{data=model.AppDealer,msg=string} "查询成功"
// @Router /ADL/findAppDealer [get]
func (a *ADL) FindAppDealer(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reADL, err := serviceAppDealer.GetAppDealer(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reADL, c)
}
// GetAppDealerList 分页获取经销商管理列表
// @Tags AppDealer
// @Summary 分页获取经销商管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AppDealerSearch true "分页获取经销商管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ADL/getAppDealerList [get]
func (a *ADL) GetAppDealerList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.AppDealerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceAppDealer.GetAppDealerInfoList(ctx,pageInfo)
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
// GetAppDealerDataSource 获取AppDealer的数据源
// @Tags AppDealer
// @Summary 获取AppDealer的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /ADL/getAppDealerDataSource [get]
func (a *ADL) GetAppDealerDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
   dataSource, err := serviceAppDealer.GetAppDealerDataSource(ctx)
   if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
		return
   }
    response.OkWithData(dataSource, c)
}
// GetAppDealerPublic 不需要鉴权的经销商管理接口
// @Tags AppDealer
// @Summary 不需要鉴权的经销商管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ADL/getAppDealerPublic [get]
func (a *ADL) GetAppDealerPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceAppDealer.GetAppDealerPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的经销商管理接口信息"}, "获取成功", c)
}
