package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

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
	err = global.GVA_DB.Create(appUsers).Error
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
func (s *appUsers) UpdateAppUsers(ctx context.Context, appUsers model.AppUsers) (err error) {
	err = global.GVA_DB.Model(&model.AppUsers{}).Where("id = ?", appUsers.ID).Updates(&appUsers).Error
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
		req.Email, req.VerificationCode, "register", false, time.Now()).First(&verification).Error; err != nil {
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

	// 创建用户
	trueVal := true
	password := string(hashedPassword)
	nickname := req.Nickname
	phone := req.Phone
	status := "1"
	roleID := plugin.Config.AuthorityId // 外部用户角色ID
	user = model.AppUsers{
		Email:         &req.Email,
		Uuid:          uuid.New(),
		Password:      &password,
		Nickname:      &nickname,
		Phone:         &phone,
		EmailVerified: &trueVal,
		Status:        &status,
		AuthorityId:   &roleID, // 关联角色ID=123
	}

	if err = global.GVA_DB.Create(&user).Error; err != nil {
		return user, err
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
