
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type UserPointsLogSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       UserID  *int `json:"userId" form:"userId"` 
       Change  *int `json:"change" form:"change"` 
       BalanceAfter  *int `json:"balanceAfter" form:"balanceAfter"` 
       Reason  *string `json:"reason" form:"reason"` 
       OrderNo  *string `json:"orderNo" form:"orderNo"` 
       Remark  *string `json:"remark" form:"remark"` 
    request.PageInfo
}
