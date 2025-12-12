package request

import (
    commonrequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "time"
)

type PointsRechargeSearch struct {
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    Status         *string     `json:"status" form:"status"`
    PayMethod      *string     `json:"payMethod" form:"payMethod"`
    commonrequest.PageInfo
}

