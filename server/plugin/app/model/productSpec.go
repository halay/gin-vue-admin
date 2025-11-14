
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProductSpec 商品规格键 结构体
type ProductSpec struct {
    global.GVA_MODEL
  ProductID  *int64 `json:"productId" form:"productId" gorm:"comment:商品ID;column:product_id;" binding:"required"`
  Name  *string `json:"name" form:"name" gorm:"comment:规格名;column:name;size:64;" binding:"required"`
  Sort  *int64 `json:"sort" form:"sort" gorm:"comment:排序;column:sort;"`
  MerchantID  *int64 `json:"merchantId" form:"merchantId" gorm:"comment:关联商户;column:merchant_id;index:idx_merchant_id"`
}


// TableName 商品规格键 ProductSpec自定义表名 app_product_specs
func (ProductSpec) TableName() string {
    return "app_product_specs"
}







