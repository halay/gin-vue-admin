package api

import (
	"net/http"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var ExtAiTask = new(extAiTask)

type extAiTask struct{}

// CreateExtAiTask 创建extAiTask表
// @Tags ExtAiTask
// @Summary 创建extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ExtAiTask true "创建extAiTask表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /extAiTask/createExtAiTask [post]
func (a *extAiTask) CreateExtAiTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.ExtAiTask
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceExtAiTask.CreateExtAiTask(ctx, &info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExtAiTask 删除extAiTask表
// @Tags ExtAiTask
// @Summary 删除extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ExtAiTask true "删除extAiTask表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /extAiTask/deleteExtAiTask [delete]
func (a *extAiTask) DeleteExtAiTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := serviceExtAiTask.DeleteExtAiTask(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteExtAiTaskByIds 批量删除extAiTask表
// @Tags ExtAiTask
// @Summary 批量删除extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /extAiTask/deleteExtAiTaskByIds [delete]
func (a *extAiTask) DeleteExtAiTaskByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := serviceExtAiTask.DeleteExtAiTaskByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateExtAiTask 更新extAiTask表
// @Tags ExtAiTask
// @Summary 更新extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ExtAiTask true "更新extAiTask表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /extAiTask/updateExtAiTask [put]
func (a *extAiTask) UpdateExtAiTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var info model.ExtAiTask
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceExtAiTask.UpdateExtAiTask(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindExtAiTask 用id查询extAiTask表
// @Tags ExtAiTask
// @Summary 用id查询extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询extAiTask表"
// @Success 200 {object} response.Response{data=model.ExtAiTask,msg=string} "查询成功"
// @Router /extAiTask/findExtAiTask [get]
func (a *extAiTask) FindExtAiTask(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reextAiTask, err := serviceExtAiTask.GetExtAiTask(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reextAiTask, c)
}

// GetExtAiTaskList 分页获取extAiTask表列表
// @Tags ExtAiTask
// @Summary 分页获取extAiTask表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.ExtAiTaskSearch true "分页获取extAiTask表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /extAiTask/getExtAiTaskList [get]
func (a *extAiTask) GetExtAiTaskList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo request.ExtAiTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = strconv.Itoa(int(utils.GetUserID(c)))
	list, total, err := serviceExtAiTask.GetExtAiTaskInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetExtAiTaskPublic 不需要鉴权的extAiTask表接口
// @Tags ExtAiTask
// @Summary 不需要鉴权的extAiTask表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /extAiTask/getExtAiTaskPublic [get]
func (a *extAiTask) GetExtAiTaskPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	serviceExtAiTask.GetExtAiTaskPublic(ctx)
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的extAiTask表接口信息"}, "获取成功", c)
}

func (a *extAiTask) GetCozeFile(c *gin.Context) {
	cozeId := c.Query("id")
	if cozeId == "" {
		response.FailWithMessage("missing coze id", c)
		return
	}
	filename := "coze-" + cozeId
	var fileModel example.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("name LIKE ?", filename+"%").First(&fileModel).Error
	if err != nil {
		response.FailWithMessage("file not found", c)
		return
	}
	c.Redirect(http.StatusFound, fileModel.Url)
}
