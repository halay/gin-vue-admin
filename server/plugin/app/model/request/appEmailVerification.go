package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppEmailVerificationSearch struct {
	CreatedAtRange  []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Email           *string     `json:"email" form:"email"`
	Purpose         *string     `json:"purpose" form:"purpose"`
	ExpireTimeRange []time.Time `json:"expireTimeRange" form:"expireTimeRange[]"`
	Used            *bool       `json:"used" form:"used"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
