package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:文件名"`                                // 文件名
	ClassId    int    `json:"classId" form:"classId" gorm:"default:0;type:int;column:class_id;comment:分类id;"` // 分类id
	Url        string `json:"url" form:"url" gorm:"column:url;comment:文件地址"`                                  // 文件地址
	Tag        string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签"`                                  // 文件标签
	Key        string `json:"key" form:"key" gorm:"column:key;comment:编号"`                                    // 编号
	UserID     uint   `json:"userId" gorm:"index;comment:用户ID;column:user_id;"`
	MerchantID *uint  `json:"merchantId" gorm:"index;comment:商户ID;column:merchant_id;"`
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
