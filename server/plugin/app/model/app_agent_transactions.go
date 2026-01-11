
package model
import (
	"time"
)

// AppAgentTransactions appAgentTransactions表 结构体
type AppAgentTransactions struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  OrderNo  *string `json:"orderNo" form:"orderNo" gorm:"comment:订单号;column:order_no;size:64;"`  //订单号
  OrderId  *int64 `json:"orderId" form:"orderId" gorm:"comment:订单ID;column:order_id;"`  //订单ID
  OrderAmount  *float64 `json:"orderAmount" form:"orderAmount" gorm:"comment:订单金额;column:order_amount;size:22;"`  //订单金额
  BaseAmount  *float64 `json:"baseAmount" form:"baseAmount" gorm:"comment:分润基数(87%);column:base_amount;size:22;"`  //分润基数(87%)
  MerchantId  *int64 `json:"merchantId" form:"merchantId" gorm:"comment:商户ID;column:merchant_id;"`  //商户ID
  BeneficiaryId  *int64 `json:"beneficiaryId" form:"beneficiaryId" gorm:"comment:受益人ID;column:beneficiary_id;"`  //受益人ID
  SourceUserId  *int64 `json:"sourceUserId" form:"sourceUserId" gorm:"comment:来源用户ID;column:source_user_id;"`  //来源用户ID
  AgentLevelId  *int64 `json:"agentLevelId" form:"agentLevelId" gorm:"comment:触发的代理等级ID;column:agent_level_id;"`  //触发的代理等级ID
  AgentLevelName  *string `json:"agentLevelName" form:"agentLevelName" gorm:"comment:触发的代理等级名称;column:agent_level_name;size:255;"`  //触发的代理等级名称
  DividendScope  *string `json:"dividendScope" form:"dividendScope" gorm:"comment:分红范围;column:dividend_scope;size:255;"`  //分红范围
  Rate1  *float64 `json:"rate1" form:"rate1" gorm:"comment:比例1;column:rate1;size:22;"`  //比例1
  Amount1  *float64 `json:"amount1" form:"amount1" gorm:"comment:金额1;column:amount1;size:22;"`  //金额1
  Rate2  *float64 `json:"rate2" form:"rate2" gorm:"comment:比例2;column:rate2;size:22;"`  //比例2
  Amount2  *float64 `json:"amount2" form:"amount2" gorm:"comment:金额2;column:amount2;size:22;"`  //金额2
  Rate3  *float64 `json:"rate3" form:"rate3" gorm:"comment:比例3;column:rate3;size:22;"`  //比例3
  Amount3  *float64 `json:"amount3" form:"amount3" gorm:"comment:金额3;column:amount3;size:22;"`  //金额3
  Rate4  *float64 `json:"rate4" form:"rate4" gorm:"comment:比例4;column:rate4;size:22;"`  //比例4
  Amount4  *float64 `json:"amount4" form:"amount4" gorm:"comment:金额4;column:amount4;size:22;"`  //金额4
  Rate5  *float64 `json:"rate5" form:"rate5" gorm:"comment:比例5;column:rate5;size:22;"`  //比例5
  Amount5  *float64 `json:"amount5" form:"amount5" gorm:"comment:金额5;column:amount5;size:22;"`  //金额5
  TotalAmount  *float64 `json:"totalAmount" form:"totalAmount" gorm:"comment:总分润金额;column:total_amount;size:22;"`  //总分润金额
  Description  *string `json:"description" form:"description" gorm:"comment:描述;column:description;size:500;"`  //描述
  SourceId  *int64 `json:"sourceId" form:"sourceId" gorm:"comment:来源id;column:source_id;"`  //来源id
  Source  *string `json:"source" form:"source" gorm:"comment:来源;column:source;size:500;"`  //来源
  ReflectStatus  *string `json:"reflectStatus" form:"reflectStatus" gorm:"comment:提现状态;column:reflect_status;size:50;"`  //提现状态
}


// TableName appAgentTransactions表 AppAgentTransactions自定义表名 app_agent_transactions
func (AppAgentTransactions) TableName() string {
    return "app_agent_transactions"
}







