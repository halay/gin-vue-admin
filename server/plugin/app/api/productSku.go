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

var ProductSku = new(SKU)

type SKU struct {}

// CreateProductSku 创建商品SKU（规格组合）
// @Tags ProductSku
// @Summary 创建商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSku true "创建商品SKU（规格组合）"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /SKU/createProductSku [post]
func (a *SKU) CreateProductSku(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.ProductSku
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法创建SKU", c)
        return
    }
    err = serviceProductSku.CreateProductSku(ctx, &info, *mid)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteProductSku 删除商品SKU（规格组合）
// @Tags ProductSku
// @Summary 删除商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSku true "删除商品SKU（规格组合）"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /SKU/deleteProductSku [delete]
func (a *SKU) DeleteProductSku(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法删除SKU", c)
        return
    }
    err := serviceProductSku.DeleteProductSku(ctx, ID, *mid)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteProductSkuByIds 批量删除商品SKU（规格组合）
// @Tags ProductSku
// @Summary 批量删除商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /SKU/deleteProductSkuByIds [delete]
func (a *SKU) DeleteProductSkuByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法批量删除SKU", c)
        return
    }
    err := serviceProductSku.DeleteProductSkuByIds(ctx, IDs, *mid)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateProductSku 更新商品SKU（规格组合）
// @Tags ProductSku
// @Summary 更新商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSku true "更新商品SKU（规格组合）"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /SKU/updateProductSku [put]
func (a *SKU) UpdateProductSku(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.ProductSku
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法更新SKU", c)
        return
    }
    err = serviceProductSku.UpdateProductSku(ctx, info, *mid)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindProductSku 用id查询商品SKU（规格组合）
// @Tags ProductSku
// @Summary 用id查询商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商品SKU（规格组合）"
// @Success 200 {object} response.Response{data=model.ProductSku,msg=string} "查询成功"
// @Router /SKU/findProductSku [get]
func (a *SKU) FindProductSku(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法查询SKU", c)
        return
    }
    reSKU, err := serviceProductSku.GetProductSku(ctx, ID, *mid)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reSKU, c)
}
// GetProductSkuList 分页获取商品SKU（规格组合）列表
// @Tags ProductSku
// @Summary 分页获取商品SKU（规格组合）列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductSkuSearch true "分页获取商品SKU（规格组合）列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /SKU/getProductSkuList [get]
func (a *SKU) GetProductSkuList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.ProductSkuSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法获取SKU列表", c)
        return
    }
    list, total, err := serviceProductSku.GetProductSkuInfoList(ctx, pageInfo, *mid)
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
// GetProductSkuDataSource 获取ProductSku的数据源
// @Tags ProductSku
// @Summary 获取ProductSku的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /SKU/getProductSkuDataSource [get]
func (a *SKU) GetProductSkuDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
   userID := utils.GetUserID(c)
   mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
   if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法获取数据源", c)
        return
   }
   dataSource, err := serviceProductSku.GetProductSkuDataSource(ctx, *mid)
   if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
		return
   }
    response.OkWithData(dataSource, c)
}
// GetProductSkuPublic 不需要鉴权的商品SKU（规格组合）接口
// @Tags ProductSku
// @Summary 不需要鉴权的商品SKU（规格组合）接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /SKU/getProductSkuPublic [get]
func (a *SKU) GetProductSkuPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceProductSku.GetProductSkuPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的商品SKU（规格组合）接口信息"}, "获取成功", c)
}
