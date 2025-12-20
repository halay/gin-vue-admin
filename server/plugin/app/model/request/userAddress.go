
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type UserAddressSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       UserID  *int `json:"userId" form:"userId"` 
       Receiver  *string `json:"receiver" form:"receiver"` 
       Phone  *string `json:"phone" form:"phone"` 
       Country  *string `json:"country" form:"country"` 
       Province  *string `json:"province" form:"province"` 
       City  *string `json:"city" form:"city"` 
       District  *string `json:"district" form:"district"` 
       Address  *string `json:"address" form:"address"` 
       Zipcode  *string `json:"zipcode" form:"zipcode"` 
       IsDefault  *bool `json:"isDefault" form:"isDefault"` 
       Status  *string `json:"status" form:"status"` 
       Remark  *string `json:"remark" form:"remark"` 
    request.PageInfo
}
