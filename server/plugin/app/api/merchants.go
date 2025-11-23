package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Merchants = new(mc)

type mc struct{}

// CreateMerchants 创建商家信息
// @Tags Merchants
// @Summary 创建商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Merchants true "创建商家信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /mc/createMerchants [post]
func (a *mc) CreateMerchants(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.Merchants
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMerchants.CreateMerchants(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMerchants 删除商家信息
// @Tags Merchants
// @Summary 删除商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Merchants true "删除商家信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /mc/deleteMerchants [delete]
func (a *mc) DeleteMerchants(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceMerchants.DeleteMerchants(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMerchantsByIds 批量删除商家信息
// @Tags Merchants
// @Summary 批量删除商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /mc/deleteMerchantsByIds [delete]
func (a *mc) DeleteMerchantsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceMerchants.DeleteMerchantsByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMerchants 更新商家信息
// @Tags Merchants
// @Summary 更新商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Merchants true "更新商家信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /mc/updateMerchants [put]
func (a *mc) UpdateMerchants(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.Merchants
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMerchants.UpdateMerchants(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMerchants 用id查询商家信息
// @Tags Merchants
// @Summary 用id查询商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商家信息"
// @Success 200 {object} response.Response{data=model.Merchants,msg=string} "查询成功"
// @Router /mc/findMerchants [get]
func (a *mc) FindMerchants(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	remc, err := serviceMerchants.GetMerchants(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remc, c)
}

// GetMerchantsList 分页获取商家信息列表
// @Tags Merchants
// @Summary 分页获取商家信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantsSearch true "分页获取商家信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /mc/getMerchantsList [get]
func (a *mc) GetMerchantsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.MerchantsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMerchants.GetMerchantsInfoList(ctx, pageInfo)
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

// GetMerchantsPublic 不需要鉴权的商家信息接口
// @Tags Merchants
// @Summary 不需要鉴权的商家信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mc/getMerchantsPublic [get]
func (a *mc) GetMerchantsPublic(c *gin.Context) {
	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{"info": "商家公共接口"}, "获取成功", c)
}

// GetRecommendedMerchants 获取推荐商家列表（公共）
// @Tags Merchants
// @Summary 推荐商家列表
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]model.Merchants,msg=string}
// @Router /mc/getRecommendedMerchants [get]
func (a *mc) GetRecommendedMerchants(c *gin.Context) {
	ctx := c.Request.Context()
	list, err := serviceMerchants.GetRecommendedList(ctx)
	if err != nil {
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}

// GetMerchantCategoryListPublic 获取商家分类列表（公共）
// @Tags Merchants
// @Summary 商家分类列表（公共）
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]model.MerchantCategory,msg=string}
// @Router /mc/getMerchantCategoryList [get]
func (a *mc) GetMerchantCategoryList(c *gin.Context) {
	type row struct {
		ID   int64
		Name string
	}
	var cats []row
	if err := global.GVA_DB.Table("app_merchant_categories").Where("status IN (?)", []string{"1", "enabled"}).Order("sort asc, id desc").Select("id,name").Scan(&cats).Error; err != nil {
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(cats, c)
}

// GetMerchantsByCategory 通过分类获取商家列表（公共）
// @Tags Merchants
// @Summary 通过分类获取商家列表（公共）
// @Accept application/json
// @Produce application/json
// @Param categoryId query int true "分类ID"
// @Success 200 {object} response.Response{data=[]model.Merchants,msg=string}
// @Router /mc/getMerchantsByCategory [get]
func (a *mc) GetMerchantsByCategory(c *gin.Context) {
	cid := c.Query("categoryId")
	if cid == "" {
		response.FailWithMessage("缺少分类ID", c)
		return
	}
	var list []model.Merchants
	if err := global.GVA_DB.Where("category_id = ? AND status = ?", cid, "").Order("id desc").Find(&list).Error; err != nil {
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
