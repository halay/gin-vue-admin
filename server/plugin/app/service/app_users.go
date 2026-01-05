package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/plugin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/utils"
)

var AppUsers = new(appUsers)

type appUsers struct{}

// CreateAppUsers 创建appUsers表记录
// Author [yourname](https://github.com/yourname)
func (s *appUsers) CreateAppUsers(ctx context.Context, appUsers *model.AppUsers) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if appUsers.MerchantID != nil {
			var cnt int64
			if e := tx.Model(&model.AppUsers{}).Where("merchant_id = ?", *appUsers.MerchantID).Count(&cnt).Error; e != nil {
				return e
			}
			if cnt > 0 {
				return errors.New("该商户已绑定其他用户")
			}
		}
		return tx.Create(appUsers).Error
	})
	return err
}

// DeleteAppUsers 删除appUsers表记录
// Author [yourname](https://github.com/yourname)
func (s *appUsers) DeleteAppUsers(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.AppUsers{}, "id = ?", ID).Error
	return err
}

// DeleteAppUsersByIds 批量删除appUsers表记录
// Author [yourname](https://github.com/yourname)
func (s *appUsers) DeleteAppUsersByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.AppUsers{}, "id in ?", IDs).Error
	return err
}

// UpdateAppUsers 更新appUsers表记录
// Author [yourname](https://github.com/yourname)
func (s *appUsers) UpdateAppUsers(ctx context.Context, appUsers request.UpdateRequest) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if appUsers.MerchantID != nil {
			var cnt int64
			if e := tx.Model(&model.AppUsers{}).Where("merchant_id = ? AND id <> ?", *appUsers.MerchantID, appUsers.ID).Count(&cnt).Error; e != nil {
				return e
			}
			if cnt > 0 {
				return errors.New("该商户已绑定其他用户")
			}
		}
		return tx.Model(&model.AppUsers{}).Where("id = ?", appUsers.ID).Updates(&appUsers).Error
	})
	return err
}

// GetAppUsers 根据ID获取appUsers表记录
// Author [yourname](https://github.com/yourname)
func (s *appUsers) GetAppUsers(ctx context.Context, ID string) (appUsers model.AppUsers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&appUsers).Error
	return
}

// GetAppUsersInfoList 分页获取appUsers表记录
// Author [yourname](https://github.com/yourname)
func (s *appUsers) GetAppUsersInfoList(ctx context.Context, info request.AppUsersSearch) (list []model.AppUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AppUsers{})
	var appUserss []model.AppUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Email != nil && *info.Email != "" {
		db = db.Where("email LIKE ?", "%"+*info.Email+"%")
	}
	if info.Nickname != nil && *info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+*info.Nickname+"%")
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if len(info.LastLoginTimeRange) == 2 {
		db = db.Where("last_login_time BETWEEN ? AND ? ", info.LastLoginTimeRange[0], info.LastLoginTimeRange[1])
	}
	if info.EmailVerified != nil {
		db = db.Where("email_verified = ?", *info.EmailVerified)
	}
	if info.PathUser != "" {
		db = db.Where("invite_path = ? OR invite_path LIKE ? OR invite_path LIKE ? OR invite_path LIKE ?",
			info.PathUser,
			info.PathUser+"/%",
			"%/"+info.PathUser,
			"%/"+info.PathUser+"/%")
	}
	if info.MembershipLevelID != nil {
		db = db.Where("membership_level_id = ?", *info.MembershipLevelID)
	}
	if info.NodeID != nil {
		db = db.Where("node_id = ?", *info.NodeID)
	}
	if info.MerchantID != nil {
		db = db.Where("merchant_id = ?", *info.MerchantID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&appUserss).Error
	return appUserss, total, err
}

func (s *appUsers) GetAppUsersPublic(ctx context.Context) {

}

// ChargeEmailExits 检测邮箱是否存在
func (s *appUsers) ChargeEmailExits(ctx context.Context, email string, purpose string) error {
	if purpose == "register" {
		// 检查邮箱是否已存在
		var count int64
		if err := global.GVA_DB.Model(&model.AppUsers{}).Where("email = ?", email).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return errors.New("邮箱已存在")
		}
	} else {
		// 检查邮箱是否已存在
		var count int64
		if err := global.GVA_DB.Model(&model.AppUsers{}).Where("email = ?", email).Count(&count).Error; err != nil {
			return err
		}
		if count == 0 {
			return errors.New("邮箱不存在")
		}
	}
	return nil
}

// SendVerificationCode 发送验证码
func (s *appUsers) SendVerificationCode(ctx context.Context, email string, purpose string) error {
	var code, ginMode string
	ginMode = gin.Mode()
	if ginMode == gin.DebugMode {
		code = "123456"
	} else {
		code = generateRandomCode(6)
	}
	// 生成6位随机验证码
	expireTime := time.Now().Add(15 * time.Minute)
	// 创建验证码记录
	verification := model.AppEmailVerification{
		Email:            &email,
		VerificationCode: &code,
		Purpose:          &purpose,
		ExpireTime:       &expireTime, // 15分钟有效期
		Used:             new(bool),   // 初始化为false
	}
	*verification.Used = false

	if ginMode == gin.ReleaseMode {
		// 发送邮件
		subject := "验证码"
		switch purpose {
		case "register":
			subject = "注册验证码"
		case "reset_password":
			subject = "重置密码验证码"
		case "change_email":
			subject = "修改邮箱验证码"
		}

		body := fmt.Sprintf("您的验证码是: %s，有效期15分钟，请勿泄露给他人。", code)
		if err := utils.Email(email, subject, body); err != nil {
			return err
		}
	}

	// 保存验证码记录
	if err := global.GVA_DB.Create(&verification).Error; err != nil {
		return err
	}
	return nil
}

// Register 用户注册
func (s *appUsers) Register(ctx context.Context, req request.RegisterRequest) (user model.AppUsers, err error) {
	// 检查邮箱是否已注册
	var count int64
	if err = global.GVA_DB.Model(&model.AppUsers{}).Where("email = ?", req.Email).Count(&count).Error; err != nil {
		return user, err
	}
	if count > 0 {
		return user, errors.New("该邮箱已注册")
	}

	// 验证验证码
	var verification model.AppEmailVerification
	if err = global.GVA_DB.Where("email = ? AND verification_code = ? AND purpose = ? AND used = ? AND expire_time > ?",
		req.Email, req.VerificationCode, "register", 0, time.Now()).First(&verification).Error; err != nil {
		return user, errors.New("验证码无效或已过期")
	}

	// 标记验证码为已使用
	trueValue := true
	verification.Used = &trueValue
	if err = global.GVA_DB.Save(&verification).Error; err != nil {
		return user, err
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	// 校验邀请码
	codeUpper := strings.ToUpper(strings.TrimSpace(req.InviteCode))
	if len(codeUpper) != 6 {
		return user, errors.New("邀请码必须为6位")
	}
	var inviter model.AppUsers
	if err = global.GVA_DB.Where("invite_code = ?", codeUpper).First(&inviter).Error; err != nil {
		return user, errors.New("邀请码不存在")
	}

	// 生成自己的唯一邀请码（大小写不敏感，统一大写保存）
	myCode, genErr := s.generateUniqueInviteCode(ctx)
	if genErr != nil {
		return user, genErr
	}

	// 计算邀请路径与层级
	var path string
	if inviter.InvitePath != nil && *inviter.InvitePath != "" {
		path = *inviter.InvitePath + "/" + fmt.Sprintf("%d", inviter.ID)
	} else {
		path = fmt.Sprintf("%d", inviter.ID)
	}
	if path == "12" {
		path = ""
	}
	level := 1
	if inviter.InviteLevel != nil {
		level = *inviter.InviteLevel + 1
	}

	// 创建用户
	trueVal := true
	password := string(hashedPassword)
	nickname := req.Nickname
	phone := req.Phone
	status := "1"
	roleID := plugin.Config.AuthorityId // 外部用户角色ID
	memberShipInfo, _ := MembershipLevel.GetMembershipLevelInfoByCode(ctx, "normal")
	var membershipId = int64(memberShipInfo.ID)
	user = model.AppUsers{
		Email:             &req.Email,
		Uuid:              uuid.New(),
		Password:          &password,
		Nickname:          &nickname,
		Phone:             &phone,
		EmailVerified:     &trueVal,
		Status:            &status,
		AuthorityId:       &roleID, // 关联角色ID=123
		InviteCode:        &myCode,
		InviterID:         &inviter.ID,
		InvitePath:        &path,
		InviteLevel:       &level,
		MembershipLevelID: &membershipId,
	}

	if err = global.GVA_DB.Create(&user).Error; err != nil {
		return user, err
	}

	// 积分奖励：注册、邀请人、被邀请人
	// 获取启用的积分设置（取最新一条）
	type pts struct {
		RegisterReward *int64
		InviterReward  *int64
		InviteeReward  *int64
	}
	var cfg pts
	_ = global.GVA_DB.Table("app_points_settings").
		Where("status = ?", 1).
		Select("register_reward,inviter_reward,invitee_reward").
		Order("sort asc, id desc").
		Limit(1).Scan(&cfg).Error
	// 奖励注册用户
	if cfg.RegisterReward != nil && *cfg.RegisterReward > 0 {
		after, _ := UserPointsAccount.AddPoints(ctx, int64(user.ID), *cfg.RegisterReward)
		_ = UserPointsAccount.AddLogDetailed(ctx, int64(user.ID), *cfg.RegisterReward, after, "register_reward", "", "注册奖励", "reward", "success", nil, nil)
	}
	// 奖励邀请人
	if user.InviterID != nil && cfg.InviterReward != nil && *cfg.InviterReward > 0 {
		invID := int64(*user.InviterID)
		after, _ := UserPointsAccount.AddPoints(ctx, invID, *cfg.InviterReward)
		rid := int64(user.ID)
		_ = UserPointsAccount.AddLogDetailed(ctx, invID, *cfg.InviterReward, after, "inviter_reward", "", "邀请人奖励", "reward", "success", &rid, nil)
	}
	// 奖励被邀请人
	if cfg.InviteeReward != nil && *cfg.InviteeReward > 0 {
		after, _ := UserPointsAccount.AddPoints(ctx, int64(user.ID), *cfg.InviteeReward)
		invID := int64(0)
		if user.InviterID != nil {
			invID = int64(*user.InviterID)
		}
		_ = UserPointsAccount.AddLogDetailed(ctx, int64(user.ID), *cfg.InviteeReward, after, "invitee_reward", "", "被邀请人奖励", "reward", "success", &invID, nil)
	}
	// 商户积分配置发放：遍历所有启用的商户积分设置
	type mptsRow struct {
		MerchantID     int64
		InviterReward  *int64
		InviteeReward  *int64
		RegisterReward *int64
		CoverageMode   *string
	}
	var mrows []mptsRow
	_ = global.GVA_DB.Table("app_merchant_points_settings").
		Where("status IN (?)", []string{"enabled", "1"}).
		Select("merchant_id,inviter_reward,invitee_reward,register_reward,coverage_mode").
		Order("sort asc, id desc").Scan(&mrows).Error
	for _, ms := range mrows {
		mode := "directed"
		if ms.CoverageMode != nil {
			mode = *ms.CoverageMode
		}
		eligible := false
		if mode == "platform" {
			eligible = true
		} else { // directed：邀请人必须属于该商户
			eligible = s.isInviterUnderMerchant(ctx, inviter, ms.MerchantID)
		}
		if !eligible {
			continue
		}
		mid64 := ms.MerchantID
		// 注册奖励（商户）
		if ms.RegisterReward != nil && *ms.RegisterReward > 0 {
			after, _ := UserPointsAccount.AddPoints(ctx, int64(user.ID), *ms.RegisterReward)
			_ = UserPointsAccount.AddLogDetailed(ctx, int64(user.ID), *ms.RegisterReward, after, "merchant_register_reward", "", "注册奖励(商户)", "reward", "success", &mid64, &mid64)
		}
		// 邀请人奖励（商户）
		if user.InviterID != nil && ms.InviterReward != nil && *ms.InviterReward > 0 {
			invID := int64(*user.InviterID)
			after, _ := UserPointsAccount.AddPoints(ctx, invID, *ms.InviterReward)
			_ = UserPointsAccount.AddLogDetailed(ctx, invID, *ms.InviterReward, after, "merchant_inviter_reward", "", "邀请人奖励(商户)", "reward", "success", &mid64, &mid64)
		}
		// 被邀请人奖励（商户）
		if ms.InviteeReward != nil && *ms.InviteeReward > 0 {
			after, _ := UserPointsAccount.AddPoints(ctx, int64(user.ID), *ms.InviteeReward)
			invID := int64(0)
			if user.InviterID != nil {
				invID = int64(*user.InviterID)
			}
			_ = UserPointsAccount.AddLogDetailed(ctx, int64(user.ID), *ms.InviteeReward, after, "merchant_invitee_reward", "", "被邀请人奖励(商户)", "reward", "success", &invID, &mid64)
		}
	}
	return user, nil
}

// Login 用户登录
func (s *appUsers) Login(ctx context.Context, req request.LoginRequest) (user model.AppUsers, err error) {
	// 查找用户
	if err = global.GVA_DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return user, errors.New("用户不存在")
	}

	// 检查用户状态
	if *user.Status != "1" {
		return user, errors.New("账户已被禁用")
	}

	// 验证密码
	if err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(req.Password)); err != nil {
		return user, errors.New("密码错误")
	}

	// 更新登录信息
	now := time.Now()
	user.LastLoginTime = &now
	user.LastLoginIP = &req.LastLoginIP
	if err = global.GVA_DB.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// ResetPassword 重置密码
func (s *appUsers) ResetPassword(ctx context.Context, req request.ResetPasswordRequest) error {
	// 验证验证码
	var verification model.AppEmailVerification
	if err := global.GVA_DB.Where("email = ? AND verification_code = ? AND purpose = ? AND used = ? AND expire_time > ?",
		req.Email, req.VerificationCode, "reset_password", false, time.Now()).First(&verification).Error; err != nil {
		return errors.New("验证码无效或已过期")
	}

	// 标记验证码为已使用
	trueVal := true
	verification.Used = &trueVal
	if err := global.GVA_DB.Save(&verification).Error; err != nil {
		return err
	}

	// 查找用户
	var user model.AppUsers
	if err := global.GVA_DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return errors.New("用户不存在")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	passwordStr := string(hashedPassword)
	user.Password = &passwordStr
	if err := global.GVA_DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

// 生成随机验证码
func generateRandomCode(length int) string {
	const charset = "0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// 生成唯一邀请码（6位字母数字，统一大写）
func (s *appUsers) generateUniqueInviteCode(ctx context.Context) (code string, err error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 50; i++ {
		b := make([]byte, 6)
		for j := 0; j < 6; j++ {
			b[j] = charset[seededRand.Intn(len(charset))]
		}
		code = string(b)
		var cnt int64
		if e := global.GVA_DB.WithContext(ctx).Model(&model.AppUsers{}).Where("invite_code = ?", code).Count(&cnt).Error; e != nil {
			return "", e
		}
		if cnt == 0 {
			return code, nil
		}
	}
	return "", errors.New("邀请码生成失败，请重试")
}

func uintPtr(u uint) *uint { return &u }

// FormatAncestors 根据邀请路径返回祖先的邮箱(ID)列表
func (s *appUsers) FormatAncestors(ctx context.Context, path *string) ([]string, error) {
	if path == nil || *path == "" {
		return []string{}, nil
	}
	parts := strings.Split(*path, "/")
	ids := make([]uint, 0, len(parts))
	for _, p := range parts {
		if p == "" {
			continue
		}
		if v, err := strconv.ParseUint(p, 10, 64); err == nil {
			ids = append(ids, uint(v))
		}
	}
	if len(ids) == 0 {
		return []string{}, nil
	}
	// 查询邮箱
	type row struct {
		ID    uint
		Email *string
	}
	var rs []row
	if err := global.GVA_DB.WithContext(ctx).Table("app_users").
		Where("id in ?", ids).
		Select("id,email").
		Scan(&rs).Error; err != nil {
		return nil, err
	}
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		mail := ""
		if r.Email != nil {
			mail = *r.Email
		}
		out = append(out, fmt.Sprintf("%s(%d)", mail, r.ID))
	}
	return out, nil
}

// FormatDescendants 查询所有后代（邀请链中包含当前用户ID），返回邮箱(ID)列表
func (s *appUsers) FormatDescendants(ctx context.Context, uid uint) ([]string, error) {
	idStr := fmt.Sprintf("%d", uid)
	type row struct {
		ID    uint
		Email *string
	}
	var rs []row
	q := global.GVA_DB.WithContext(ctx).Table("app_users").
		Where("invite_path = ?", idStr).
		Or("invite_path LIKE ?", idStr+"/%").
		Or("invite_path LIKE ?", "%/"+idStr).
		Or("invite_path LIKE ?", "%/"+idStr+"/%").
		Select("id,email")
	if err := q.Scan(&rs).Error; err != nil {
		return nil, err
	}
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		mail := ""
		if r.Email != nil {
			mail = *r.Email
		}
		out = append(out, fmt.Sprintf("%s(%d)", mail, r.ID))
	}
	return out, nil
}

// isInviterUnderMerchant 判断邀请人是否是指定商户任意级别的下级用户：
// 原理：查询该商户绑定的用户ID集合，检测这些ID是否出现在邀请人的 invite_path 字段中
func (s *appUsers) isInviterUnderMerchant(ctx context.Context, inviter model.AppUsers, merchantID int64) bool {
	if inviter.InvitePath == nil || *inviter.InvitePath == "" {
		return false
	}
	// 查询商户对应的用户ID集合
	type row struct{ ID uint }
	var rs []row
	if err := global.GVA_DB.WithContext(ctx).
		Table("app_users").
		Where("merchant_id = ?", merchantID).
		Select("id").
		Scan(&rs).Error; err != nil {
		return false
	}
	if len(rs) == 0 {
		return false
	}
	path := *inviter.InvitePath
	for _, r := range rs {
		idStr := fmt.Sprintf("%d", r.ID)
		if path == idStr ||
			strings.HasPrefix(path, idStr+"/") ||
			strings.Contains(path, "/"+idStr+"/") ||
			strings.HasSuffix(path, "/"+idStr) {
			return true
		}
	}
	return false
}
