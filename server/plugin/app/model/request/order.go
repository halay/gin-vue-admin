package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	OrderNo        *string     `json:"orderNo" form:"orderNo"`
	UserID         *int        `json:"userId" form:"userId"`
	MerchantID     *int        `json:"merchantId" form:"merchantId"`
	TotalAmount    *float64    `json:"totalAmount" form:"totalAmount"`
	TotalPoints    *int        `json:"totalPoints" form:"totalPoints"`
	PayMethod      *string     `json:"payMethod" form:"payMethod"`
	PayStatus      *string     `json:"payStatus" form:"payStatus"`
	Status         *string     `json:"status" form:"status"`
	ConsigneeName  *string     `json:"consigneeName" form:"consigneeName"`
	ConsigneePhone *string     `json:"consigneePhone" form:"consigneePhone"`
	Address        *string     `json:"address" form:"address"`
	Country        *string     `json:"country" form:"country"`
	PostalCode     *string     `json:"postalCode" form:"postalCode"`
	DeliveryStatus *string     `json:"deliveryStatus" form:"deliveryStatus"`
	ExpressName    *string     `json:"expressName" form:"expressName"`
	ExpressNo      *string     `json:"expressNo" form:"expressNo"`
	Remark         *string     `json:"remark" form:"remark"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
