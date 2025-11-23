package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var MerchantCategory = new(MCAT)

type MCAT struct {}

// CreateMerchantCategory 创建商户分类
// @Tags MerchantCategory
// @Summary 创建商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantCategory true "创建商户分类"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /MCAT/createMerchantCategory [post]
func (a *MCAT) CreateMerchantCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MerchantCategory
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMerchantCategory.CreateMerchantCategory(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMerchantCategory 删除商户分类
// @Tags MerchantCategory
// @Summary 删除商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantCategory true "删除商户分类"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /MCAT/deleteMerchantCategory [delete]
func (a *MCAT) DeleteMerchantCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceMerchantCategory.DeleteMerchantCategory(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteMerchantCategoryByIds 批量删除商户分类
// @Tags MerchantCategory
// @Summary 批量删除商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /MCAT/deleteMerchantCategoryByIds [delete]
func (a *MCAT) DeleteMerchantCategoryByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceMerchantCategory.DeleteMerchantCategoryByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateMerchantCategory 更新商户分类
// @Tags MerchantCategory
// @Summary 更新商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantCategory true "更新商户分类"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /MCAT/updateMerchantCategory [put]
func (a *MCAT) UpdateMerchantCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MerchantCategory
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMerchantCategory.UpdateMerchantCategory(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindMerchantCategory 用id查询商户分类
// @Tags MerchantCategory
// @Summary 用id查询商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商户分类"
// @Success 200 {object} response.Response{data=model.MerchantCategory,msg=string} "查询成功"
// @Router /MCAT/findMerchantCategory [get]
func (a *MCAT) FindMerchantCategory(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reMCAT, err := serviceMerchantCategory.GetMerchantCategory(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reMCAT, c)
}
// GetMerchantCategoryList 分页获取商户分类列表
// @Tags MerchantCategory
// @Summary 分页获取商户分类列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantCategorySearch true "分页获取商户分类列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /MCAT/getMerchantCategoryList [get]
func (a *MCAT) GetMerchantCategoryList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.MerchantCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMerchantCategory.GetMerchantCategoryInfoList(ctx,pageInfo)
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
// GetMerchantCategoryPublic 不需要鉴权的商户分类接口
// @Tags MerchantCategory
// @Summary 不需要鉴权的商户分类接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MCAT/getMerchantCategoryPublic [get]
func (a *MCAT) GetMerchantCategoryPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceMerchantCategory.GetMerchantCategoryPublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的商户分类接口信息"}, "获取成功", c)
}
