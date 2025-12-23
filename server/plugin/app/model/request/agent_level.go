package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AgentLevelSearch struct {
	request.PageInfo
	MerchantId *int64 `json:"merchantId" form:"merchantId"`
}
