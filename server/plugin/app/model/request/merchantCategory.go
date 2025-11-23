
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type MerchantCategorySearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Name  *string `json:"name" form:"name"` 
       Status  *string `json:"status" form:"status"` 
       SortValue  *int `json:"sortValue" form:"sortValue"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
