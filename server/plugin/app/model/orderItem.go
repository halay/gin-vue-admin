package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// OrderItem 订单明细 结构体
type OrderItem struct {
	global.GVA_MODEL
	OrderID     *int64         `json:"orderId" form:"orderId" gorm:"comment:订单ID;column:order_id;" binding:"required"`                       //订单
	ProductID   *int64         `json:"productId" form:"productId" gorm:"comment:商品ID;column:product_id;" binding:"required"`                 //商品
	SkuID       *int64         `json:"skuId" form:"skuId" gorm:"comment:SKU ID;column:sku_id;" binding:"required"`                           //SKU
	ProductName *string        `json:"productName" form:"productName" gorm:"comment:商品名快照;column:product_name;size:255;" binding:"required"` //商品名快照
	SkuAttrs    datatypes.JSON `json:"skuAttrs" form:"skuAttrs" gorm:"comment:SKU属性快照;column:sku_attrs;" swaggertype:"object"`               //SKU属性
	Price       *float64       `json:"price" form:"price" gorm:"comment:单价;column:price;" binding:"required"`                                //单价
	Points      *int64         `json:"points" form:"points" gorm:"comment:单件积分;column:points;"`                                              //单件积分
	Quantity    *int64         `json:"quantity" form:"quantity" gorm:"comment:数量;column:quantity;" binding:"required"`                       //数量
	TotalAmount *float64       `json:"totalAmount" form:"totalAmount" gorm:"comment:金额小计;column:total_amount;"`                              //小计金额
	TotalPoints *int64         `json:"totalPoints" form:"totalPoints" gorm:"comment:积分小计;column:total_points;"`                              //小计积分
	CoverImage  string         `json:"coverImage" form:"coverImage" gorm:"comment:图片快照;column:cover_image;"`                                 //图片
	MerchantID  *int64         `json:"merchantId" form:"merchantId" gorm:"comment:商户ID;column:merchant_id;"`                                 //商户
}

// TableName 订单明细 OrderItem自定义表名 app_order_items
func (OrderItem) TableName() string {
	return "app_order_items"
}
