package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserPointsAccount 用户积分账户 结构体
type UserPointsAccount struct {
	global.GVA_MODEL
	UserID  *int64 `json:"userId" form:"userId" gorm:"comment:app用户ID;column:user_id;" binding:"required"`          //用户
	Balance    *int64  `json:"balance" form:"balance" gorm:"default:0;comment:积分余额;column:balance;" binding:"required"` //积分余额
	MerchantID *int64  `json:"merchantId" form:"merchantId" gorm:"index;comment:商户ID(0表示平台);column:merchant_id;"`       //商户ID
	TokenName  *string `json:"tokenName" form:"tokenName" gorm:"comment:积分名称;column:token_name;size:64;"`               //积分名称
	Symbol     *string `json:"symbol" form:"symbol" gorm:"comment:积分符号;column:symbol;size:32;"`                         //积分符号
}

// TableName 用户积分账户 UserPointsAccount自定义表名 app_user_points_accounts
func (UserPointsAccount) TableName() string {
	return "app_user_points_accounts"
}
