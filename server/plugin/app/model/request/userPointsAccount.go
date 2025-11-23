package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserPointsAccountSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	UserID         *int        `json:"userId" form:"userId"`
	Balance        *int        `json:"balance" form:"balance"`
	request.PageInfo
}
