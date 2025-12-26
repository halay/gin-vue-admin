package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DownlinePurchaseRecordSearch struct {
	request.PageInfo
	OrderNo           *string    `json:"orderNo" form:"orderNo"`
	MerchantId        *uint      `json:"merchantId" form:"merchantId"`
	UplineUserId      *uint      `json:"uplineUserId" form:"uplineUserId"`
	UplineUserEmail   *string    `json:"uplineUserEmail" form:"uplineUserEmail"`
	PurchaseUserId    *uint      `json:"purchaseUserId" form:"purchaseUserId"`
	PurchaseUserEmail *string    `json:"purchaseUserEmail" form:"purchaseUserEmail"`
	IsDirect          *bool      `json:"isDirect" form:"isDirect"`
	PayStatus         *string    `json:"payStatus" form:"payStatus"`
	OrderStatus       *string    `json:"orderStatus" form:"orderStatus"`
	StartCreatedAt    *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt      *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
}
