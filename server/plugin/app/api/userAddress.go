package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	appUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var UserAddress = new(UA)

type UA struct{}

// CreateUserAddress 创建用户收货地址
// @Tags UserAddress
// @Summary 创建用户收货地址
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddress true "创建用户收货地址"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /UA/createUserAddress [post]
func (a *UA) CreateUserAddress(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录或无权限", c)
		return
	}
	var info model.UserAddress
	u := int64(uid)
	info.UserID = &u
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceUserAddress.CreateUserAddress(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserAddress 删除用户收货地址
// @Tags UserAddress
// @Summary 删除用户收货地址
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddress true "删除用户收货地址"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /UA/deleteUserAddress [delete]
func (a *UA) DeleteUserAddress(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录或无权限", c)
		return
	}
	// 所属校验
	var addr model.UserAddress
	u := int64(uid)
	addr.UserID = &u
	if e := global.GVA_DB.Where("id = ?", ID).First(&addr).Error; e != nil {
		if e == gorm.ErrRecordNotFound {
			response.FailWithMessage("记录不存在", c)
			return
		}
		response.FailWithMessage(e.Error(), c)
		return
	}
	if addr.UserID == nil || *addr.UserID != int64(uid) {
		response.FailWithMessage("无权删除他人地址", c)
		return
	}
	err := serviceUserAddress.DeleteUserAddress(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserAddressByIds 批量删除用户收货地址
// @Tags UserAddress
// @Summary 批量删除用户收货地址
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /UA/deleteUserAddressByIds [delete]
func (a *UA) DeleteUserAddressByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录或无权限", c)
		return
	}
	// 只允许删除属于自己的记录
	var cnt int64
	_ = global.GVA_DB.Table("app_user_addresses").Where("id in ? AND user_id = ?", IDs, uid).Count(&cnt).Error
	if cnt != int64(len(IDs)) {
		response.FailWithMessage("仅能批量删除属于自己的地址", c)
		return
	}
	err := serviceUserAddress.DeleteUserAddressByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserAddress 更新用户收货地址
// @Tags UserAddress
// @Summary 更新用户收货地址
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserAddress true "更新用户收货地址"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /UA/updateUserAddress [put]
func (a *UA) UpdateUserAddress(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录或无权限", c)
		return
	}
	var info model.UserAddress
	userId := int64(uid)
	info.UserID = &userId
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验归属
	var orig model.UserAddress
	if e := global.GVA_DB.Where("id = ?", info.ID).First(&orig).Error; e != nil {
		if e == gorm.ErrRecordNotFound {
			response.FailWithMessage("记录不存在", c)
			return
		}
		response.FailWithMessage(e.Error(), c)
		return
	}
	if orig.UserID == nil || *orig.UserID != int64(uid) {
		response.FailWithMessage("无权修改他人地址", c)
		return
	}
	u := int64(uid)
	info.UserID = &u
	err = serviceUserAddress.UpdateUserAddress(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserAddress 用id查询用户收货地址
// @Tags UserAddress
// @Summary 用id查询用户收货地址
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户收货地址"
// @Success 200 {object} response.Response{data=model.UserAddress,msg=string} "查询成功"
// @Router /UA/findUserAddress [get]
func (a *UA) FindUserAddress(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录或无权限", c)
		return
	}
	reUA, err := serviceUserAddress.GetUserAddress(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	if reUA.UserID == nil || *reUA.UserID != int64(uid) {
		response.FailWithMessage("无权查看他人地址", c)
		return
	}
	response.OkWithData(reUA, c)
}

// GetUserAddressList 分页获取用户收货地址列表
// @Tags UserAddress
// @Summary 分页获取用户收货地址列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.UserAddressSearch true "分页获取用户收货地址列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /UA/getUserAddressList [get]
func (a *UA) GetUserAddressList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.UserAddressSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := appUtils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录或无权限", c)
		return
	}
	// 限定为当前用户
	v := int(uid)
	pageInfo.UserID = &v
	list, total, err := serviceUserAddress.GetUserAddressInfoList(ctx, pageInfo)
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

// GetUserAddressPublic 不需要鉴权的用户收货地址接口
// @Tags UserAddress
// @Summary 不需要鉴权的用户收货地址接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /UA/getUserAddressPublic [get]
func (a *UA) GetUserAddressPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceUserAddress.GetUserAddressPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的用户收货地址接口信息"}, "获取成功", c)
}
