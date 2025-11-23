package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MerchantsSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	MerchantName   *string     `json:"merchantName" form:"merchantName"`
	BusinessScope  *string     `json:"businessScope" form:"businessScope"`
	ContactPerson  *string     `json:"contactPerson" form:"contactPerson"`
	ContactPhone   *string     `json:"contactPhone" form:"contactPhone"`
	Email          *string     `json:"email" form:"email"`
    Status         *string     `json:"status" form:"status"`
    IsRecommended  *bool       `json:"isRecommended" form:"isRecommended"`
    CategoryID     *int64      `json:"categoryId" form:"categoryId"`
    request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
	// 在结构体中新增如下字段
	Founder          *string `json:"founder" form:"founder"`
	FounderPhone     *string `json:"founderPhone" form:"founderPhone"`
	DigitalAssetName *string `json:"digitalAssetName" form:"digitalAssetName"`
}
