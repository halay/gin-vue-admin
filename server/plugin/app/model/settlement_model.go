
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Settlement 结算管理 结构体
type Settlement struct {
    global.GVA_MODEL
  SerialNumber  *int64 `json:"serialNumber" form:"serialNumber" gorm:"comment:序号;column:serial_number;"`  //序号
  Name  *string `json:"name" form:"name" gorm:"comment:姓名;column:name;size:50;" binding:"required"`  //姓名
  AccountID  *string `json:"accountID" form:"accountID" gorm:"comment:账号ID;column:account_id;size:64;" binding:"required"`  //账号ID
  SettlementPeriod  *string `json:"settlementPeriod" form:"settlementPeriod" gorm:"comment:结算周期;column:settlement_period;size:20;" binding:"required"`  //结算周期
  SubsidyAmount  *float64 `json:"subsidyAmount" form:"subsidyAmount" gorm:"comment:补贴金额;column:subsidy_amount;size:(10,2);" binding:"required"`  //补贴金额
  Status  string `json:"status" form:"status" gorm:"comment:结算状态;column:status;type:enum('待结算','已结算','已驳回');" binding:"required"`  //结算状态
}


// TableName 结算管理 Settlement自定义表名 settlement
func (Settlement) TableName() string {
    return "settlement"
}







