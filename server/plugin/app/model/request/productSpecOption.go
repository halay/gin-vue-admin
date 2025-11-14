
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type ProductSpecOptionSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       ProductID  *int `json:"productId" form:"productId"` 
       SpecID  *int `json:"specId" form:"specId"` 
       Value  *string `json:"value" form:"value"` 
       SortValue  *int `json:"sortValue" form:"sortValue"` 
       MerchantID  *int64 `json:"merchantId" form:"merchantId"` 
    request.PageInfo
}
