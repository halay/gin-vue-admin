
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserPointsLog 用户积分流水 结构体
type UserPointsLog struct {
    global.GVA_MODEL
  UserID  *int64 `json:"userId" form:"userId" gorm:"comment:app用户ID;column:user_id;" binding:"required"`  //用户
  Change  *int64 `json:"change" form:"change" gorm:"comment:积分变动，正负;column:change;" binding:"required"`  //变动值
  BalanceAfter  *int64 `json:"balanceAfter" form:"balanceAfter" gorm:"comment:变动后余额;column:balance_after;" binding:"required"`  //变动后余额
  Reason  *string `json:"reason" form:"reason" gorm:"comment:变动原因;column:reason;size:255;"`  //原因
  OrderNo  *string `json:"orderNo" form:"orderNo" gorm:"comment:关联订单号;column:order_no;size:64;"`  //关联订单号
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:255;"`  //备注
}


// TableName 用户积分流水 UserPointsLog自定义表名 app_user_points_logs
func (UserPointsLog) TableName() string {
    return "app_user_points_logs"
}







