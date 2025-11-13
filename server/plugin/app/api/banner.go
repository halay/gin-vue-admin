package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var Banner = new(banner)

type banner struct {}

// CreateBanner 创建Banner图片
// @Tags Banner
// @Summary 创建Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Banner true "创建Banner图片"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /banner/createBanner [post]
func (a *banner) CreateBanner(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Banner
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceBanner.CreateBanner(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteBanner 删除Banner图片
// @Tags Banner
// @Summary 删除Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Banner true "删除Banner图片"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /banner/deleteBanner [delete]
func (a *banner) DeleteBanner(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceBanner.DeleteBanner(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteBannerByIds 批量删除Banner图片
// @Tags Banner
// @Summary 批量删除Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /banner/deleteBannerByIds [delete]
func (a *banner) DeleteBannerByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceBanner.DeleteBannerByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateBanner 更新Banner图片
// @Tags Banner
// @Summary 更新Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Banner true "更新Banner图片"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /banner/updateBanner [put]
func (a *banner) UpdateBanner(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Banner
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceBanner.UpdateBanner(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindBanner 用id查询Banner图片
// @Tags Banner
// @Summary 用id查询Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询Banner图片"
// @Success 200 {object} response.Response{data=model.Banner,msg=string} "查询成功"
// @Router /banner/findBanner [get]
func (a *banner) FindBanner(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	rebanner, err := serviceBanner.GetBanner(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(rebanner, c)
}
// GetBannerList 分页获取Banner图片列表
// @Tags Banner
// @Summary 分页获取Banner图片列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.BannerSearch true "分页获取Banner图片列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /banner/getBannerList [get]
func (a *banner) GetBannerList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.BannerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceBanner.GetBannerInfoList(ctx,pageInfo)
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
// GetBannerPublic 不需要鉴权的Banner图片接口
// @Tags Banner
// @Summary 不需要鉴权的Banner图片接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /banner/getBannerPublic [get]
func (a *banner) GetBannerPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    position := c.Query("position")
    list, err := serviceBanner.GetBannerPublic(ctx, position)
    if err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:"+err.Error(), c)
        return
    }
    response.OkWithData(list, c)
}
