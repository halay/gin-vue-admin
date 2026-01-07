package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	appUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var Order = new(ORD)

type ORD struct{}

// CreateOrder 创建订单
// @Tags Order
// @Summary 创建订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "创建订单"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ORD/createOrder [post]
func (a *ORD) CreateOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.Order
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 后台商户创建/修改订单，限定商户作用域
	userID := appUtils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建订单", c)
		return
	}
	err = serviceOrder.CreateOrder(ctx, &info, *mid)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteOrder 删除订单
// @Tags Order
// @Summary 删除订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "删除订单"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ORD/deleteOrder [delete]
func (a *ORD) DeleteOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := appUtils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法删除订单", c)
		return
	}
	err := serviceOrder.DeleteOrder(ctx, ID, *mid)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteOrderByIds 批量删除订单
// @Tags Order
// @Summary 批量删除订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ORD/deleteOrderByIds [delete]
func (a *ORD) DeleteOrderByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := appUtils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法批量删除订单", c)
		return
	}
	err := serviceOrder.DeleteOrderByIds(ctx, IDs, *mid)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateOrder 更新订单
// @Tags Order
// @Summary 更新订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "更新订单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ORD/updateOrder [put]
func (a *ORD) UpdateOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.Order
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := appUtils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法更新订单", c)
		return
	}
	err = serviceOrder.UpdateOrder(ctx, info, *mid)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindOrder 用id查询订单
// @Tags Order
// @Summary 用id查询订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询订单"
// @Success 200 {object} response.Response{data=model.Order,msg=string} "查询成功"
// @Router /ORD/findOrder [get]
func (a *ORD) FindOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := appUtils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法查询订单", c)
		return
	}
	reORD, err := serviceOrder.GetOrder(ctx, ID, *mid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reORD, c)
}

// GetOrderList 分页获取订单列表
// @Tags Order
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.OrderSearch true "分页获取订单列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ORD/getOrderList [get]
func (a *ORD) GetOrderList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := appUtils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法获取订单列表", c)
		return
	}
	list, total, err := serviceOrder.GetOrderInfoList(ctx, pageInfo, *mid)
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

// GetOrderDataSource 获取Order的数据源
// @Tags Order
// @Summary 获取Order的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /ORD/getOrderDataSource [get]
func (a *ORD) GetOrderDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := serviceOrder.GetOrderDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetOrderPublic 不需要鉴权的订单接口
// @Tags Order
// @Summary 不需要鉴权的订单接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ORD/getOrderPublic [get]
func (a *ORD) GetOrderPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceOrder.GetOrderPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的订单接口信息"}, "获取成功", c)
}

// GetMyOrderList 获取当前登录用户的订单列表
// @Tags Order
// @Summary 获取当前登录用户订单列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.OrderSearch true "分页查询参数"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ORD/getMyOrderList [get]
func (a *ORD) GetMyOrderList(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo request.OrderSearch
	if err := c.ShouldBind(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := appUtils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	list, total, err := serviceOrder.GetOrderInfoListByUser(ctx, int64(userID), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	// 聚合订单明细（商品信息）
	ids := make([]int64, 0, len(list))
	mids := make([]int64, 0, len(list))
	for _, o := range list {
		ids = append(ids, int64(o.ID))
		if o.MerchantID != nil {
			mids = append(mids, *o.MerchantID)
		}
	}
	var items []model.OrderItem
	if len(ids) > 0 {
		_ = global.GVA_DB.WithContext(ctx).Where("order_id in ?", ids).Find(&items).Error
	}
	itemMap := make(map[int64][]model.OrderItem)
	for _, it := range items {
		if it.OrderID != nil {
			itemMap[*it.OrderID] = append(itemMap[*it.OrderID], it)
		}
	}
	// 聚合商户信息
	type mrow struct {
		ID           int64  `gorm:"column:id"`
		MerchantName string `gorm:"column:merchant_name"`
		Logo         string `gorm:"column:logo"`
	}
	mMap := make(map[int64]mrow)
	if len(mids) > 0 {
		var ms []mrow
		_ = global.GVA_DB.WithContext(ctx).Table("app_merchants").Where("id in ?", mids).Select("id, merchant_name, logo").Find(&ms).Error
		for _, m := range ms {
			mMap[m.ID] = m
		}
	}
	enriched := make([]gin.H, 0, len(list))
	for _, o := range list {
		var mid int64
		if o.MerchantID != nil {
			mid = *o.MerchantID
		}
		mi := mMap[mid]
		enriched = append(enriched, gin.H{
			"order": o,
			"items": itemMap[int64(o.ID)],
			"merchant": gin.H{
				"id":           mid,
				"merchantName": mi.MerchantName,
				"logo":         mi.Logo,
			},
		})
	}
	response.OkWithDetailed(response.PageResult{List: enriched, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}

// GetMyOrderDetail 获取当前登录用户的订单详情
// @Tags Order
// @Summary 获取当前登录用户的订单详情
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param orderNo query string true "订单号"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ORD/getMyOrderDetail [get]
func (a *ORD) GetMyOrderDetail(c *gin.Context) {
	ctx := c.Request.Context()
	orderNo := c.Query("orderNo")
	userID := appUtils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	if orderNo == "" {
		response.FailWithMessage("订单号不能为空", c)
		return
	}
	ord, items, err := serviceOrder.GetOrderDetailByUser(ctx, int64(userID), orderNo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"order": ord, "items": items}, "获取成功", c)
}

// CreateOrderByPoints C端：创建积分订单（未支付）
// @Tags Order
// @Summary C端创建订单（points或card）
// @Accept application/json
// @Produce application/json
// @Param skuId body string true "SKU ID"
// @Param quantity body string true "购买数量"
// @Param consigneeName body string false "收货人"
// @Param consigneePhone body string false "收货电话"
// @Param address body string false "收货地址"
// @Param province body string false "省"
// @Param city body string false "市"
// @Param district body string false "区县"
// @Param payMethod body string false "支付方式 points|card"
// @Success 200 {object} response.Response{data=object,msg=string} "创建成功"
// @Router /ORD/createOrderByPoints [post]
func (a *ORD) CreateOrderByPoints(c *gin.Context) {
	ctx := c.Request.Context()
	userID := appUtils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	var body struct {
		SkuID          int64  `json:"skuId"`
		Quantity       int64  `json:"quantity"`
		ConsigneeName  string `json:"consigneeName"`
		ConsigneePhone string `json:"consigneePhone"`
		Address        string `json:"address"`
		Province       string `json:"province"`
		City           string `json:"city"`
		District       string `json:"district"`
		Country        string `json:"country"`
		PostalCode     string `json:"postalCode"`
		PayMethod      string `json:"payMethod"`
		Remark         string `json:"remark"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if body.Quantity <= 0 {
		body.Quantity = 1
	}

	var sku model.ProductSku
	if err := global.GVA_DB.WithContext(ctx).Where("id = ?", body.SkuID).First(&sku).Error; err != nil {
		response.FailWithMessage("SKU不存在", c)
		return
	}
	pm := body.PayMethod
	if pm == "" {
		pm = "card"
	}
	ord, item, err := serviceOrder.CreateOrderByPoints(ctx, int64(userID), sku, body.Quantity, body.ConsigneeName, body.ConsigneePhone, body.Address, body.Province, body.City, body.District, body.Country, body.PostalCode, pm, body.Remark)
	if err != nil {
		global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
		response.FailWithMessage("创建订单失败:"+err.Error(), c)
		return
	}
	// 若为card支付，立即创建支付意图并返回
	if pm != "" && ord.OrderNo != nil {
		piID, clientSecret, err2 := serviceOrder.CreateOrderPaymentIntent(ctx, uint(userID), *ord.OrderNo, pm)
		if err2 != nil {
			global.GVA_LOG.Error("创建支付意图失败!", zap.Error(err2))
			response.FailWithMessage("创建支付意图失败:"+err2.Error(), c)
			return
		}
		response.OkWithData(gin.H{"order": ord, "item": item, "paymentIntentId": piID, "clientSecret": clientSecret}, c)
		return
	}
	response.OkWithData(gin.H{"order": ord, "item": item}, c)
}

// PayOrderByPoints C端：订单积分支付
// @Tags Order
// @Summary C端订单积分支付
// @Accept application/json
// @Produce application/json
// @Param orderNo body string true "订单号"
// @Success 200 {object} response.Response{msg=string} "支付成功"
// @Router /ORD/payOrderByPoints [post]
func (a *ORD) PayOrderByPoints(c *gin.Context) {
	ctx := c.Request.Context()
	userID := appUtils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	var body struct {
		OrderNo string `json:"orderNo"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if body.OrderNo == "" {
		response.FailWithMessage("订单号不能为空", c)
		return
	}
	if err := serviceOrder.PayOrderByPoints(ctx, int64(userID), body.OrderNo); err != nil {
		global.GVA_LOG.Error("积分支付失败!", zap.Error(err))
		response.FailWithMessage("积分支付失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("支付成功", c)
}

// CreateOrderPaymentIntent C端：订单银行卡支付（Stripe）
// @Tags Order
// @Summary C端订单银行卡支付：创建支付意图
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param orderNo body string true "订单号"
// @Success 200 {object} response.Response{data=object,msg=string} "创建成功"
// @Router /ORD/createOrderPaymentIntent [post]
func (a *ORD) CreateOrderPaymentIntent(c *gin.Context) {
	ctx := c.Request.Context()
	userID := appUtils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	var body struct {
		OrderNo   string `json:"orderNo"`
		PayMethod string `json:"payMethod"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if body.OrderNo == "" {
		response.FailWithMessage("订单号不能为空", c)
		return
	}
	if body.PayMethod == "" {
		body.PayMethod = "card"
	}
	id, clientSecret, err := serviceOrder.CreateOrderPaymentIntent(ctx, uint(userID), body.OrderNo, body.PayMethod)
	if err != nil {
		global.GVA_LOG.Error("创建支付意图失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"paymentIntentId": id, "clientSecret": clientSecret}, c)
}
