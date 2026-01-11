package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppDealer 经销商管理 结构体
type AppDealer struct {
	global.GVA_MODEL
	Name             *string  `json:"name" form:"name" gorm:"comment:经销商名称;column:name;size:64;" binding:"required"`                                              //经销商名称
	SalesCommission  *float64 `json:"salesCommission" form:"salesCommission" gorm:"comment:销售提成;column:sales_commission;size:64;"`                                //销售提成
	CommissionType   *int     `json:"commissionType" form:"commissionType" gorm:"default:1;comment:提成类型(1:比例,2:固定金额);column:commission_type;" binding:"required"` //提成类型
	ExpenseAllowance *float64 `json:"expenseAllowance" form:"expenseAllowance" gorm:"comment:费用补贴;column:expense_allowance;size:64;"`                             //费用补贴
	AllowanceType    *int     `json:"allowanceType" form:"allowanceType" gorm:"default:1;comment:补贴类型(1:比例,2:固定金额);column:allowance_type;" binding:"required"`    //补贴类型
	MerchantId       *int64   `json:"merchantId" form:"merchantId" gorm:"comment:关联商户;column:merchant_id;"`                                                       //关联商户
	Remark           *string  `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:255;"`                                                             //备注
	IsDefault        *bool    `json:"isDefault" form:"isDefault" gorm:"default:false;comment:是否默认;column:is_default;"`                                            //是否默认
}

// TableName 经销商管理 AppDealer自定义表名 app_dealers
func (AppDealer) TableName() string {
	return "app_dealers"
}
