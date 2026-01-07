package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AgentTransactionSearch struct {
	request.PageInfo
	OrderNo        *string    `json:"orderNo" form:"orderNo"`
	MerchantId     *uint      `json:"merchantId" form:"merchantId"`
	BeneficiaryId  *uint      `json:"beneficiaryId" form:"beneficiaryId"`
	SourceUserId   *uint      `json:"sourceUserId" form:"sourceUserId"`
	AgentLevelName *string    `json:"agentLevelName" form:"agentLevelName"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
}
