package request

import (
	commonrequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SearchPublic struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Keyword        string      `json:"keyword" form:"keyword"`
	MerchantID     *int64      `json:"merchantId" form:"merchantId"`
	CategoryID     *int        `json:"categoryId" form:"categoryId"`
	Days           int         `json:"days" form:"days"`
	Metric         string      `json:"metric" form:"metric"`
	commonrequest.PageInfo
}
