package model

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// ProductSkuOption SKU-规格值中间表
type ProductSkuOption struct {
    global.GVA_MODEL
    SkuID        *int64 `json:"skuId" form:"skuId" gorm:"comment:SKU ID;column:sku_id;index:idx_sku"`
    SpecID       *int64 `json:"specId" form:"specId" gorm:"comment:规格键ID;column:spec_id;index:idx_spec"`
    OptionID     *int64 `json:"optionId" form:"optionId" gorm:"comment:规格值ID;column:option_id;index:idx_option"`
    MerchantID   *int64 `json:"merchantId" form:"merchantId" gorm:"comment:商户ID;column:merchant_id;index:idx_merchant_id"`
}

func (ProductSkuOption) TableName() string { return "app_product_sku_options" }
