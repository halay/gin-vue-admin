
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type AppDealerSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Name  *string `json:"name" form:"name"` 
       SalesCommission  *float64 `json:"salesCommission" form:"salesCommission"` 
       CommissionType  *int `json:"commissionType" form:"commissionType"` 
       ExpenseAllowance  *float64 `json:"expenseAllowance" form:"expenseAllowance"` 
       AllowanceType  *int `json:"allowanceType" form:"allowanceType"` 
       MerchantId  *int `json:"merchantId" form:"merchantId"` 
       Remark  *string `json:"remark" form:"remark"` 
    request.PageInfo
}
