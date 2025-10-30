package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppEmailVerification 邮箱验证码 结构体
type AppEmailVerification struct {
	global.GVA_MODEL
	Email            *string    `json:"email" form:"email" gorm:"comment:邮箱地址;column:email;size:255;" binding:"required"`                                 //邮箱
	VerificationCode *string    `json:"verificationCode" form:"verificationCode" gorm:"comment:验证码;column:verification_code;size:10;" binding:"required"` //验证码
	Purpose          *string    `json:"purpose" form:"purpose" gorm:"comment:验证码用途;column:purpose;size:20;" binding:"required"`                           //用途
	ExpireTime       *time.Time `json:"expireTime" form:"expireTime" gorm:"comment:验证码过期时间;column:expire_time;" binding:"required"`                       //过期时间
	Used             *bool      `json:"used" form:"used" gorm:"default:false;comment:验证码是否已使用;column:used;"`                                              //是否已使用
}

// TableName 邮箱验证码 AppEmailVerification自定义表名 app_email_verifications
func (AppEmailVerification) TableName() string {
	return "app_email_verifications"
}
