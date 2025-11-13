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

var ProductCategory = new(PC)

type PC struct {}

// CreateProductCategory 创建商户商品分类
// @Tags ProductCategory
// @Summary 创建商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductCategory true "创建商户商品分类"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PC/createProductCategory [post]
func (a *PC) CreateProductCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var info model.ProductCategory
    err := c.ShouldBindJSON(&info)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法创建分类", c)
        return
    }
    err = serviceProductCategory.CreateProductCategory(ctx, &info, *mid)
    if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("创建成功", c)
}

// DeleteProductCategory 删除商户商品分类
// @Tags ProductCategory
// @Summary 删除商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductCategory true "删除商户商品分类"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PC/deleteProductCategory [delete]
func (a *PC) DeleteProductCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    ID := c.Query("ID")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法删除分类", c)
        return
    }
    err := serviceProductCategory.DeleteProductCategory(ctx, ID, *mid)
    if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("删除成功", c)
}

// DeleteProductCategoryByIds 批量删除商户商品分类
// @Tags ProductCategory
// @Summary 批量删除商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PC/deleteProductCategoryByIds [delete]
func (a *PC) DeleteProductCategoryByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法批量删除分类", c)
        return
    }
    err := serviceProductCategory.DeleteProductCategoryByIds(ctx, IDs, *mid)
    if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("批量删除成功", c)
}

// UpdateProductCategory 更新商户商品分类
// @Tags ProductCategory
// @Summary 更新商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductCategory true "更新商户商品分类"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PC/updateProductCategory [put]
func (a *PC) UpdateProductCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var info model.ProductCategory
    err := c.ShouldBindJSON(&info)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法更新分类", c)
        return
    }
    err = serviceProductCategory.UpdateProductCategory(ctx, info, *mid)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("更新成功", c)
}

// FindProductCategory 用id查询商户商品分类
// @Tags ProductCategory
// @Summary 用id查询商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商户商品分类"
// @Success 200 {object} response.Response{data=model.ProductCategory,msg=string} "查询成功"
// @Router /PC/findProductCategory [get]
func (a *PC) FindProductCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    ID := c.Query("ID")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法查询分类", c)
        return
    }
    rePC, err := serviceProductCategory.GetProductCategory(ctx, ID, *mid)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
        return
    }
    response.OkWithData(rePC, c)
}
// GetProductCategoryList 分页获取商户商品分类列表
// @Tags ProductCategory
// @Summary 分页获取商户商品分类列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductCategorySearch true "分页获取商户商品分类列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PC/getProductCategoryList [get]
func (a *PC) GetProductCategoryList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var pageInfo request.ProductCategorySearch
    err := c.ShouldBindQuery(&pageInfo)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法获取分类列表", c)
        return
    }
    list, total, err := serviceProductCategory.GetProductCategoryInfoList(ctx, pageInfo, *mid)
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
// GetProductCategoryPublic 不需要鉴权的商户商品分类接口
// @Tags ProductCategory
// @Summary 不需要鉴权的商户商品分类接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PC/getProductCategoryPublic [get]
func (a *PC) GetProductCategoryPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceProductCategory.GetProductCategoryPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的商户商品分类接口信息"}, "获取成功", c)
}
