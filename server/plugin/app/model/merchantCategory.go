
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MerchantCategory 商户分类 结构体
type MerchantCategory struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:分类名称;column:name;size:64;" binding:"required"`  //分类名称
  Status  *string `json:"status" form:"status" gorm:"default:1;comment:启用/禁用;column:status;"`  //状态
  Sort  *int64 `json:"sort" form:"sort" gorm:"comment:排序;column:sort;"`  //排序
}


// TableName 商户分类 MerchantCategory自定义表名 app_merchant_categories
func (MerchantCategory) TableName() string {
    return "app_merchant_categories"
}







