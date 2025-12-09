package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type AppConsultation struct {
    global.GVA_MODEL
    Title      *string    `json:"title" form:"title" gorm:"column:title;size:255;"`
    CoverImage string     `json:"coverImage" form:"coverImage" gorm:"column:cover_image;"`
    Content    *string    `json:"content" form:"content" gorm:"column:content;type:text;"`
    JumpURL    *string    `json:"jumpUrl" form:"jumpUrl" gorm:"column:jump_url;size:512;"`
    Category   *string    `json:"category" form:"category" gorm:"column:category;size:16;default:1;"`
    Status     *string    `json:"status" form:"status" gorm:"column:status;default:enabled;"`
    Sort       *int64     `json:"sort" form:"sort" gorm:"column:sort;"`
    PublishAt  *time.Time `json:"publishAt" form:"publishAt" gorm:"column:publish_at;"`
    ExpireAt   *time.Time `json:"expireAt" form:"expireAt" gorm:"column:expire_at;"`
    Excerpt    *string    `json:"excerpt" form:"excerpt" gorm:"column:excerpt;type:text;"`
    Clicks     *int64     `json:"clicks" form:"clicks" gorm:"column:clicks;default:0;"`
}

func (AppConsultation) TableName() string { return "app_consultations" }
