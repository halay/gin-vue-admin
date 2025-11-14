
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// ProductSku 商品SKU（规格组合） 结构体
type ProductSku struct {
    global.GVA_MODEL
  ProductID  *int64 `json:"productId" form:"productId" gorm:"comment:商品ID;column:product_id;uniqueIndex:uniq_product_sku,priority:1" binding:"required"`
  SKUCode  *string `json:"skuCode" form:"skuCode" gorm:"comment:SKU编码;column:sku_code;size:64;uniqueIndex:uniq_product_sku,priority:2"`
  Attrs  datatypes.JSON `json:"attrs" form:"attrs" gorm:"comment:属性组合JSON;column:attrs;" swaggertype:"object"`
  Price  *float64 `json:"price" form:"price" gorm:"comment:SKU价格;column:price;" binding:"required"`
  Points  *int64 `json:"points" form:"points" gorm:"comment:SKU积分;column:points;"`
  Image  string `json:"image" form:"image" gorm:"comment:SKU图片;column:image;"`
  Stock  *int64 `json:"stock" form:"stock" gorm:"comment:SKU库存;column:stock;" binding:"required"`
  Status  *string `json:"status" form:"status" gorm:"default:enabled;comment:启用/禁用;column:status;" binding:"required"`
  MerchantID  *int64 `json:"merchantId" form:"merchantId" gorm:"comment:关联商户;column:merchant_id;index:idx_merchant_id"`
}


// TableName 商品SKU（规格组合） ProductSku自定义表名 app_product_skus
func (ProductSku) TableName() string {
    return "app_product_skus"
}







