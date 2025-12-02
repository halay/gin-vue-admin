package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUsersSearch struct {
	CreatedAtRange     []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Email              *string     `json:"email" form:"email"`
	Nickname           *string     `json:"nickname" form:"nickname"`
	Status             *string     `json:"status" form:"status"`
	LastLoginTimeRange []time.Time `json:"lastLoginTimeRange" form:"lastLoginTimeRange[]"`
	EmailVerified      *bool       `json:"emailVerified" form:"emailVerified"`
	MembershipLevelID  *int64      `json:"membershipLevelId" form:"membershipLevelId"`
	NodeID             *int64      `json:"nodeId" form:"nodeId"`
	request.PageInfo
}

// SendVerificationCodeRequest 发送验证码请求
type SendVerificationCodeRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Purpose string `json:"purpose" binding:"required"` // register, reset_password, change_email
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=6"`
	VerificationCode string `json:"verificationCode" binding:"required"`
	Nickname         string `json:"nickname"`
	Phone            string `json:"phone"`
	InviteCode       string `json:"inviteCode" binding:"required"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	LastLoginIP string `json:"lastLoginIP"`
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required"`
	NewPassword      string `json:"newPassword" binding:"required,min=6"`
}
