
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ShareholderProfit 股东分润 结构体
type ShareholderProfit struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:股东名称;column:name;size:255;" binding:"required"`  //股东名称
  ShareRatio  *float64 `json:"shareRatio" form:"shareRatio" gorm:"default:0;comment:分润比例;column:share_ratio;size:10,2;" binding:"required"`  //分润比例
}


// TableName 股东分润 ShareholderProfit自定义表名 shareholder_profit
func (ShareholderProfit) TableName() string {
    return "shareholder_profit"
}







