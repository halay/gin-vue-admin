package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppAgentTransactionsSearch struct {
	request.PageInfo
	MerchantId *uint `json:"merchantId" form:"merchantId"`
	BeneficiaryId *uint `json:"beneficiaryId" form:"beneficiaryId"`
}
