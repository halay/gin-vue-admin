
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type MerchantAnnouncementSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Title  *string `json:"title" form:"title"` 
       Status  string `json:"status" form:"status"` 
       PublishAtRange  []time.Time  `json:"publishAtRange" form:"publishAtRange[]"`
       MerchantID  *int `json:"merchantId" form:"merchantId"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
