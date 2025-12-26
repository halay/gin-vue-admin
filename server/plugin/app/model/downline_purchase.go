package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DownlinePurchaseRecord 下级购买记录
type DownlinePurchaseRecord struct {
	global.GVA_MODEL
	OrderId              *uint    `json:"orderId" gorm:"index;comment:订单ID;column:order_id;"`
	OrderNo              *string  `json:"orderNo" gorm:"comment:订单号;column:order_no;size:64;"`
	PurchaseUserId       *uint    `json:"purchaseUserId" gorm:"index;comment:购买人ID;column:purchase_user_id;"`
	PurchaseUserNickname *string  `json:"purchaseUserNickname" gorm:"comment:购买人昵称;column:purchase_user_nickname;size:50;"`
	PurchaseUserPhone    *string  `json:"purchaseUserPhone" gorm:"comment:购买人手机号;column:purchase_user_phone;size:20;"`
	PurchaseUserEmail    *string  `json:"purchaseUserEmail" gorm:"comment:购买人邮箱;column:purchase_user_email;size:100;"`
	MerchantId           *uint    `json:"merchantId" gorm:"index;comment:商户ID;column:merchant_id;"`
	MerchantName         *string  `json:"merchantName" gorm:"comment:商户名称;column:merchant_name;size:64;"`
	UplineUserId         *uint    `json:"uplineUserId" gorm:"index;comment:上级用户ID(受益人);column:upline_user_id;"`
	UplineUserEmail      *string  `json:"uplineUserEmail" gorm:"comment:上级用户邮箱(受益人);column:upline_user_email;size:100;"`
	IsDirect             *bool    `json:"isDirect" gorm:"comment:是否直接下级;column:is_direct;"`
	InviteLevel          *int     `json:"inviteLevel" gorm:"comment:层级深度(相对);column:invite_level;"`
	TotalAmount          *float64 `json:"totalAmount" gorm:"comment:订单金额;column:total_amount;"`
	TotalPoints          *int64   `json:"totalPoints" gorm:"comment:订单积分;column:total_points;"`
	PayStatus            *string  `json:"payStatus" gorm:"comment:支付状态;column:pay_status;"`
	OrderStatus          *string  `json:"orderStatus" gorm:"comment:订单状态;column:order_status;"`
}

func (DownlinePurchaseRecord) TableName() string {
	return "app_downline_purchase_records"
}
