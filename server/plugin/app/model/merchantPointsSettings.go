
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MerchantPointsSettings 商户积分营销配置 结构体
type MerchantPointsSettings struct {
    global.GVA_MODEL
  MerchantID  *int64 `json:"merchantId" form:"merchantId" gorm:"comment:关联商户ID;column:merchant_id;"`  //商户ID
  TokenName  *string `json:"tokenName" form:"tokenName" gorm:"comment:品牌积分资产全称;column:token_name;size:128;" binding:"required"`  //积分全称
  Symbol  *string `json:"symbol" form:"symbol" gorm:"comment:品牌积分缩写;column:symbol;size:32;" binding:"required"`  //积分缩写(Symbol)
  InviterReward  *int64 `json:"inviterReward" form:"inviterReward" gorm:"comment:邀请人奖励积分;column:inviter_reward;" binding:"required"`  //邀请人奖励积分
  InviteeReward  *int64 `json:"inviteeReward" form:"inviteeReward" gorm:"comment:被邀请人奖励积分;column:invitee_reward;" binding:"required"`  //被邀请人奖励积分
  RegisterReward  *int64 `json:"registerReward" form:"registerReward" gorm:"comment:注册用户奖励积分;column:register_reward;" binding:"required"`  //注册奖励积分
  CoverageMode  *string `json:"coverageMode" form:"coverageMode" gorm:"default:directed;comment:全平台(platform)/品牌定向(directed);column:coverage_mode;size:16;" binding:"required"`  //分发受众策略
  DailyLimit  *int64 `json:"dailyLimit" form:"dailyLimit" gorm:"comment:每日限制人次;column:daily_limit;"`  //每日发放上限(人次)
  Status  *string `json:"status" form:"status" gorm:"default:enabled;comment:启用/禁用;column:status;"`  //状态
  Sort  *int64 `json:"sort" form:"sort" gorm:"comment:排序;column:sort;"`  //排序
  Remark  *string `json:"remark" form:"remark" gorm:"comment:说明;column:remark;size:255;"`  //说明
}


// TableName 商户积分营销配置 MerchantPointsSettings自定义表名 app_merchant_points_settings
func (MerchantPointsSettings) TableName() string {
    return "app_merchant_points_settings"
}






