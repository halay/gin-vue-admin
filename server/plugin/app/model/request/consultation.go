package request

import (
	commonrequest "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ConsultationSearch struct {
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
    Title          *string     `json:"title" form:"title"`
    Category       *string     `json:"category" form:"category"`
    Status         *string     `json:"status" form:"status"`
    PublishAtRange []time.Time `json:"publishAtRange" form:"publishAtRange[]"`
    commonrequest.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
