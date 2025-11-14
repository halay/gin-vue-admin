package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProductSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Name           *string     `json:"name" form:"name"`
	Type           *string     `json:"type" form:"type"`
	Status         *string     `json:"status" form:"status"`
	CategoryID     *int        `json:"categoryId" form:"categoryId"`
	CoverImage     string      `json:"coverImage" form:"coverImage"`
	CarouselImages string      `json:"carouselImages" form:"carouselImages"`
	Video          string      `json:"video" form:"video"`
	Detail         string      `json:"detail" form:"detail"`
	HasVariants    *bool       `json:"hasVariants" form:"hasVariants"`
	EnablePoints   *bool       `json:"enablePoints" form:"enablePoints"`
	Price          *float64    `json:"price" form:"price"`
	Points         *int        `json:"points" form:"points"`
	Stock          *int        `json:"stock" form:"stock"`
	MerchantID     *int64      `json:"merchantId" form:"merchantId"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
