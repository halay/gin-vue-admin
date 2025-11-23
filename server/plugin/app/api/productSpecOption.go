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

var ProductSpecOption = new(PSO)

type PSO struct{}

// CreateProductSpecOption 创建商品规格值
// @Tags ProductSpecOption
// @Summary 创建商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpecOption true "创建商品规格值"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PSO/createProductSpecOption [post]
func (a *PSO) CreateProductSpecOption(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.ProductSpecOption
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法创建规格值", c)
		return
	}
	info.MerchantID = nil
	err = serviceProductSpecOption.CreateProductSpecOption(ctx, &info, *mid)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProductSpecOption 删除商品规格值
// @Tags ProductSpecOption
// @Summary 删除商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpecOption true "删除商品规格值"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PSO/deleteProductSpecOption [delete]
func (a *PSO) DeleteProductSpecOption(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法删除规格值", c)
		return
	}
	err := serviceProductSpecOption.DeleteProductSpecOption(ctx, ID, *mid)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductSpecOptionByIds 批量删除商品规格值
// @Tags ProductSpecOption
// @Summary 批量删除商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PSO/deleteProductSpecOptionByIds [delete]
func (a *PSO) DeleteProductSpecOptionByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法批量删除规格值", c)
		return
	}
	err := serviceProductSpecOption.DeleteProductSpecOptionByIds(ctx, IDs, *mid)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProductSpecOption 更新商品规格值
// @Tags ProductSpecOption
// @Summary 更新商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpecOption true "更新商品规格值"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PSO/updateProductSpecOption [put]
func (a *PSO) UpdateProductSpecOption(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.ProductSpecOption
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法更新规格值", c)
		return
	}
	err = serviceProductSpecOption.UpdateProductSpecOption(ctx, info, *mid)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProductSpecOption 用id查询商品规格值
// @Tags ProductSpecOption
// @Summary 用id查询商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商品规格值"
// @Success 200 {object} response.Response{data=model.ProductSpecOption,msg=string} "查询成功"
// @Router /PSO/findProductSpecOption [get]
func (a *PSO) FindProductSpecOption(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法查询规格值", c)
		return
	}
	rePSO, err := serviceProductSpecOption.GetProductSpecOption(ctx, ID, *mid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rePSO, c)
}

// GetProductSpecOptionList 分页获取商品规格值列表
// @Tags ProductSpecOption
// @Summary 分页获取商品规格值列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductSpecOptionSearch true "分页获取商品规格值列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PSO/getProductSpecOptionList [get]
func (a *PSO) GetProductSpecOptionList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.ProductSpecOptionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法获取规格值列表", c)
		return
	}
	list, total, err := serviceProductSpecOption.GetProductSpecOptionInfoList(ctx, pageInfo, *mid)
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

// GetProductSpecOptionDataSource 获取ProductSpecOption的数据源
// @Tags ProductSpecOption
// @Summary 获取ProductSpecOption的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /PSO/getProductSpecOptionDataSource [get]
func (a *PSO) GetProductSpecOptionDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	userID := utils.GetUserID(c)
	mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if errMid != nil || mid == nil {
		response.FailWithMessage("未绑定商户，无法获取数据源", c)
		return
	}
	dataSource, err := serviceProductSpecOption.GetProductSpecOptionDataSource(ctx, *mid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetProductSpecOptionPublic 不需要鉴权的商品规格值接口
// @Tags ProductSpecOption
// @Summary 不需要鉴权的商品规格值接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PSO/getProductSpecOptionPublic [get]
func (a *PSO) GetProductSpecOptionPublic(c *gin.Context) {
	ctx := c.Request.Context()
	var pageInfo request.ProductSpecOptionSearch
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.MerchantID == nil {
		response.FailWithMessage("merchantId为必填", c)
		return
	}
	list, total, err := serviceProductSpecOption.GetProductSpecOptionInfoList(ctx, pageInfo, int(*pageInfo.MerchantID))
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
