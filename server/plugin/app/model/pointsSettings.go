
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PointsSettings 平台积分基础设置与奖励规则 结构体
type PointsSettings struct {
    global.GVA_MODEL
  TokenName  *string `json:"tokenName" form:"tokenName" gorm:"comment:积分资产全称;column:token_name;size:128;" binding:"required"`  //积分全称
  Symbol  *string `json:"symbol" form:"symbol" gorm:"comment:积分缩写;column:symbol;size:32;" binding:"required"`  //积分缩写(Symbol)
  RegisterReward  *int64 `json:"registerReward" form:"registerReward" gorm:"comment:注册用户奖励积分;column:register_reward;" binding:"required"`  //注册奖励积分
  InviterReward  *int64 `json:"inviterReward" form:"inviterReward" gorm:"comment:邀请人奖励积分;column:inviter_reward;" binding:"required"`  //邀请人奖励积分
  InviteeReward  *int64 `json:"inviteeReward" form:"inviteeReward" gorm:"comment:被邀请人奖励积分;column:invitee_reward;" binding:"required"`  //被邀请人奖励积分
  Status  *string `json:"status" form:"status" gorm:"default:enabled;comment:启用/禁用;column:status;"`  //状态
  Sort  *int64 `json:"sort" form:"sort" gorm:"comment:排序;column:sort;"`  //排序
  Remark  *string `json:"remark" form:"remark" gorm:"comment:说明;column:remark;size:255;"`  //说明
}


// TableName 平台积分基础设置与奖励规则 PointsSettings自定义表名 app_points_settings
func (PointsSettings) TableName() string {
    return "app_points_settings"
}







