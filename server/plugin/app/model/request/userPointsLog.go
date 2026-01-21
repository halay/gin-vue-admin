package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserPointsLogSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	UserID         *int        `json:"userId" form:"userId"`
	Change         *int        `json:"change" form:"change"`
	BalanceAfter   *int        `json:"balanceAfter" form:"balanceAfter"`
	Type           *string     `json:"type" form:"type"`
	Status         *string     `json:"status" form:"status"`
	Reason         *string     `json:"reason" form:"reason"`
	OrderNo        *string     `json:"orderNo" form:"orderNo"`
	Remark         *string     `json:"remark" form:"remark"`
	MerchantID     *int        `json:"merchantId" form:"merchantId"`
	request.PageInfo
}
