package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MerchantAnnouncement 商户公告信息 结构体
type MerchantAnnouncement struct {
	global.GVA_MODEL
	Title      *string    `json:"title" form:"title" gorm:"comment:公告标题;column:title;size:255;" binding:"required"`               //公告标题
	Content    *string    `json:"content" form:"content" gorm:"comment:公告内容;column:content;size:0;type:text;" binding:"required"` //公告内容
	Status     *string    `json:"status" form:"status" gorm:"default:enabled;comment:启用/禁用;column:status;" binding:"required"`    //状态
	PublishAt  *time.Time `json:"publishAt" form:"publishAt" gorm:"comment:发布时间;column:publish_at;size:0;"`                       //发布时间
	ExpireAt   *time.Time `json:"expireAt" form:"expireAt" gorm:"comment:过期时间;column:expire_at;size:0;"`                          //过期时间
	MerchantID *int       `json:"merchantId" form:"merchantId" gorm:"comment:关联商户;column:merchant_id;"`                           //关联商户
}

// TableName 商户公告信息 MerchantAnnouncement自定义表名 app_merchant_announcements
func (MerchantAnnouncement) TableName() string {
	return "app_merchant_announcements"
}
