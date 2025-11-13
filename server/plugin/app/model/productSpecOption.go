
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProductSpecOption 商品规格值 结构体
type ProductSpecOption struct {
    global.GVA_MODEL
  ProductID  *int64 `json:"productId" form:"productId" gorm:"comment:商品ID;column:product_id;" binding:"required"`  //商品
  SpecID  *int64 `json:"specId" form:"specId" gorm:"comment:规格键ID;column:spec_id;" binding:"required"`  //规格键
  Value  *string `json:"value" form:"value" gorm:"comment:规格值;column:value;size:64;" binding:"required"`  //规格值
  Sort  *int64 `json:"sort" form:"sort" gorm:"comment:排序;column:sort;"`  //排序
  MerchantID  *int64 `json:"merchantId" form:"merchantId" gorm:"comment:关联商户;column:merchant_id;"`  //关联商户
}


// TableName 商品规格值 ProductSpecOption自定义表名 app_product_spec_options
func (ProductSpecOption) TableName() string {
    return "app_product_spec_options"
}







