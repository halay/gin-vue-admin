
package model
import (
)

// AgentLevel 营销方案设置 结构体
type AgentLevel struct {
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;"`  //id字段
  MerchantId  *int64 `json:"merchantId" form:"merchantId" gorm:"comment:品牌方id;column:merchant_id;"`  //品牌方id
  Name  *string `json:"name" form:"name" gorm:"comment:名称;column:name;size:255;"`  //名称
  QualificationAmount  *int32 `json:"qualificationAmount" form:"qualificationAmount" gorm:"comment:代理资格总金额(万);column:qualification_amount;"`  //代理资格(万)
  DividendScope  *string `json:"dividendScope" form:"dividendScope" gorm:"comment:分红范围;column:dividend_scope;size:255;"`  //分红范围
  CfRate1  *float64 `json:"cfRate1" form:"cfRate1" gorm:"comment:比例1;column:cf_rate1;size:22;"`  //业务提成
  CfRate2  *float64 `json:"cfRate2" form:"cfRate2" gorm:"comment:比例2;column:cf_rate2;size:22;"`  //销售费用补贴
  CfRate3  *float64 `json:"cfRate3" form:"cfRate3" gorm:"comment:比例3;column:cf_rate3;size:22;"`  //经销补贴
  CfRate4  *float64 `json:"cfRate4" form:"cfRate4" gorm:"comment:比例4;column:cf_rate4;size:22;"`  //分红比例
  CfRate5  *float64 `json:"cfRate5" form:"cfRate5" gorm:"comment:比例4;column:cf_rate5;size:22;"`  //额外补贴
  Status  *string `json:"status" form:"status" gorm:"index;comment:是否启用;column:status;"`  //状态
}


// TableName 营销方案设置 AgentLevel自定义表名 app_agent_level
func (AgentLevel) TableName() string {
    return "app_agent_level"
}







