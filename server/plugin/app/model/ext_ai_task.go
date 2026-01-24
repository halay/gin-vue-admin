package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ExtAiTask extAiTask表 结构体
type ExtAiTask struct {
	global.GVA_MODEL
	CompleteAt  int64  `json:"completeAt" form:"completeAt" gorm:"comment:完成时间;column:complete_at;size:19;"` //完成时间
	UserId      int64  `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;"`                     //用户ID
	TaskId      string `json:"taskId" form:"taskId" gorm:"comment:任务ID;column:task_id;"`                     //任务ID
	Options     string `json:"options" form:"options" gorm:"comment:任务参数;column:options;size:2048;"`         //任务参数
	Type        string `json:"type" form:"type" gorm:"comment:任务类型;column:type;size:48;"`                    //任务类型
	Status      string `json:"status" form:"status" gorm:"comment:任务状态;column:status;size:48;"`              //任务状态
	Result      string `json:"result" form:"result" gorm:"comment:任务结果;column:result;size:4096;"`
	Description string `json:"description" form:"description" gorm:"comment:备注说明;column:description;size:1024;"` //备注说明
}

// TableName extAiTask表 ExtAiTask自定义表名 ext_ai_task
func (ExtAiTask) TableName() string {
	return "ext_ai_task"
}
