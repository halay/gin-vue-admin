
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Node 节点策略 结构体
type Node struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:节点名称;column:name;size:64;" binding:"required"`  //节点名称
  SystemLevel  *string `json:"systemLevel" form:"systemLevel" gorm:"comment:系统层级;column:system_level;size:32;" binding:"required"`  //系统层级ID
  LimitCount  *int64 `json:"limitCount" form:"limitCount" gorm:"comment:限额数量;column:limit_count;"`  //限额数量(0为不限)
  JoinFee  *float64 `json:"joinFee" form:"joinFee" gorm:"comment:加盟费用;column:join_fee;"`  //加盟费用
  SettlementCurrency  *string `json:"settlementCurrency" form:"settlementCurrency" gorm:"comment:结算货币;column:settlement_currency;size:16;"`  //结算货币
  CitySurchargePercent  *float64 `json:"citySurchargePercent" form:"citySurchargePercent" gorm:"comment:加收比例%;column:city_surcharge_percent;"`  //一线城市/特定区域加收比例(%)
  TokenDistributionX  *float64 `json:"tokenDistributionX" form:"tokenDistributionX" gorm:"comment:代币配送倍数X;column:token_distribution_x;"`  //代币配送倍数(X倍)
  ReleaseWeeks  *int64 `json:"releaseWeeks" form:"releaseWeeks" gorm:"comment:释放周期周;column:release_weeks;"`  //释放周期(周)
  CashCouponPercent  *float64 `json:"cashCouponPercent" form:"cashCouponPercent" gorm:"comment:现金券赠送%;column:cash_coupon_percent;"`  //平台消费现金券赠送(%)
  PVBonusPercent  *float64 `json:"pvBonusPercent" form:"pvBonusPercent" gorm:"comment:消费级差PV分润%;column:pv_bonus_percent;"`  //消费级差PV分润(%)
  RwaBrandFeeSharePercent  *float64 `json:"rwaBrandFeeSharePercent" form:"rwaBrandFeeSharePercent" gorm:"comment:RWA品牌费用分成%;column:rwa_brand_fee_share_percent;"`  //RWA品牌上市费用分成(%)
  AnnualFeeSharePercent  *float64 `json:"annualFeeSharePercent" form:"annualFeeSharePercent" gorm:"comment:入驻年费级差分成%;column:annual_fee_share_percent;"`  //入驻年费级差分成(%)
  OnlineTradeFeeSharePercent  *float64 `json:"onlineTradeFeeSharePercent" form:"onlineTradeFeeSharePercent" gorm:"comment:在线交易手续费分成%;column:online_trade_fee_share_percent;"`  //在线交易手续费分成(%)
  EquityPercent  *float64 `json:"equityPercent" form:"equityPercent" gorm:"comment:股权百分比;column:equity_percent;"`  //股权(%)
  Status  *string `json:"status" form:"status" gorm:"default:1;comment:启用/禁用;column:status;"`  //状态
  Sort  *int64 `json:"sort" form:"sort" gorm:"comment:排序;column:sort;"`  //排序
}


// TableName 节点策略 Node自定义表名 app_nodes
func (Node) TableName() string {
    return "app_nodes"
}







