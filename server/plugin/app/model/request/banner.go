package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BannerSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	Title          *string     `json:"title" form:"title"`
	Position       *string     `json:"position" form:"position"`
	Status         *string     `json:"status" form:"status"`
	StartTimeRange []time.Time `json:"startTimeRange" form:"startTimeRange[]"`
	EndTimeRange   []time.Time `json:"endTimeRange" form:"endTimeRange[]"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
