package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var MerchantAdmin = new(MADM)

type MADM struct{}

// CreateMerchantAdmin 创建后台用户与商户关联映射
// @Tags MerchantAdmin
// @Summary 创建后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAdmin true "创建后台用户与商户关联映射"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /MADM/createMerchantAdmin [post]
func (a *MADM) CreateMerchantAdmin(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.MerchantAdmin
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMerchantAdmin.CreateMerchantAdmin(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMerchantAdmin 删除后台用户与商户关联映射
// @Tags MerchantAdmin
// @Summary 删除后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAdmin true "删除后台用户与商户关联映射"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /MADM/deleteMerchantAdmin [delete]
func (a *MADM) DeleteMerchantAdmin(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceMerchantAdmin.DeleteMerchantAdmin(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMerchantAdminByIds 批量删除后台用户与商户关联映射
// @Tags MerchantAdmin
// @Summary 批量删除后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /MADM/deleteMerchantAdminByIds [delete]
func (a *MADM) DeleteMerchantAdminByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceMerchantAdmin.DeleteMerchantAdminByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMerchantAdmin 更新后台用户与商户关联映射
// @Tags MerchantAdmin
// @Summary 更新后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAdmin true "更新后台用户与商户关联映射"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /MADM/updateMerchantAdmin [put]
func (a *MADM) UpdateMerchantAdmin(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.MerchantAdmin
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMerchantAdmin.UpdateMerchantAdmin(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMerchantAdmin 用id查询后台用户与商户关联映射
// @Tags MerchantAdmin
// @Summary 用id查询后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询后台用户与商户关联映射"
// @Success 200 {object} response.Response{data=model.MerchantAdmin,msg=string} "查询成功"
// @Router /MADM/findMerchantAdmin [get]
func (a *MADM) FindMerchantAdmin(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reMADM, err := serviceMerchantAdmin.GetMerchantAdmin(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reMADM, c)
}

// GetMerchantAdminList 分页获取后台用户与商户关联映射列表
// @Tags MerchantAdmin
// @Summary 分页获取后台用户与商户关联映射列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantAdminSearch true "分页获取后台用户与商户关联映射列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /MADM/getMerchantAdminList [get]
func (a *MADM) GetMerchantAdminList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.MerchantAdminSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMerchantAdmin.GetMerchantAdminInfoList(ctx, pageInfo)
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

// GetMerchantAdminDataSource 获取MerchantAdmin的数据源
// @Tags MerchantAdmin
// @Summary 获取MerchantAdmin的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /MADM/getMerchantAdminDataSource [get]
func (a *MADM) GetMerchantAdminDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := serviceMerchantAdmin.GetMerchantAdminDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetMerchantAdminPublic 不需要鉴权的后台用户与商户关联映射接口
// @Tags MerchantAdmin
// @Summary 不需要鉴权的后台用户与商户关联映射接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MADM/getMerchantAdminPublic [get]
func (a *MADM) GetMerchantAdminPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceMerchantAdmin.GetMerchantAdminPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的后台用户与商户关联映射接口信息"}, "获取成功", c)
}
