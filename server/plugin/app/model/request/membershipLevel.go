
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type MembershipLevelSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       Name  *string `json:"name" form:"name"` 
       Code  *string `json:"code" form:"code"` 
       Status  *string `json:"status" form:"status"` 
       SortValue  *int `json:"sortValue" form:"sortValue"` 
       Requirement  *string `json:"requirement" form:"requirement"` 
       RequirementValue  *float64 `json:"requirementValue" form:"requirementValue"` 
       Benefits  string `json:"benefits" form:"benefits"` 
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
