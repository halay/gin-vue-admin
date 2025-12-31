package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ExaAttachmentCategory struct {
	global.GVA_MODEL
	Name       string                   `json:"name" form:"name" gorm:"default:null;type:varchar(255);column:name;comment:分类名称;"`
	Pid        uint                     `json:"pid" form:"pid" gorm:"default:0;type:int;column:pid;comment:父节点ID;"`
	UserID     uint                     `json:"userId" gorm:"index;comment:用户ID;column:user_id;"`
	MerchantID *uint                    `json:"merchantId" gorm:"index;comment:商户ID;column:merchant_id;"`
	Children   []*ExaAttachmentCategory `json:"children" gorm:"-"`
}

func (ExaAttachmentCategory) TableName() string {
	return "exa_attachment_category"
}
