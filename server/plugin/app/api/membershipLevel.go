package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var MembershipLevel = new(ML)

type ML struct {}

// CreateMembershipLevel 创建会员等级
// @Tags MembershipLevel
// @Summary 创建会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MembershipLevel true "创建会员等级"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ML/createMembershipLevel [post]
func (a *ML) CreateMembershipLevel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MembershipLevel
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMembershipLevel.CreateMembershipLevel(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMembershipLevel 删除会员等级
// @Tags MembershipLevel
// @Summary 删除会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MembershipLevel true "删除会员等级"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /ML/deleteMembershipLevel [delete]
func (a *ML) DeleteMembershipLevel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceMembershipLevel.DeleteMembershipLevel(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteMembershipLevelByIds 批量删除会员等级
// @Tags MembershipLevel
// @Summary 批量删除会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /ML/deleteMembershipLevelByIds [delete]
func (a *ML) DeleteMembershipLevelByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceMembershipLevel.DeleteMembershipLevelByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateMembershipLevel 更新会员等级
// @Tags MembershipLevel
// @Summary 更新会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MembershipLevel true "更新会员等级"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /ML/updateMembershipLevel [put]
func (a *ML) UpdateMembershipLevel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.MembershipLevel
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceMembershipLevel.UpdateMembershipLevel(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindMembershipLevel 用id查询会员等级
// @Tags MembershipLevel
// @Summary 用id查询会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询会员等级"
// @Success 200 {object} response.Response{data=model.MembershipLevel,msg=string} "查询成功"
// @Router /ML/findMembershipLevel [get]
func (a *ML) FindMembershipLevel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reML, err := serviceMembershipLevel.GetMembershipLevel(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reML, c)
}
// GetMembershipLevelList 分页获取会员等级列表
// @Tags MembershipLevel
// @Summary 分页获取会员等级列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.MembershipLevelSearch true "分页获取会员等级列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /ML/getMembershipLevelList [get]
func (a *ML) GetMembershipLevelList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.MembershipLevelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceMembershipLevel.GetMembershipLevelInfoList(ctx,pageInfo)
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
// GetMembershipLevelPublic 不需要鉴权的会员等级接口
// @Tags MembershipLevel
// @Summary 不需要鉴权的会员等级接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ML/getMembershipLevelPublic [get]
func (a *ML) GetMembershipLevelPublic(c *gin.Context) {
    ctx := c.Request.Context()
    list, err := serviceMembershipLevel.ListEnabled(ctx)
    if err != nil { response.FailWithMessage("获取失败:"+err.Error(), c); return }
    response.OkWithData(list, c)
}
