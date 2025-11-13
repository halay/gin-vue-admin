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

var ProductSpec = new(PS)

type PS struct {}

// CreateProductSpec 创建商品规格键
// @Tags ProductSpec
// @Summary 创建商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpec true "创建商品规格键"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PS/createProductSpec [post]
func (a *PS) CreateProductSpec(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.ProductSpec
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法创建规格", c)
        return
    }
    err = serviceProductSpec.CreateProductSpec(ctx, &info, *mid)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteProductSpec 删除商品规格键
// @Tags ProductSpec
// @Summary 删除商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpec true "删除商品规格键"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PS/deleteProductSpec [delete]
func (a *PS) DeleteProductSpec(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法删除规格", c)
        return
    }
    err := serviceProductSpec.DeleteProductSpec(ctx, ID, *mid)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteProductSpecByIds 批量删除商品规格键
// @Tags ProductSpec
// @Summary 批量删除商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PS/deleteProductSpecByIds [delete]
func (a *PS) DeleteProductSpecByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法批量删除规格", c)
        return
    }
    err := serviceProductSpec.DeleteProductSpecByIds(ctx, IDs, *mid)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateProductSpec 更新商品规格键
// @Tags ProductSpec
// @Summary 更新商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpec true "更新商品规格键"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PS/updateProductSpec [put]
func (a *PS) UpdateProductSpec(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.ProductSpec
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法更新规格", c)
        return
    }
    err = serviceProductSpec.UpdateProductSpec(ctx, info, *mid)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindProductSpec 用id查询商品规格键
// @Tags ProductSpec
// @Summary 用id查询商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商品规格键"
// @Success 200 {object} response.Response{data=model.ProductSpec,msg=string} "查询成功"
// @Router /PS/findProductSpec [get]
func (a *PS) FindProductSpec(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法查询规格", c)
        return
    }
    rePS, err := serviceProductSpec.GetProductSpec(ctx, ID, *mid)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(rePS, c)
}
// GetProductSpecList 分页获取商品规格键列表
// @Tags ProductSpec
// @Summary 分页获取商品规格键列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductSpecSearch true "分页获取商品规格键列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PS/getProductSpecList [get]
func (a *PS) GetProductSpecList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.ProductSpecSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法获取规格列表", c)
        return
    }
    list, total, err := serviceProductSpec.GetProductSpecInfoList(ctx, pageInfo, *mid)
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
// GetProductSpecDataSource 获取ProductSpec的数据源
// @Tags ProductSpec
// @Summary 获取ProductSpec的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /PS/getProductSpecDataSource [get]
func (a *PS) GetProductSpecDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
   userID := utils.GetUserID(c)
   mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
   if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法获取数据源", c)
        return
   }
   dataSource, err := serviceProductSpec.GetProductSpecDataSource(ctx, *mid)
   if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
		return
   }
    response.OkWithData(dataSource, c)
}
// GetProductSpecPublic 不需要鉴权的商品规格键接口
// @Tags ProductSpec
// @Summary 不需要鉴权的商品规格键接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PS/getProductSpecPublic [get]
func (a *PS) GetProductSpecPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceProductSpec.GetProductSpecPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的商品规格键接口信息"}, "获取成功", c)
}
