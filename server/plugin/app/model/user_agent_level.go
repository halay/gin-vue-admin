package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserAgentLevel 用户代理等级记录
type UserAgentLevel struct {
	global.GVA_MODEL
	UserID         uint      `json:"userId" gorm:"index;comment:用户ID;column:user_id;"`
	MerchantID     uint      `json:"merchantId" gorm:"index;comment:商户ID;column:merchant_id;"`
	AgentLevelID   uint      `json:"agentLevelId" gorm:"index;comment:代理等级ID;column:agent_level_id;"`
	AgentLevelName string    `json:"agentLevelName" gorm:"comment:代理等级名称;column:agent_level_name;size:255;"`
	Status         string    `json:"status" gorm:"comment:状态(active/inactive);column:status;size:20;"`
	AchievedAt     time.Time `json:"achievedAt" gorm:"comment:达成时间;column:achieved_at;"`
}

func (UserAgentLevel) TableName() string {
	return "app_user_agent_levels"
}
