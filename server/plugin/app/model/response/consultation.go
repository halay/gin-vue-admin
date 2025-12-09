package response

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ConsultationResponse struct {
    global.GVA_MODEL
    Title      *string    `json:"title" form:"title" gorm:"column:title;size:255;"`
    CoverImage string     `json:"coverImage" form:"coverImage" gorm:"column:cover_image;"`
    JumpURL    *string    `json:"jumpUrl" form:"jumpUrl" gorm:"column:jump_url;size:512;"`
    Category   *string    `json:"category" form:"category" gorm:"column:category;size:16;"`
    Sort       *int64     `json:"sort" form:"sort" gorm:"column:sort;"`
    PublishAt  *time.Time `json:"publishAt" form:"publishAt" gorm:"column:publish_at;"`
    Clicks     *int64     `json:"clicks" form:"clicks" gorm:"column:clicks;default:0;"`
}
