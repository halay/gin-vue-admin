
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProductCategory 商户商品分类 结构体
type ProductCategory struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:分类名称;column:name;size:255;" binding:"required"`  //分类名称
  Status  *string `json:"status" form:"status" gorm:"default:enabled;comment:启用/禁用;column:status;" binding:"required"`  //分类状态
  Sort  *int64 `json:"sort" form:"sort" gorm:"comment:排序值;column:sort;"`  //排序值
  Description  *string `json:"description" form:"description" gorm:"comment:分类说明;column:description;type:text;"`  //分类说明
  MerchantID  *int64 `json:"merchantId" form:"merchantId" gorm:"comment:关联商户;column:merchant_id;"`  //关联商户
}


// TableName 商户商品分类 ProductCategory自定义表名 app_product_categories
func (ProductCategory) TableName() string {
    return "app_product_categories"
}







