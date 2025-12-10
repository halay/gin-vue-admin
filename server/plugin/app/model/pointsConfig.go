package model

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "time"
)

// PointsConfig 积分配置：现金购买积分、首单优惠等
type PointsConfig struct {
    global.GVA_MODEL
    Type         *string    `json:"type" form:"type" gorm:"comment:类型(first_order/recharge);column:type;size:32;"`
    Points       *int64     `json:"points" form:"points" gorm:"comment:积分数量;column:points;"`
    BonusPoints  *int64     `json:"bonusPoints" form:"bonusPoints" gorm:"comment:赠送积分数量;column:bonus_points;"`
    CurrentPrice *float64   `json:"currentPrice" form:"currentPrice" gorm:"comment:当前价格(人民币);column:current_price;"`
    OriginalPrice *float64  `json:"originalPrice" form:"originalPrice" gorm:"comment:原价(人民币);column:original_price;"`
    Sort         *int64     `json:"sort" form:"sort" gorm:"comment:显示顺序;column:sort;"`
    Limited      *bool      `json:"limited" form:"limited" gorm:"comment:是否限时;column:limited;"`
    StartAt      *time.Time `json:"startAt" form:"startAt" gorm:"comment:开始时间;column:start_at;"`
    EndAt        *time.Time `json:"endAt" form:"endAt" gorm:"comment:结束时间;column:end_at;"`
    Status       *string    `json:"status" form:"status" gorm:"comment:启用/禁用;column:status;default:enabled;"`
    Description  *string    `json:"description" form:"description" gorm:"comment:说明;column:description;type:text;"`
}

func (PointsConfig) TableName() string { return "app_points_configs" }

