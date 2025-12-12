package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type PointsRechargeOrder struct {
	global.GVA_MODEL
	OrderNo            *string    `json:"orderNo" form:"orderNo" gorm:"comment:订单号;column:order_no;uniqueIndex;size:64;"`
	UserID             *int64     `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;"`
	ConfigID           *uint      `json:"configId" form:"configId" gorm:"comment:积分配置ID;column:config_id;"`
	Points             *int64     `json:"points" form:"points" gorm:"comment:购买积分;column:points;"`
	BonusPoints        *int64     `json:"bonusPoints" form:"bonusPoints" gorm:"comment:赠送积分;column:bonus_points;"`
	Amount             *float64   `json:"amount" form:"amount" gorm:"comment:支付金额;column:amount;"`
	PayMethod          *string    `json:"payMethod" form:"payMethod" gorm:"comment:支付方式;column:pay_method;size:32;"`
	Status             *string    `json:"status" form:"status" gorm:"comment:状态(pending/success/failed);column:status;size:32;default:pending;"`
	TransactionID      *string    `json:"transactionId" form:"transactionId" gorm:"comment:第三方交易号;column:transaction_id;size:128;"`
	PaidAt             *time.Time `json:"paidAt" form:"paidAt" gorm:"comment:支付时间;column:paid_at;"`
	Remark             *string    `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:255;"`
	StripeClientSecret *string    `json:"stripeClientSecret" gorm:"comment:Stripe客户端密钥;column:stripe_client_secret;size:255;"`
}

func (PointsRechargeOrder) TableName() string { return "app_points_recharge_orders" }
