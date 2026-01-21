package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserPointsAccountSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	UserID         *int        `json:"userId" form:"userId"`
	Balance        *int        `json:"balance" form:"balance"`
	MerchantID     *int        `json:"merchantId" form:"merchantId"`
	request.PageInfo
}
