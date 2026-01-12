package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	appResponse "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/response"
	system2 "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	appUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var AppUsers = new(appUsers)

type appUsers struct{}

// CreateAppUsers 创建appUsers表
// @Tags AppUsers
// @Summary 创建appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "创建appUsers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /appUsers/createAppUsers [post]
func (a *appUsers) CreateAppUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.AppUsers
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAppUsers.CreateAppUsers(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAppUsers 删除appUsers表
// @Tags AppUsers
// @Summary 删除appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "删除appUsers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /appUsers/deleteAppUsers [delete]
func (a *appUsers) DeleteAppUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceAppUsers.DeleteAppUsers(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAppUsersByIds 批量删除appUsers表
// @Tags AppUsers
// @Summary 批量删除appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /appUsers/deleteAppUsersByIds [delete]
func (a *appUsers) DeleteAppUsersByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceAppUsers.DeleteAppUsersByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAppUsers 更新appUsers表
// @Tags AppUsers
// @Summary 更新appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "更新appUsers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /appUsers/updateAppUsers [put]
func (a *appUsers) UpdateAppUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info request.UpdateRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceAppUsers.UpdateAppUsers(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAppUsers 用id查询appUsers表
// @Tags AppUsers
// @Summary 用id查询appUsers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询appUsers表"
// @Success 200 {object} response.Response{data=model.AppUsers,msg=string} "查询成功"
// @Router /appUsers/findAppUsers [get]
func (a *appUsers) FindAppUsers(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reappUsers, err := serviceAppUsers.GetAppUsers(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	ancestors, _ := serviceAppUsers.FormatAncestors(ctx, reappUsers.InvitePath)
	descendants, _ := serviceAppUsers.FormatDescendants(ctx, reappUsers.ID)
	// enrich rows with invite info
	u := reappUsers
	var enriched = make(gin.H)
	enriched = gin.H{
		"ID":                u.ID,
		"CreatedAt":         u.CreatedAt,
		"email":             derefStr(u.Email),
		"nickname":          derefStr(u.Nickname),
		"avatar":            derefStr(u.Avatar),
		"phone":             derefStr(u.Phone),
		"status":            derefStr(u.Status),
		"lastLoginTime":     u.LastLoginTime,
		"lastLoginIp":       derefStr(u.LastLoginIP),
		"emailVerified":     derefBool(u.EmailVerified),
		"authorityId":       derefInt64u(u.AuthorityId),
		"inviteCode":        derefStr(u.InviteCode),
		"ancestors":         ancestors,
		"descendants":       descendants,
		"membershipLevelId": derefInt64(u.MembershipLevelID),
		"nodeId":            derefInt64(u.NodeID),
		"merchantId":        derefInt64(u.MerchantID),
	}

	response.OkWithData(enriched, c)
}

// GetAppUsersList 分页获取appUsers表列表
// @Tags AppUsers
// @Summary 分页获取appUsers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AppUsersSearch true "分页获取appUsers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /appUsers/getAppUsersList [get]
func (a *appUsers) GetAppUsersList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.AppUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := appUtils.GetUserID(c)
	mid, _ := serviceMerchantAdmin.GetMerchantIDByUserID(ctx, userID)
	if mid != nil {
		//通过mid查询app_users表，查询商家的app_users信息
		var merchantUser model.AppUsers
		if err := global.GVA_DB.Where("merchant_id=?", mid).First(&merchantUser).Error; err == nil {
			if merchantUser.ID != 0 {
				pageInfo.PathUser = strconv.Itoa(int(merchantUser.ID))
			}
		}
	}
	list, total, err := serviceAppUsers.GetAppUsersInfoList(ctx, pageInfo)
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

// GetAppUsersPublic 不需要鉴权的appUsers表接口
// @Tags AppUsers
// @Summary 不需要鉴权的appUsers表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appUsers/getAppUsersPublic [get]
func (a *appUsers) GetAppUsersPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceAppUsers.GetAppUsersPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的appUsers表接口信息"}, "获取成功", c)
}

// SendVerificationCode 发送验证码
// @Tags     AppUser
// @Summary  发送验证码
// @Accept   application/json
// @Produce  application/json
// @Param    data body request.SendVerificationCodeRequest true "邮箱和验证码用途"
// @Success  200  {object} response.Response{msg=string} "发送成功"
// @Router   /appUsers/sendVerificationCode [post]
func (a *appUsers) SendVerificationCode(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var req request.SendVerificationCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := serviceAppUsers.ChargeEmailExits(ctx, req.Email, req.Purpose)
	if err != nil {
		global.GVA_LOG.Error("邮箱是否存在判断!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 调用服务层发送验证码
	err = serviceAppUsers.SendVerificationCode(ctx, req.Email, req.Purpose)
	if err != nil {
		global.GVA_LOG.Error("发送验证码失败!", zap.Error(err))
		response.FailWithMessage("发送验证码失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("验证码已发送", c)
}

// Register 用户注册
// @Tags     AppUser
// @Summary  用户注册
// @accept   application/json
// @Produce  application/json
// @Param    data body request.RegisterRequest true "注册信息"
// @Success  200  {object} response.Response{msg=string} "注册成功"
// @Router   /appUsers/register [post]
func (a *appUsers) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user, err := serviceAppUsers.Register(c, req)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage("注册失败: "+err.Error(), c)
		return
	}
	token, claims, err := appUtils.LoginToken(user)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	appUtils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	// 积分账户余额 (兼容旧逻辑: 获取平台积分)
	accs, _ := serviceUserPointsAccount.GetUserAllPointsAccounts(c.Request.Context(), int64(user.ID))
	var balance int64
	var pointsAccounts []appResponse.PointsAccount
	for _, a := range accs {
		bal := int64(0)
		if a.Balance != nil {
			bal = *a.Balance
		}
		// 转换到 Response 模型
		pa := appResponse.PointsAccount{
			MerchantID: a.MerchantID,
			TokenName:  a.TokenName,
			Symbol:     a.Symbol,
			Balance:    bal,
			IsPlatform: a.MerchantID == nil || *a.MerchantID == 0,
		}
		pointsAccounts = append(pointsAccounts, pa)

		// 平台积分赋值给 PointsBalance 以兼容旧逻辑
		if pa.IsPlatform {
			balance = bal
		}
	}

	response.OkWithDetailed(appResponse.AppLoginResponse{
		User:           user,
		Token:          token,
		ExpiresAt:      claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		PointsBalance:  balance,
		PointsAccounts: pointsAccounts,
	}, "注册成功", c)
	//response.OkWithData(user, c)
}

// Login 用户登录
// @Tags     AppUser
// @Summary  用户登录
// @accept   application/json
// @Produce  application/json
// @Param    data body request.LoginRequest true "登录信息"
// @Success  200  {object} response.AppLoginResponse{data=response.AppLoginResponse,msg=string} "登录成功"
// @Router   /appUsers/login [post]
func (a *appUsers) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.LastLoginIP = c.ClientIP()
	user, err := serviceAppUsers.Login(c, req)
	if err != nil {
		global.GVA_LOG.Error("登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败: "+err.Error(), c)
		return
	}
	token, claims, err := appUtils.LoginToken(user)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	appUtils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	var merchant model.Merchants
	if user.MerchantID != nil {
		_ = global.GVA_DB.Where("id = ?", *user.MerchantID).First(&merchant).Error
		user.Merchant = merchant
	} else {
		if user.InvitePath != nil && *user.InvitePath != "" {
			parts := strings.Split(*user.InvitePath, "/")
			// 无论是否有"/"，parts至少包含一个元素（本身）
			// 从parts中最后一个开始查找merchantUser，直到找到了商家信息才停止遍历
			for i := len(parts) - 1; i >= 0; i-- {
				if parts[i] == "" {
					continue
				}
				num, err := strconv.Atoi(parts[i])
				if err != nil {
					continue
				}
				// 排除自己（虽然InvitePath通常不包含自己，但为了保险）
				if num == int(user.ID) {
					continue
				}

				var merchantUser model.AppUsers
				if err := global.GVA_DB.Where("id=?", num).First(&merchantUser).Error; err == nil {
					if merchantUser.MerchantID != nil {
						_ = global.GVA_DB.Where("id = ?", merchantUser.MerchantID).First(&merchant).Error
						user.Merchant = merchant
						break // 找到商户信息后停止遍历
					}
				}
			}
		}
	}

	// 积分账户余额 (兼容旧逻辑: 获取平台积分)
	accs, _ := serviceUserPointsAccount.GetUserAllPointsAccounts(c.Request.Context(), int64(user.ID))
	var balance int64
	var pointsAccounts []appResponse.PointsAccount
	for _, a := range accs {
		bal := int64(0)
		if a.Balance != nil {
			bal = *a.Balance
		}
		// 转换到 Response 模型
		pa := appResponse.PointsAccount{
			MerchantID: a.MerchantID,
			TokenName:  a.TokenName,
			Symbol:     a.Symbol,
			Balance:    bal,
			IsPlatform: a.MerchantID == nil || *a.MerchantID == 0,
		}
		pointsAccounts = append(pointsAccounts, pa)

		// 平台积分赋值给 PointsBalance 以兼容旧逻辑
		if pa.IsPlatform {
			balance = bal
		}
	}

	var shareholderProfitId int64
	var shareholderProfitName string
	if user.ShareholderProfitID != nil {
		shareholderProfitId = *user.ShareholderProfitID
		var sp model.ShareholderProfit
		if global.GVA_DB.Where("id = ?", shareholderProfitId).First(&sp).Error == nil && sp.Name != nil {
			shareholderProfitName = *sp.Name
		}
	}

	var appDealerId int64
	var appDealerName string
	if user.AppDealerID != nil {
		appDealerId = *user.AppDealerID
		var ad model.AppDealer
		if global.GVA_DB.Where("id = ?", appDealerId).First(&ad).Error == nil && ad.Name != nil {
			appDealerName = *ad.Name
		}
	}

	response.OkWithDetailed(appResponse.AppLoginResponse{
		User:           user,
		Token:          token,
		ExpiresAt:      claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		PointsBalance:  balance,
		PointsAccounts: pointsAccounts,
		//Merchant:  model.Merchants `json:"merchant"`
		ShareholderProfitID:   shareholderProfitId,
		ShareholderProfitName: shareholderProfitName,
		AppDealerID:           appDealerId,
		AppDealerName:         appDealerName,
	}, "登录成功", c)
}

// ResetPassword 重置密码
// @Tags     AppUser
// @Summary  重置密码
// @accept   application/json
// @Produce  application/json
// @Param    data body request.ResetPasswordRequest true "重置密码信息"
// @Success  200  {object} response.Response{msg=string} "重置成功"
// @Router   /appUsers/resetPassword [post]
func (a *appUsers) ResetPassword(c *gin.Context) {
	var req request.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := serviceAppUsers.ResetPassword(c, req); err != nil {
		global.GVA_LOG.Error("重置密码失败!", zap.Error(err))
		response.FailWithMessage("重置密码失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("重置成功", c)
}

// ResetPassword 退出登录
// @Tags     AppUser
// @Summary  退出登录
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200  {object} response.Response{msg=string} "退出成功"
// @Router   /appUsers/logout [post]
func (a *appUsers) Logout(c *gin.Context) {
	token := appUtils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := system2.JwtServiceApp.JsonInBlacklist(jwt)
	if err != nil {
		global.GVA_LOG.Error("退出登录,jwt作废失败!", zap.Error(err))
		response.FailWithMessage("退出登录失败", c)
		return
	}
	appUtils.ClearToken(c)
	response.OkWithMessage("退出成功", c)

}

// GetInviteCount 获取用户的邀请统计
// @Tags AppUser
// @Summary 获取用户的邀请统计（累计和今日）
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]int64,msg=string} "获取成功"
// @Router /appUsers/getInviteCount [get]
func (a *appUsers) GetInviteCount(c *gin.Context) {
	ctx := c.Request.Context()
	userID := appUtils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("无法从token提取用户标识", c)
		return
	}

	total, today, err := serviceAppUsers.GetInviteCount(ctx, userID)
	if err != nil {
		global.GVA_LOG.Error("获取邀请统计失败!", zap.Error(err))
		response.FailWithMessage("获取邀请统计失败:"+err.Error(), c)
		return
	}

	response.OkWithData(gin.H{
		"totalInvite": total,
		"todayInvite": today,
	}, c)
}

// GetSubUsers 获取下级用户列表
// @Tags AppUser
// @Summary 获取下级用户列表
// @Security ApiKeyAuth
// @Produce application/json
// @Param type query string true "查询类型: direct(直属), indirect(间接)"
// @Success 200 {object} response.Response{data=[]model.AppUsers,msg=string} "获取成功"
// @Router /appUsers/getSubUsers [get]
func (a *appUsers) GetSubUsers(c *gin.Context) {
	userID := appUtils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("无法从token提取用户标识", c)
		return
	}

	queryType := c.Query("type")
	if queryType != "direct" && queryType != "indirect" {
		response.FailWithMessage("参数错误，type必须为direct或indirect", c)
		return
	}

	list, err := serviceAppUsers.GetSubUsers(c.Request.Context(), userID, queryType)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	response.OkWithData(list, c)
}

// GetUserInfo 通过x-token获取当前用户信息
// @Tags     AppUser
// @Summary  获取当前登录用户信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200 {object} response.Response{data=response.UserResponse,msg=string} "获取成功"
// @Router   /appUsers/getUserInfo [get]
func (a *appUsers) GetUserInfo(c *gin.Context) {
	var user model.AppUsers
	userId := appUtils.GetUserID(c)
	if userId == 0 {
		response.FailWithMessage("无法从token提取用户标识", c)
		return
	}

	// 根据id查询
	if err := global.GVA_DB.Where("id = ?", userId).First(&user).Error; err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	// 组装返回体，避免泄露敏感字段（如密码）
	resp := appResponse.UserResponse{
		ID:                user.ID,
		Email:             derefStr(user.Email),
		Nickname:          derefStr(user.Nickname),
		Avatar:            derefStr(user.Avatar),
		Phone:             derefStr(user.Phone),
		Status:            derefStr(user.Status),
		EmailVerified:     derefBool(user.EmailVerified),
		LastLoginTime:     user.LastLoginTime,
		LastLoginIP:       derefStr(user.LastLoginIP),
		InviteCode:        derefStr(user.InviteCode),
		InviterID:         uint(derefInt64u(user.InviterID)),
		InviteLevel:       derefInt(user.InviteLevel),
		MembershipLevelID: derefInt64(user.MembershipLevelID),
		NodeID:            derefInt64(user.NodeID),
		MerchantID:        derefInt64(user.MerchantID),
	}
	// 积分余额 (兼容旧逻辑: 获取平台积分)
	accs, _ := serviceUserPointsAccount.GetUserAllPointsAccounts(c.Request.Context(), int64(user.ID))
	var pointsAccounts []appResponse.PointsAccount
	for _, a := range accs {
		bal := int64(0)
		if a.Balance != nil {
			bal = *a.Balance
		}
		// 转换到 Response 模型
		pa := appResponse.PointsAccount{
			MerchantID: a.MerchantID,
			TokenName:  a.TokenName,
			Symbol:     a.Symbol,
			Balance:    bal,
			IsPlatform: a.MerchantID == nil || *a.MerchantID == 0,
		}
		pointsAccounts = append(pointsAccounts, pa)

		// 平台积分赋值给 PointsBalance 以兼容旧逻辑
		if pa.IsPlatform {
			resp.PointsBalance = bal
		}
	}
	resp.PointsAccounts = pointsAccounts

	var merchant model.Merchants
	if user.MerchantID != nil {
		_ = global.GVA_DB.Where("id = ?", *user.MerchantID).First(&merchant).Error
		resp.Merchant = merchant
	} else {
		if user.InvitePath != nil && *user.InvitePath != "" {
			parts := strings.Split(*user.InvitePath, "/")
			// 无论是否有"/"，parts至少包含一个元素（本身）
			// 从parts中最后一个开始查找merchantUser，直到找到了商家信息才停止遍历
			for i := len(parts) - 1; i >= 0; i-- {
				if parts[i] == "" {
					continue
				}
				num, err := strconv.Atoi(parts[i])
				if err != nil {
					continue
				}
				// 排除自己（虽然InvitePath通常不包含自己，但为了保险）
				if num == int(user.ID) {
					continue
				}

				var merchantUser model.AppUsers
				if err := global.GVA_DB.Where("id=?", num).First(&merchantUser).Error; err == nil {
					if merchantUser.MerchantID != nil {
						_ = global.GVA_DB.Where("id = ?", merchantUser.MerchantID).First(&merchant).Error
						resp.Merchant = merchant
						break // 找到商户信息后停止遍历
					}
				}
			}
		}
	}

	var shareholderProfitId int64
	var shareholderProfitName string
	if user.ShareholderProfitID != nil {
		shareholderProfitId = *user.ShareholderProfitID
		// 尝试获取股东名称
		var sp model.ShareholderProfit
		if global.GVA_DB.Where("id = ?", shareholderProfitId).First(&sp).Error == nil && sp.Name != nil {
			shareholderProfitName = *sp.Name
		}
	}
	resp.ShareholderProfitID = shareholderProfitId
	resp.ShareholderProfitName = shareholderProfitName

	var appDealerId int64
	var appDealerName string
	if user.AppDealerID != nil {
		appDealerId = *user.AppDealerID
		if user.AppDealer.Name != nil {
			appDealerName = *user.AppDealer.Name
		} else {
			// 如果预加载未成功，尝试单独查询
			var ad model.AppDealer
			if global.GVA_DB.Where("id = ?", appDealerId).First(&ad).Error == nil && ad.Name != nil {
				appDealerName = *ad.Name
			}
		}
	}
	resp.AppDealerID = appDealerId
	resp.AppDealerName = appDealerName

	response.OkWithDetailed(resp, "获取成功", c)
}

// GetDashboard 获取APP用户首页/个人中心聚合数据
// @Tags     AppUser
// @Summary  获取APP用户首页/个人中心聚合数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200 {object} response.Response{data=response.DashboardResponse,msg=string} "获取成功"
// @Router   /appUsers/getDashboard [get]
func (a *appUsers) GetDashboard(c *gin.Context) {
	userId := appUtils.GetUserID(c)
	if userId == 0 {
		response.FailWithMessage("无法从token提取用户标识", c)
		return
	}
	mid := c.DefaultQuery("mid", "0")
	if mid != "" {
		merchant, err := serviceMerchants.GetMerchants(c.Request.Context(), mid)
		if err != nil {
			global.GVA_LOG.Error("获取商户信息失败!", zap.Error(err))
			response.FailWithMessage("获取商户信息失败: "+err.Error(), c)
			return
		}
		if merchant.ID == 0 {
			response.FailWithMessage("商户不存在", c)
			return
		}
	}
	merchantId, _ := strconv.ParseInt(mid, 10, 64)
	dashboard, err := serviceAppUsers.GetDashboard(c.Request.Context(), userId, merchantId)
	if err != nil {
		global.GVA_LOG.Error("获取Dashboard失败!", zap.Error(err))
		response.FailWithMessage("获取Dashboard失败: "+err.Error(), c)
		return
	}

	response.OkWithData(dashboard, c)
}

func derefStr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func derefInt64(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}

func derefInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

func derefInt64u(p *uint) int64 {
	if p == nil {
		return 0
	}
	return int64(*p)
}

func derefBool(p *bool) bool {
	if p == nil {
		return false
	}
	return *p
}
