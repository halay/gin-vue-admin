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
	MerchantID         *int64      `json:"merchantId" form:"merchantId"`
	PathUser           string      `json:"pathUser" form:"pathUser"`
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
type UpdateRequest struct {
	ID                uint    `gorm:"primarykey" json:"ID"` // 主键ID
	MembershipLevelID *int64  `json:"membershipLevelId" form:"membershipLevelId" gorm:"comment:会员等级ID;column:membership_level_id;"`
	NodeID            *int64  `json:"nodeId" form:"nodeId" gorm:"comment:所属节点ID;column:node_id;"`
	MerchantID        *int64  `json:"merchantId" form:"merchantId" gorm:"comment:绑定商户ID;column:merchant_id;"`
	Nickname          *string `json:"nickname" form:"nickname" gorm:"comment:用户昵称;column:nickname;size:50;"` //用户昵称 //用户头像URL
	Phone             *string `json:"phone" form:"phone" gorm:"comment:用户手机号;column:phone;size:20;"`
}
