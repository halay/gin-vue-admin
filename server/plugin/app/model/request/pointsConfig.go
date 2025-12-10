package request

import (
    commonrequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "time"
)

type PointsConfigSearch struct {
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    Type           *string     `json:"type" form:"type"`
    Status         *string     `json:"status" form:"status"`
    Limited        *bool       `json:"limited" form:"limited"`
    StartAtRange   []time.Time `json:"startAtRange" form:"startAtRange[]"`
    commonrequest.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}

