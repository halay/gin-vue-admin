package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MerchantAdmin 后台用户与商户关联映射 结构体
type MerchantAdmin struct {
	global.GVA_MODEL
	UserID     *int64 `json:"userId" form:"userId" gorm:"comment:后台用户ID;column:user_id;" binding:"required"`           //后台用户ID
	MerchantID *int   `json:"merchantId" form:"merchantId" gorm:"comment:商户ID;column:merchant_id;" binding:"required"` //商户ID
}

// TableName 后台用户与商户关联映射 MerchantAdmin自定义表名 app_merchant_admins
func (MerchantAdmin) TableName() string {
	return "app_merchant_admins"
}
