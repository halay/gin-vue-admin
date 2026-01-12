package response

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
)

// LoginResponse 登录响应
type LoginResponse struct {
	User  model.AppUsers `json:"user"`
	Token string         `json:"token"`
}

// UserResponse 用户信息响应
type UserResponse struct {
	ID                    uint            `json:"id"`
	Email                 string          `json:"email"`
	Nickname              string          `json:"nickname"`
	Avatar                string          `json:"avatar"`
	Phone                 string          `json:"phone"`
	Status                string          `json:"status"`
	EmailVerified         bool            `json:"emailVerified"`
	LastLoginTime         *time.Time      `json:"lastLoginTime,omitempty"`
	LastLoginIP           string          `json:"lastLoginIP,omitempty"`
	InviteCode            string          `json:"inviteCode"`
	InviterID             uint            `json:"inviterId"`
	InviteLevel           int             `json:"inviteLevel"`
	MembershipLevelID     int64           `json:"membershipLevelId"`
	NodeID                int64           `json:"nodeId"`
	MerchantID            int64           `json:"MerchantID"`
	PointsBalance         int64           `json:"pointsBalance"`  // 废弃，保留兼容
	PointsAccounts        []PointsAccount `json:"pointsAccounts"` // 新增：所有积分账户
	Merchant              model.Merchants `json:"merchant"`
	ShareholderProfitID   int64           `json:"shareholderProfitId"`
	ShareholderProfitName string          `json:"shareholderProfitName"`
	AppDealerID           int64           `json:"appDealerId"`
	AppDealerName         string          `json:"appDealerName"`
}
type SubUserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
}

// DashboardResponse APP首页/个人中心数据聚合响应
type DashboardResponse struct {
	User       UserSummary       `json:"user"`
	Commission CommissionSummary `json:"commission"`
	Points     PointsSummary     `json:"points"`
	Invite     InviteSummary     `json:"invite"`
}

type InviteSummary struct {
	DirectCount   int64 `json:"directCount"`
	IndirectCount int64 `json:"indirectCount"`
}
type UserSummary struct {
	Nickname            string `json:"nickname"`
	Avatar              string `json:"avatar"`
	MembershipLevelName string `json:"membershipLevelName"`
	AgentLevelName      string `json:"agentLevelName"` // 代理等级或身份名称
	DealerName          string `json:"dealerName"`
	ProfitName          string `json:"profitName"`
	Email               string `json:"email"`
	ID                  uint   `json:"id"`
}

type CommissionSummary struct {
	Total     float64 `json:"total"`     // 总返佣金额
	Pending   float64 `json:"pending"`   // 待提现金额
	Withdrawn float64 `json:"withdrawn"` // 已提现金额
}

type PointsSummary struct {
	Total    int64  `json:"total"`    // 总获取积分
	Current  int64  `json:"current"`  // 当前剩余积分 (通常指平台积分或主积分)
	Consumed int64  `json:"consumed"` // 已消耗积分
	Pending  int64  `json:"pending"`  // 待入账积分
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}

type PointsAccount struct {
	MerchantID *int64  `json:"merchantId"`
	TokenName  *string `json:"tokenName"`
	Symbol     *string `json:"symbol"`
	Balance    int64   `json:"balance"`
	IsPlatform bool    `json:"isPlatform"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	List     []UserResponse `json:"list"`
	Total    int64          `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
}
type AppLoginResponse struct {
	User           model.AppUsers  `json:"user"`
	Token          string          `json:"token"`
	ExpiresAt      int64           `json:"expiresAt"`
	PointsBalance  int64           `json:"pointsBalance"`  // 废弃，保留兼容
	PointsAccounts []PointsAccount `json:"pointsAccounts"` // 新增：所有积分账户
	//Merchant  model.Merchants `json:"merchant"`
	ShareholderProfitID   int64  `json:"shareholderProfitId"`
	ShareholderProfitName string `json:"shareholderProfitName"`
	AppDealerID           int64  `json:"appDealerId"`
	AppDealerName         string `json:"appDealerName"`
}
