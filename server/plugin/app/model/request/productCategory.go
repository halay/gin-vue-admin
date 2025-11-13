
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type ProductCategorySearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Name  *string `json:"name" form:"name"` 
       Status  *string `json:"status" form:"status"` 
       SortValue  *int `json:"sortValue" form:"sortValue"` 
       Description  string `json:"description" form:"description"` 
       MerchantID  *int `json:"merchantId" form:"merchantId"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
