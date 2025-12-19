
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type PointsSettingsSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       TokenName  *string `json:"tokenName" form:"tokenName"` 
       Symbol  *string `json:"symbol" form:"symbol"` 
       RegisterReward  *int `json:"registerReward" form:"registerReward"` 
       InviterReward  *int `json:"inviterReward" form:"inviterReward"` 
       InviteeReward  *int `json:"inviteeReward" form:"inviteeReward"` 
       Status  *string `json:"status" form:"status"` 
       SortValue  *int `json:"sortValue" form:"sortValue"` 
       Remark  *string `json:"remark" form:"remark"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
