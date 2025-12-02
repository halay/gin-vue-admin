package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

var _ system.Login = new(AppUsers)

// AppUsers appUsers表 结构体
type AppUsers struct {
	global.GVA_MODEL
	Email         *string    `json:"email" form:"email" gorm:"uniqueIndex;comment:用户邮箱;column:email;size:255;" binding:"required"`         //用户邮箱
	Password      *string    `json:"password" form:"password" gorm:"comment:用户密码;column:password;size:255;" binding:"required"`            //用户密码
	Nickname      *string    `json:"nickname" form:"nickname" gorm:"comment:用户昵称;column:nickname;size:50;"`                                //用户昵称
	Uuid          uuid.UUID  `json:"uuid" form:"uuid" gorm:"comment:用户UUID;column:uuid;size:191;"`                                         //用户UUID
	Avatar        *string    `json:"avatar" form:"avatar" gorm:"comment:用户头像URL;column:avatar;size:255;"`                                  //用户头像URL
	Phone         *string    `json:"phone" form:"phone" gorm:"comment:用户手机号;column:phone;size:20;"`                                        //用户手机号
	Status        *string    `json:"status" form:"status" gorm:"comment:用户状态;column:status;size:19;"`                                      //用户状态
	LastLoginTime *time.Time `json:"lastLoginTime" form:"lastLoginTime" gorm:"comment:最后登录时间;column:last_login_time;"`                     //最后登录时间
	LastLoginIP   *string    `json:"lastLoginIp" form:"lastLoginIp" gorm:"comment:最后登录IP;column:last_login_ip;size:50;"`                   //最后登录IP
	EmailVerified *bool      `json:"emailVerified" form:"emailVerified" gorm:"comment:邮箱是否已验证;column:email_verified;"`                     //邮箱是否已验证
	AuthorityId   *uint      `json:"authorityId" form:"authorityId" gorm:"default:123;comment:用户角色ID;column:authority_id;"`                //用户角色ID
	InviteCode    *string    `json:"inviteCode" form:"inviteCode" gorm:"uniqueIndex;comment:邀请码(不区分大小写存储为大写);column:invite_code;size:16;"` //邀请码
	InviterID     *uint      `json:"inviterId" form:"inviterId" gorm:"comment:邀请人用户ID;column:inviter_id;"`                                 //邀请人ID
	InvitePath    *string    `json:"invitePath" form:"invitePath" gorm:"comment:邀请路径;column:invite_path;type:text;"`                       //邀请路径 a/b/c...
	InviteLevel   *int       `json:"inviteLevel" form:"inviteLevel" gorm:"comment:邀请层级;column:invite_level;"`                              //层级深度
	MembershipLevelID *int64 `json:"membershipLevelId" form:"membershipLevelId" gorm:"comment:会员等级ID;column:membership_level_id;"`
	NodeID *int64 `json:"nodeId" form:"nodeId" gorm:"comment:所属节点ID;column:node_id;"`
}

func (u AppUsers) GetUsername() string {
	return *u.Email
}

func (u AppUsers) GetNickname() string {
	return *u.Nickname
}

func (u AppUsers) GetUUID() uuid.UUID {
	return u.Uuid
}

func (u AppUsers) GetUserId() uint {
	return u.ID
}

func (u AppUsers) GetAuthorityId() uint {
	return *u.AuthorityId
}

func (u AppUsers) GetUserInfo() any {
	return u
}

// TableName appUsers表 AppUsers自定义表名 app_users
func (AppUsers) TableName() string {
	return "app_users"
}
