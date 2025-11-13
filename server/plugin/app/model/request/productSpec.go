
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type ProductSpecSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       ProductID  *int `json:"productId" form:"productId"` 
       Name  *string `json:"name" form:"name"` 
       SortValue  *int `json:"sortValue" form:"sortValue"` 
       MerchantID  *int `json:"merchantId" form:"merchantId"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
