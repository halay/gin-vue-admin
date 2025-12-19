
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type MerchantPointsSettingsSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       MerchantID  *int `json:"merchantId" form:"merchantId"` 
       TokenName  *string `json:"tokenName" form:"tokenName"` 
       Symbol  *string `json:"symbol" form:"symbol"` 
       InviterReward  *int `json:"inviterReward" form:"inviterReward"` 
       InviteeReward  *int `json:"inviteeReward" form:"inviteeReward"` 
       RegisterReward  *int `json:"registerReward" form:"registerReward"` 
       CoverageMode  *string `json:"coverageMode" form:"coverageMode"` 
       DailyLimit  *int `json:"dailyLimit" form:"dailyLimit"` 
       Status  *string `json:"status" form:"status"` 
       Sort  *int `json:"sort" form:"sort"` 
       Remark  *string `json:"remark" form:"remark"` 
    request.PageInfo
}
