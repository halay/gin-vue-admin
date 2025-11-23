package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Banner Banner图片 结构体
type Banner struct {
	global.GVA_MODEL
	Title     *string    `json:"title" form:"title" gorm:"comment:标题;column:title;" binding:"required"`                       //标题
	ImageUrl  string     `json:"imageUrl" form:"imageUrl" gorm:"comment:图片地址;column:image_url;" binding:"required"`           //图片地址
	LinkUrl   *string    `json:"linkUrl" form:"linkUrl" gorm:"comment:跳转链接;column:link_url;"`                                 //跳转链接
	Position  *string    `json:"position" form:"position" gorm:"comment:位置标识;column:position;" binding:"required"`            //位置标识
	Sort      *int32     `json:"sort" form:"sort" gorm:"default:0;comment:排序;column:sort;"`                                   //排序
	Status    *string    `json:"status" form:"status" gorm:"default:enabled;comment:启用/禁用;column:status;" binding:"required"` //状态
	StartTime *time.Time `json:"startTime" form:"startTime" gorm:"comment:开始时间;column:start_time;"`                           //开始时间
	EndTime   *time.Time `json:"endTime" form:"endTime" gorm:"comment:结束时间;column:end_time;"`                                 //结束时间
}

// TableName Banner图片 Banner自定义表名 app_banners
func (Banner) TableName() string {
	return "app_banners"
}
