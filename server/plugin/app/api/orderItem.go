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

var OrderItem = new(ORDI)

type ORDI struct{}

// CreateOrderItem 创建订单明细
// @Tags OrderItem
// @Summary 创建订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "创建订单明细"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ORDI/createOrderItem [post]
func (a *ORDI) CreateOrderItem(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.OrderItem
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceOrderItem.CreateOrderItem(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteOrderItem 删除订单明细
// @Tags OrderItem
// @Summary 删除订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "删除订单明细"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ORDI/deleteOrderItem [delete]
func (a *ORDI) DeleteOrderItem(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceOrderItem.DeleteOrderItem(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteOrderItemByIds 批量删除订单明细
// @Tags OrderItem
// @Summary 批量删除订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ORDI/deleteOrderItemByIds [delete]
func (a *ORDI) DeleteOrderItemByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceOrderItem.DeleteOrderItemByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateOrderItem 更新订单明细
// @Tags OrderItem
// @Summary 更新订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "更新订单明细"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ORDI/updateOrderItem [put]
func (a *ORDI) UpdateOrderItem(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.OrderItem
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceOrderItem.UpdateOrderItem(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindOrderItem 用id查询订单明细
// @Tags OrderItem
// @Summary 用id查询订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询订单明细"
// @Success 200 {object} response.Response{data=model.OrderItem,msg=string} "查询成功"
// @Router /ORDI/findOrderItem [get]
func (a *ORDI) FindOrderItem(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reORDI, err := serviceOrderItem.GetOrderItem(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reORDI, c)
}

// GetOrderItemList 分页获取订单明细列表
// @Tags OrderItem
// @Summary 分页获取订单明细列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.OrderItemSearch true "分页获取订单明细列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ORDI/getOrderItemList [get]
func (a *ORDI) GetOrderItemList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.OrderItemSearch
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
	pageInfo.MerchantID = mid

	list, total, err := serviceOrderItem.GetOrderItemInfoList(ctx, pageInfo)
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

// GetOrderItemDataSource 获取OrderItem的数据源
// @Tags OrderItem
// @Summary 获取OrderItem的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /ORDI/getOrderItemDataSource [get]
func (a *ORDI) GetOrderItemDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := serviceOrderItem.GetOrderItemDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetOrderItemPublic 不需要鉴权的订单明细接口
// @Tags OrderItem
// @Summary 不需要鉴权的订单明细接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ORDI/getOrderItemPublic [get]
func (a *ORDI) GetOrderItemPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceOrderItem.GetOrderItemPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的订单明细接口信息"}, "获取成功", c)
}
