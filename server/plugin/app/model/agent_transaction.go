
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// AgentTransactionDetail 代理交易明细 结构体
type AgentTransactionDetail struct {
    global.GVA_MODEL
  SequenceNumber  *int64 `json:"sequenceNumber" form:"sequenceNumber" gorm:"comment:代理交易序号;column:sequence_number;"`  //序号
  Name  *string `json:"name" form:"name" gorm:"comment:代理人姓名;column:name;size:50;"`  //姓名
  AccountId  *string `json:"accountId" form:"accountId" gorm:"comment:代理账号ID;column:account_id;size:100;"`  //账号ID
  Email  *string `json:"email" form:"email" gorm:"comment:代理人邮箱;column:email;size:100;"`  //邮箱
  AgentLevel  string `json:"agentLevel" form:"agentLevel" gorm:"comment:代理人级别;column:agent_level;type:enum('一级','二级','三级');"`  //代理级别
  TransactionTime  *time.Time `json:"transactionTime" form:"transactionTime" gorm:"comment:交易发生时间;column:transaction_time;"`  //交易时间
  Product  *string `json:"product" form:"product" gorm:"comment:购买的商品名称;column:product;size:200;"`  //购买商品
  SubsidyAmount  *float64 `json:"subsidyAmount" form:"subsidyAmount" gorm:"comment:补贴金额;column:subsidy_amount;size:10,2;"`  //补贴金额
  Status  string `json:"status" form:"status" gorm:"comment:交易状态;column:status;type:enum('待处理','已确认','已取消');"`  //状态
}


// TableName 代理交易明细 AgentTransactionDetail自定义表名 agent_transaction_detail
func (AgentTransactionDetail) TableName() string {
    return "agent_transaction_detail"
}







