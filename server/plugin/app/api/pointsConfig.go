package api

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var PointsCfg = new(PCFG)

type PCFG struct{}

// CreatePointsConfig 新建积分配置
func (a *PCFG) CreatePointsConfig(c *gin.Context) {
    ctx := c.Request.Context()
    var info model.PointsConfig
    if err := c.ShouldBindJSON(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
    if err := servicePointsCfg.CreatePointsConfig(ctx, &info); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败:"+err.Error(), c)
        return
    }
    response.OkWithMessage("创建成功", c)
}

// DeletePointsConfig 删除积分配置
func (a *PCFG) DeletePointsConfig(c *gin.Context) {
    ctx := c.Request.Context()
    ID := c.Query("ID")
    if err := servicePointsCfg.DeletePointsConfig(ctx, ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败:"+err.Error(), c)
        return
    }
    response.OkWithMessage("删除成功", c)
}

// DeletePointsConfigByIds 批量删除
func (a *PCFG) DeletePointsConfigByIds(c *gin.Context) {
    ctx := c.Request.Context()
    IDs := c.QueryArray("IDs[]")
    if err := servicePointsCfg.DeletePointsConfigByIds(ctx, IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败:"+err.Error(), c)
        return
    }
    response.OkWithMessage("批量删除成功", c)
}

// UpdatePointsConfig 更新
func (a *PCFG) UpdatePointsConfig(c *gin.Context) {
    ctx := c.Request.Context()
    var info model.PointsConfig
    if err := c.ShouldBindJSON(&info); err != nil { response.FailWithMessage(err.Error(), c); return }
    if err := servicePointsCfg.UpdatePointsConfig(ctx, info); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败:"+err.Error(), c)
        return
    }
    response.OkWithMessage("更新成功", c)
}

// FindPointsConfig 查询详情
func (a *PCFG) FindPointsConfig(c *gin.Context) {
    ctx := c.Request.Context()
    ID := c.Query("ID")
    re, err := servicePointsCfg.GetPointsConfig(ctx, ID)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:"+err.Error(), c)
        return
    }
    response.OkWithData(re, c)
}

// GetPointsConfigList 后台列表
func (a *PCFG) GetPointsConfigList(c *gin.Context) {
    ctx := c.Request.Context()
    var pageInfo request.PointsConfigSearch
    if err := c.ShouldBindQuery(&pageInfo); err != nil { response.FailWithMessage(err.Error(), c); return }
    list, total, err := servicePointsCfg.GetPointsConfigInfoList(ctx, pageInfo)
    if err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:"+err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{List:list, Total:total, Page:pageInfo.Page, PageSize:pageInfo.PageSize}, "获取成功", c)
}

// GetPointsConfigPublic APP端公共获取配置
func (a *PCFG) GetPointsConfigPublic(c *gin.Context) {
    ctx := c.Request.Context()
    var pageInfo request.PointsConfigSearch
    if err := c.ShouldBindQuery(&pageInfo); err != nil { response.FailWithMessage(err.Error(), c); return }
    list, total, err := servicePointsCfg.GetPointsConfigPublic(ctx, pageInfo)
    if err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:"+err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{List:list, Total:total, Page:pageInfo.Page, PageSize:pageInfo.PageSize}, "获取成功", c)
}

