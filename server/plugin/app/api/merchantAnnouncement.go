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

var MerchantAnnouncement = new(MA)

type MA struct {}

// CreateMerchantAnnouncement 创建商户公告信息
// @Tags MerchantAnnouncement
// @Summary 创建商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAnnouncement true "创建商户公告信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /MA/createMerchantAnnouncement [post]
func (a *MA) CreateMerchantAnnouncement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var info model.MerchantAnnouncement
    err := c.ShouldBindJSON(&info)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法创建公告", c)
        return
    }
    err = serviceMerchantAnnouncement.CreateMerchantAnnouncement(ctx, &info, *mid)
    if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
        response.FailWithMessage("创建失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("创建成功", c)
}

// DeleteMerchantAnnouncement 删除商户公告信息
// @Tags MerchantAnnouncement
// @Summary 删除商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAnnouncement true "删除商户公告信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /MA/deleteMerchantAnnouncement [delete]
func (a *MA) DeleteMerchantAnnouncement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    ID := c.Query("ID")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法删除公告", c)
        return
    }
    err := serviceMerchantAnnouncement.DeleteMerchantAnnouncement(ctx, ID, *mid)
    if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
        response.FailWithMessage("删除失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("删除成功", c)
}

// DeleteMerchantAnnouncementByIds 批量删除商户公告信息
// @Tags MerchantAnnouncement
// @Summary 批量删除商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /MA/deleteMerchantAnnouncementByIds [delete]
func (a *MA) DeleteMerchantAnnouncementByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法批量删除公告", c)
        return
    }
    err := serviceMerchantAnnouncement.DeleteMerchantAnnouncementByIds(ctx, IDs, *mid)
    if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
        response.FailWithMessage("批量删除失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("批量删除成功", c)
}

// UpdateMerchantAnnouncement 更新商户公告信息
// @Tags MerchantAnnouncement
// @Summary 更新商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAnnouncement true "更新商户公告信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /MA/updateMerchantAnnouncement [put]
func (a *MA) UpdateMerchantAnnouncement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var info model.MerchantAnnouncement
    err := c.ShouldBindJSON(&info)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法更新公告", c)
        return
    }
    err = serviceMerchantAnnouncement.UpdateMerchantAnnouncement(ctx, info, *mid)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
        response.FailWithMessage("更新失败:" + err.Error(), c)
        return
    }
    response.OkWithMessage("更新成功", c)
}

// FindMerchantAnnouncement 用id查询商户公告信息
// @Tags MerchantAnnouncement
// @Summary 用id查询商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询商户公告信息"
// @Success 200 {object} response.Response{data=model.MerchantAnnouncement,msg=string} "查询成功"
// @Router /MA/findMerchantAnnouncement [get]
func (a *MA) FindMerchantAnnouncement(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    ID := c.Query("ID")
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法查询公告", c)
        return
    }
    reMA, err := serviceMerchantAnnouncement.GetMerchantAnnouncement(ctx, ID, *mid)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
        return
    }
    response.OkWithData(reMA, c)
}
// GetMerchantAnnouncementList 分页获取商户公告信息列表
// @Tags MerchantAnnouncement
// @Summary 分页获取商户公告信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantAnnouncementSearch true "分页获取商户公告信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /MA/getMerchantAnnouncementList [get]
func (a *MA) GetMerchantAnnouncementList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    var pageInfo request.MerchantAnnouncementSearch
    err := c.ShouldBindQuery(&pageInfo)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    userID := utils.GetUserID(c)
    mid, errMid := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
    if errMid != nil || mid == nil {
        response.FailWithMessage("未绑定商户，无法获取公告列表", c)
        return
    }
    list, total, err := serviceMerchantAnnouncement.GetMerchantAnnouncementInfoList(ctx, pageInfo, *mid)
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
// GetMerchantAnnouncementDataSource 获取MerchantAnnouncement的数据源
// @Tags MerchantAnnouncement
// @Summary 获取MerchantAnnouncement的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /MA/getMerchantAnnouncementDataSource [get]
func (a *MA) GetMerchantAnnouncementDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
   dataSource, err := serviceMerchantAnnouncement.GetMerchantAnnouncementDataSource(ctx)
   if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
		return
   }
    response.OkWithData(dataSource, c)
}
// GetMerchantAnnouncementPublic 不需要鉴权的商户公告信息接口
// @Tags MerchantAnnouncement
// @Summary 不需要鉴权的商户公告信息接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MA/getMerchantAnnouncementPublic [get]
func (a *MA) GetMerchantAnnouncementPublic(c *gin.Context) {
    ctx := c.Request.Context()
    var pageInfo request.MerchantAnnouncementSearch
    if err := c.ShouldBindQuery(&pageInfo); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    if pageInfo.MerchantID == nil {
        response.FailWithMessage("merchantId为必填", c)
        return
    }
    list, total, err := serviceMerchantAnnouncement.GetMerchantAnnouncementInfoList(ctx, pageInfo, *pageInfo.MerchantID)
    if err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:"+err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{List:list, Total:total, Page:pageInfo.Page, PageSize:pageInfo.PageSize}, "获取成功", c)
}
