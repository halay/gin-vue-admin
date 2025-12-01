
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MembershipLevel 会员等级 结构体
type MembershipLevel struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:普通/VIP/SVIP;column:name;size:64;" binding:"required"`  //等级名称
  Code  *string `json:"code" form:"code" gorm:"comment:normal/vip/svip;column:code;size:32;" binding:"required"`  //等级编码
  Status  *string `json:"status" form:"status" gorm:"default:1;comment:启用/禁用;column:status;"`  //状态
  Sort  *int64 `json:"sort" form:"sort" gorm:"comment:排序;column:sort;"`  //排序
  Requirement  *string `json:"requirement" form:"requirement" gorm:"comment:如100美元/年;column:requirement;size:255;"`  //要求描述
  RequirementValue  *float64 `json:"requirementValue" form:"requirementValue" gorm:"comment:数值要求;column:requirement_value;"`  //要求值
  Benefits  *string `json:"benefits" form:"benefits" gorm:"comment:权益列表说明;column:benefits;type:text;"`  //权益说明
}


// TableName 会员等级 MembershipLevel自定义表名 app_membership_levels
func (MembershipLevel) TableName() string {
    return "app_membership_levels"
}







