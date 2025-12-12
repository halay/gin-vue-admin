package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	appUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var PointsRecharge = new(PR)

type PR struct{}

type CreateRechargeRequest struct {
	ConfigID  uint   `json:"configId" binding:"required"`
	PayMethod string `json:"payMethod" binding:"required"`
}

// CreateRechargeOrder 生成充值订单（需登录）
func (a *PR) CreateRechargeOrder(c *gin.Context) {
	ctx := c.Request.Context()
	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	var req CreateRechargeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ord, cfg, err := servicePointsRecharge.CreateRechargeOrder(ctx, int64(uid), req.ConfigID, req.PayMethod)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"order": ord, "config": cfg}, c)
}

type PayCallbackRequest struct {
	OrderNo       string `json:"orderNo" binding:"required"`
	Success       bool   `json:"success"`
	TransactionID string `json:"transactionId"`
}

// PayCallback 支付回调（公共，无需鉴权）
func (a *PR) PayCallback(c *gin.Context) {
	ctx := c.Request.Context()
	var req PayCallbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := servicePointsRecharge.PayCallback(ctx, req.OrderNo, req.Success, req.TransactionID); err != nil {
		global.GVA_LOG.Error("回调失败!", zap.Error(err))
		response.FailWithMessage("回调失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("回调处理成功", c)
}

// GetMyRechargeOrders 获取当前登录用户的积分充值订单列表
func (a *PR) GetMyRechargeOrders(c *gin.Context) {
	ctx := c.Request.Context()
	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	var pageInfo request.PointsRechargeSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := servicePointsRecharge.GetRechargeOrdersByUser(ctx, int64(uid), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}

// GetMyRechargeOrder 获取当前登录用户的单条充值订单
func (a *PR) GetMyRechargeOrder(c *gin.Context) {
	ctx := c.Request.Context()
	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	orderNo := c.Query("orderNo")
	id := c.Query("ID")
	re, err := servicePointsRecharge.GetRechargeOrderByUser(ctx, int64(uid), orderNo, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(re, c)
}
