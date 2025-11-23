package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OrderItemSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	OrderID        *int        `json:"orderId" form:"orderId"`
	ProductID      *int        `json:"productId" form:"productId"`
	SkuID          *int        `json:"skuId" form:"skuId"`
	ProductName    *string     `json:"productName" form:"productName"`
	SkuAttrs       string      `json:"skuAttrs" form:"skuAttrs"`
	Price          *float64    `json:"price" form:"price"`
	Points         *int        `json:"points" form:"points"`
	Quantity       *int        `json:"quantity" form:"quantity"`
	TotalAmount    *float64    `json:"totalAmount" form:"totalAmount"`
	TotalPoints    *int        `json:"totalPoints" form:"totalPoints"`
	CoverImage     string      `json:"coverImage" form:"coverImage"`
	MerchantID     *int        `json:"merchantId" form:"merchantId"`
	request.PageInfo
}
