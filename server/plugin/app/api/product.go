package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	commonrequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var Product = new(P)

type P struct{}

// CreateProduct 创建商户商品
// @Tags Product
// @Summary 创建商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "创建商户商品"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /P/createProduct [post]
func (a *P) CreateProduct(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.Product
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建商品", c)
		return
	}
	err = serviceProduct.CreateProduct(ctx, &info, *mid)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProduct 删除商户商品
// @Tags Product
// @Summary 删除商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "删除商户商品"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /P/deleteProduct [delete]
func (a *P) DeleteProduct(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法删除商品", c)
		return
	}
	err := serviceProduct.DeleteProduct(ctx, ID, *mid)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductByIds 批量删除商户商品
// @Tags Product
// @Summary 批量删除商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /P/deleteProductByIds [delete]
func (a *P) DeleteProductByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法批量删除商品", c)
		return
	}
	err := serviceProduct.DeleteProductByIds(ctx, IDs, *mid)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProduct 更新商户商品
// @Tags Product
// @Summary 更新商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "更新商户商品"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /P/updateProduct [put]
func (a *P) UpdateProduct(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.Product
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法更新商品", c)
		return
	}
	err = serviceProduct.UpdateProduct(ctx, info, *mid)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProduct 用id查询商户商品
// @Tags Product
// @Summary 用id查询商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商户商品"
// @Success 200 {object} response.Response{data=model.Product,msg=string} "查询成功"
// @Router /P/findProduct [get]
func (a *P) FindProduct(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法查询商品", c)
		return
	}
	reP, err := serviceProduct.GetProduct(ctx, ID, *mid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reP, c)
}

// GetProductList 分页获取商户商品列表
// @Tags Product
// @Summary 分页获取商户商品列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductSearch true "分页获取商户商品列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /P/getProductList [get]
func (a *P) GetProductList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.ProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法获取商品列表", c)
		return
	}
	list, total, err := serviceProduct.GetProductInfoList(ctx, pageInfo, *mid)
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

// GetProductDataSource 获取Product的数据源
// @Tags Product
// @Summary 获取Product的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /P/getProductDataSource [get]
func (a *P) GetProductDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法获取数据源", c)
		return
	}
	dataSource, err := serviceProduct.GetProductDataSource(ctx, *mid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetProductPublic 不需要鉴权的商户商品接口
// @Tags Product
// @Summary 不需要鉴权的商户商品接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /P/getProductPublic [get]
func (a *P) GetProductPublic(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo request.ProductSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.MerchantID == nil {
		response.FailWithMessage("merchantId为必填", c)
		return
	}
	// 如果传入ID则返回详情
	if id := c.Query("ID"); id != "" {
		prod, err := serviceProduct.GetProduct(ctx, id, int(*pageInfo.MerchantID))
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败:"+err.Error(), c)
			return
		}
		// 组装SKU、规格键、规格值
		skuSearch := request.ProductSkuSearch{ProductID: intPtr(int(prod.ID)), PageInfo: commonrequest.PageInfo{Page: 1, PageSize: 9999}}
		skus, _, _ := serviceProductSku.GetProductSkuInfoList(ctx, skuSearch, int(*pageInfo.MerchantID))
		//specSearch := request.ProductSpecSearch{ProductID: intPtr(int(prod.ID)), PageInfo: commonrequest.PageInfo{Page: 1, PageSize: 9999}}
		//specs, _, _ := serviceProductSpec.GetProductSpecInfoList(ctx, specSearch, int(*pageInfo.MerchantID))
		//specOptSearch := request.ProductSpecOptionSearch{ProductID: intPtr(int(prod.ID)), PageInfo: commonrequest.PageInfo{Page: 1, PageSize: 9999}}
		//specOptions, _, _ := serviceProductSpecOption.GetProductSpecOptionInfoList(ctx, specOptSearch, int(*pageInfo.MerchantID))
		response.OkWithDetailed(gin.H{
			"product": prod,
			"skus":    skus,
			//"specs":       specs,
			//"specOptions": specOptions,
		}, "获取成功", c)
		return
	}
	// 列表
	list, total, err := serviceProduct.GetProductInfoList(ctx, pageInfo, int(*pageInfo.MerchantID))
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

func intPtr(v int) *int { return &v }
