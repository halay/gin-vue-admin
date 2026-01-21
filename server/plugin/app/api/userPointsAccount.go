package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	appUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var UserPointsAccount = new(UPA)

type UPA struct{}

// CreateUserPointsAccount 创建用户积分账户
// @Tags UserPointsAccount
// @Summary 创建用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsAccount true "创建用户积分账户"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /UPA/createUserPointsAccount [post]
func (a *UPA) CreateUserPointsAccount(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.UserPointsAccount
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceUserPointsAccount.CreateUserPointsAccount(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserPointsAccount 删除用户积分账户
// @Tags UserPointsAccount
// @Summary 删除用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsAccount true "删除用户积分账户"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /UPA/deleteUserPointsAccount [delete]
func (a *UPA) DeleteUserPointsAccount(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceUserPointsAccount.DeleteUserPointsAccount(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserPointsAccountByIds 批量删除用户积分账户
// @Tags UserPointsAccount
// @Summary 批量删除用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /UPA/deleteUserPointsAccountByIds [delete]
func (a *UPA) DeleteUserPointsAccountByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceUserPointsAccount.DeleteUserPointsAccountByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserPointsAccount 更新用户积分账户
// @Tags UserPointsAccount
// @Summary 更新用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsAccount true "更新用户积分账户"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /UPA/updateUserPointsAccount [put]
func (a *UPA) UpdateUserPointsAccount(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.UserPointsAccount
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceUserPointsAccount.UpdateUserPointsAccount(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserPointsAccount 用id查询用户积分账户
// @Tags UserPointsAccount
// @Summary 用id查询用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询用户积分账户"
// @Success 200 {object} response.Response{data=model.UserPointsAccount,msg=string} "查询成功"
// @Router /UPA/findUserPointsAccount [get]
func (a *UPA) FindUserPointsAccount(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reUPA, err := serviceUserPointsAccount.GetUserPointsAccount(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reUPA, c)
}

// GetUserPointsAccountList 分页获取用户积分账户列表
// @Tags UserPointsAccount
// @Summary 分页获取用户积分账户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.UserPointsAccountSearch true "分页获取用户积分账户列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /UPA/getUserPointsAccountList [get]
func (a *UPA) GetUserPointsAccountList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.UserPointsAccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := appUtils.GetUserID(c)
	mid, _ := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	pageInfo.MerchantID = mid
	list, total, err := serviceUserPointsAccount.GetUserPointsAccountList(ctx, pageInfo)
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

// GetUserPointsAccountDataSource 获取UserPointsAccount的数据源
// @Tags UserPointsAccount
// @Summary 获取UserPointsAccount的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /UPA/getUserPointsAccountDataSource [get]
func (a *UPA) GetUserPointsAccountDataSource(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口为获取数据源定义的数据
	dataSource, err := serviceUserPointsAccount.GetUserPointsAccountDataSource(ctx)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetUserPointsAccountPublic 不需要鉴权的用户积分账户接口
// @Tags UserPointsAccount
// @Summary 不需要鉴权的用户积分账户接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /UPA/getUserPointsAccountPublic [get]
func (a *UPA) GetUserPointsAccountPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceUserPointsAccount.GetUserPointsAccountPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的用户积分账户接口信息"}, "获取成功", c)
}

// GetMyPointsBalance 获取当前登录用户的积分余额（包括平台和商户）
// @Tags UserPointsAccount
// @Summary 获取当前登录用户积分余额（包括平台和商户）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /UPA/getMyPointsBalance [get]
func (a *UPA) GetMyPointsBalance(c *gin.Context) {
	ctx := c.Request.Context()
	uid := utils.GetUserID(c)
	if uid == 0 {
		response.FailWithMessage("未登录", c)
		return
	}
	// 获取用户所有积分账户
	accounts, err := serviceUserPointsAccount.GetUserAllPointsAccounts(ctx, int64(uid))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	// 构造返回数据
	var res []map[string]interface{}
	for _, acc := range accounts {
		item := map[string]interface{}{
			"merchantId": acc.MerchantID,
			"tokenName":  acc.TokenName,
			"symbol":     acc.Symbol,
			"balance":    0,
		}
		if acc.Balance != nil {
			item["balance"] = *acc.Balance
		}
		// 平台积分特殊标记
		if acc.MerchantID != nil && *acc.MerchantID == 0 {
			item["isPlatform"] = true
		} else {
			item["isPlatform"] = false
		}
		res = append(res, item)
	}

	response.OkWithData(gin.H{"accounts": res}, c)
}
