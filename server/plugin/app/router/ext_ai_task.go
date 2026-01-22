package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var ExtAiTask = new(extAiTask)

type extAiTask struct {}

// Init 初始化 extAiTask表 路由信息
func (r *extAiTask) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("extAiTask").Use(middleware.OperationRecord())
		group.POST("createExtAiTask", apiExtAiTask.CreateExtAiTask)   // 新建extAiTask表
		group.DELETE("deleteExtAiTask", apiExtAiTask.DeleteExtAiTask) // 删除extAiTask表
		group.DELETE("deleteExtAiTaskByIds", apiExtAiTask.DeleteExtAiTaskByIds) // 批量删除extAiTask表
		group.PUT("updateExtAiTask", apiExtAiTask.UpdateExtAiTask)    // 更新extAiTask表
	}
	{
	    group := private.Group("extAiTask")
		group.GET("findExtAiTask", apiExtAiTask.FindExtAiTask)        // 根据ID获取extAiTask表
		group.GET("getExtAiTaskList", apiExtAiTask.GetExtAiTaskList)  // 获取extAiTask表列表
	}
	{
	    group := public.Group("extAiTask")
	    group.GET("getExtAiTaskPublic", apiExtAiTask.GetExtAiTaskPublic)  // extAiTask表开放接口
	}
}
