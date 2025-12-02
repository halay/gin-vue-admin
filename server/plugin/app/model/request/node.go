
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type NodeSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Name  *string `json:"name" form:"name"` 
       SystemLevel  *string `json:"systemLevel" form:"systemLevel"` 
       LimitCount  *int `json:"limitCount" form:"limitCount"` 
       JoinFee  *float64 `json:"joinFee" form:"joinFee"` 
       SettlementCurrency  *string `json:"settlementCurrency" form:"settlementCurrency"` 
       CitySurchargePercent  *float64 `json:"citySurchargePercent" form:"citySurchargePercent"` 
       TokenDistributionX  *float64 `json:"tokenDistributionX" form:"tokenDistributionX"` 
       ReleaseWeeks  *int `json:"releaseWeeks" form:"releaseWeeks"` 
       CashCouponPercent  *float64 `json:"cashCouponPercent" form:"cashCouponPercent"` 
       PVBonusPercent  *float64 `json:"pvBonusPercent" form:"pvBonusPercent"` 
       RwaBrandFeeSharePercent  *float64 `json:"rwaBrandFeeSharePercent" form:"rwaBrandFeeSharePercent"` 
       AnnualFeeSharePercent  *float64 `json:"annualFeeSharePercent" form:"annualFeeSharePercent"` 
       OnlineTradeFeeSharePercent  *float64 `json:"onlineTradeFeeSharePercent" form:"onlineTradeFeeSharePercent"` 
       EquityPercent  *float64 `json:"equityPercent" form:"equityPercent"` 
       Status  *string `json:"status" form:"status"` 
       SortValue  *int `json:"sortValue" form:"sortValue"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
