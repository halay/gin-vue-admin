
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type ProductSkuSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       ProductID  *int `json:"productId" form:"productId"` 
       SKUCode  *string `json:"skuCode" form:"skuCode"` 
       Attrs  string `json:"attrs" form:"attrs"` 
       Price  *float64 `json:"price" form:"price"` 
       Points  *int `json:"points" form:"points"` 
       Image  string `json:"image" form:"image"` 
       Stock  *int `json:"stock" form:"stock"` 
       Status  *string `json:"status" form:"status"` 
       MerchantID  *int64 `json:"merchantId" form:"merchantId"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
