package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AgentTransaction 代理分润记录
type AgentTransaction struct {
	global.GVA_MODEL
	OrderNo        string  `json:"orderNo" gorm:"index;comment:订单号;column:order_no;size:64;"`
	OrderID        uint    `json:"orderId" gorm:"index;comment:订单ID;column:order_id;"`
	OrderAmount    float64 `json:"orderAmount" gorm:"comment:订单金额;column:order_amount;"`
	BaseAmount     float64 `json:"baseAmount" gorm:"comment:分润基数(87%);column:base_amount;"`
	MerchantID     uint    `json:"merchantId" gorm:"index;comment:商户ID;column:merchant_id;"`
	BeneficiaryID  uint    `json:"beneficiaryId" gorm:"index;comment:受益人ID;column:beneficiary_id;"`
	SourceUserID   uint    `json:"sourceUserId" gorm:"index;comment:来源用户ID;column:source_user_id;"`
	AgentLevelID   uint    `json:"agentLevelId" gorm:"comment:触发的代理等级ID;column:agent_level_id;"`
	AgentLevelName string  `json:"agentLevelName" gorm:"comment:触发的代理等级名称;column:agent_level_name;size:255;"`
	DividendScope  string  `json:"dividendScope" gorm:"comment:分红范围;column:dividend_scope;size:255;"`
	
	// 分润明细
	Rate1   float64 `json:"rate1" gorm:"comment:比例1;column:rate1;"`
	Amount1 float64 `json:"amount1" gorm:"comment:金额1;column:amount1;"`
	Rate2   float64 `json:"rate2" gorm:"comment:比例2;column:rate2;"`
	Amount2 float64 `json:"amount2" gorm:"comment:金额2;column:amount2;"`
	Rate3   float64 `json:"rate3" gorm:"comment:比例3;column:rate3;"`
	Amount3 float64 `json:"amount3" gorm:"comment:金额3;column:amount3;"`
	Rate4   float64 `json:"rate4" gorm:"comment:比例4;column:rate4;"`
	Amount4 float64 `json:"amount4" gorm:"comment:金额4;column:amount4;"`
	Rate5   float64 `json:"rate5" gorm:"comment:比例5;column:rate5;"`
	Amount5 float64 `json:"amount5" gorm:"comment:金额5;column:amount5;"`
	
	TotalAmount float64 `json:"totalAmount" gorm:"comment:总分润金额;column:total_amount;"`
	Description string  `json:"description" gorm:"comment:描述;column:description;size:500;"`
}

func (AgentTransaction) TableName() string {
	return "app_agent_transactions"
}
