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
	ID                uint            `json:"id"`
	Email             string          `json:"email"`
	Nickname          string          `json:"nickname"`
	Avatar            string          `json:"avatar"`
	Phone             string          `json:"phone"`
	Status            string          `json:"status"`
	EmailVerified     bool            `json:"emailVerified"`
	LastLoginTime     *time.Time      `json:"lastLoginTime,omitempty"`
	LastLoginIP       string          `json:"lastLoginIP,omitempty"`
	InviteCode        string          `json:"inviteCode"`
	InviterID         uint            `json:"inviterId"`
	InviteLevel       int             `json:"inviteLevel"`
	MembershipLevelID int64           `json:"membershipLevelId"`
	NodeID            int64           `json:"nodeId"`
	MerchantID        int64           `json:"MerchantID"`
	PointsBalance     int64           `json:"pointsBalance"` // 废弃，保留兼容
	PointsAccounts    []PointsAccount `json:"pointsAccounts"` // 新增：所有积分账户
	Merchant          model.Merchants `json:"merchant"`
	ShareholderProfitID int64 `json:"shareholderProfitId"`
	ShareholderProfitName string `json:"shareholderProfitName"`
	AppDealerID       int64   `json:"appDealerId"`
	AppDealerName     string  `json:"appDealerName"`
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
	User          model.AppUsers  `json:"user"`
	Token         string          `json:"token"`
	ExpiresAt     int64           `json:"expiresAt"`
	PointsBalance int64           `json:"pointsBalance"` // 废弃，保留兼容
	PointsAccounts []PointsAccount `json:"pointsAccounts"` // 新增：所有积分账户
	//Merchant  model.Merchants `json:"merchant"`
	ShareholderProfitID int64 `json:"shareholderProfitId"`
	ShareholderProfitName string `json:"shareholderProfitName"`
	AppDealerID       int64   `json:"appDealerId"`
	AppDealerName     string  `json:"appDealerName"`
}
