package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/plugin"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

var AppRelease = new(AR)

type AR struct{}

// CreateAppRelease 创建App客户端版本升级
// @Tags AppRelease
// @Summary 创建App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppRelease true "创建App客户端版本升级"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /AR/createAppRelease [post]
func (a *AR) CreateAppRelease(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.AppRelease
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAppRelease.CreateAppRelease(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAppRelease 删除App客户端版本升级
// @Tags AppRelease
// @Summary 删除App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppRelease true "删除App客户端版本升级"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /AR/deleteAppRelease [delete]
func (a *AR) DeleteAppRelease(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceAppRelease.DeleteAppRelease(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAppReleaseByIds 批量删除App客户端版本升级
// @Tags AppRelease
// @Summary 批量删除App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /AR/deleteAppReleaseByIds [delete]
func (a *AR) DeleteAppReleaseByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceAppRelease.DeleteAppReleaseByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAppRelease 更新App客户端版本升级
// @Tags AppRelease
// @Summary 更新App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppRelease true "更新App客户端版本升级"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /AR/updateAppRelease [put]
func (a *AR) UpdateAppRelease(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.AppRelease
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAppRelease.UpdateAppRelease(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAppRelease 用id查询App客户端版本升级
// @Tags AppRelease
// @Summary 用id查询App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询App客户端版本升级"
// @Success 200 {object} response.Response{data=model.AppRelease,msg=string} "查询成功"
// @Router /AR/findAppRelease [get]
func (a *AR) FindAppRelease(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reAR, err := serviceAppRelease.GetAppRelease(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reAR, c)
}

// GetAppReleaseList 分页获取App客户端版本升级列表
// @Tags AppRelease
// @Summary 分页获取App客户端版本升级列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AppReleaseSearch true "分页获取App客户端版本升级列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /AR/getAppReleaseList [get]
func (a *AR) GetAppReleaseList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.AppReleaseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceAppRelease.GetAppReleaseInfoList(ctx, pageInfo)
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

// GetAppReleasePublic 不需要鉴权的App客户端版本升级接口
// @Tags AppRelease
// @Summary 不需要鉴权的App客户端版本升级接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /AR/getAppReleasePublic [get]
func (a *AR) GetAppReleasePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	platform := c.Query("platform")
	vcStr := c.Query("versionCode")
	if platform == "" || vcStr == "" {
		response.FailWithMessage("缺少参数", c)
		return
	}
	vc, _ := strconv.ParseInt(vcStr, 10, 64)
	need, rel, err := serviceAppRelease.CheckUpdate(ctx, platform, vc)
	if err != nil {
		response.FailWithMessage("检查失败:"+err.Error(), c)
		return
	}
	var hrefUrl = plugin.Config.HrefUrl
	var androidUrl = *rel.AndroidUrl
	if androidUrl != "" {
		androidUrl = hrefUrl + *rel.AndroidUrl
	}
	var iosUrl = *rel.IosUrl
	if iosUrl != "" {
		iosUrl = hrefUrl + *rel.IosUrl
	}
	// 返回必要信息
	data := gin.H{
		"need":        need,
		"versionCode": rel.VersionCode,
		"versionName": rel.VersionName,
		"forceUpdate": rel.ForceUpdate,
		"minSupport":  rel.MinSupport,
		"androidUrl":  androidUrl,
		"iosUrl":      iosUrl,
		"changelog":   rel.Changelog,
	}
	response.OkWithData(data, c)
}
