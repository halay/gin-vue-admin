package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// Product 商户商品 结构体
type Product struct {
	global.GVA_MODEL
	Name           *string        `json:"name" form:"name" gorm:"comment:商品名称;column:name;size:255;" binding:"required"`
	Type           *string        `json:"type" form:"type" gorm:"default:physical;comment:虚拟/实体;column:type;" binding:"required"`
	Status         *string        `json:"status" form:"status" gorm:"default:enabled;comment:草稿/上架/下线;column:status;" binding:"required"`
	CategoryID     *int64         `json:"categoryId" form:"categoryId" gorm:"comment:所属分类;column:category_id;" binding:"required"`
	CoverImage     string         `json:"coverImage" form:"coverImage" gorm:"comment:封面图;column:cover_image;"`
	CarouselImages datatypes.JSON `json:"carouselImages" form:"carouselImages" gorm:"comment:轮播图;column:carousel_images;" swaggertype:"array,object"`
	Video          string         `json:"video" form:"video" gorm:"comment:视频;column:video;"`
	Detail         *string        `json:"detail" form:"detail" gorm:"comment:详情;column:detail;type:text;"`
	HasVariants    *bool          `json:"hasVariants" form:"hasVariants" gorm:"default:false;comment:是否多规格;column:has_variants;"`
	EnablePoints   *bool          `json:"enablePoints" form:"enablePoints" gorm:"default:false;comment:开启积分支付;column:enable_points;"`
	Price          *float64       `json:"price" form:"price" gorm:"comment:价格(无规格时);column:price;"`
	Points         *int64         `json:"points" form:"points" gorm:"comment:积分(无规格时);column:points;"`
	Stock          *int64         `json:"stock" form:"stock" gorm:"comment:库存(无规格时);column:stock;"`
	MerchantID     *int64         `json:"merchantId" form:"merchantId" gorm:"comment:关联商户;column:merchant_id;index:idx_merchant_id"`
}

// TableName 商户商品 Product自定义表名 app_products
func (Product) TableName() string {
	return "app_products"
}
