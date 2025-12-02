package api

import (

	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var Node = new(NODE)

type NODE struct {}

// CreateNode 创建节点策略
// @Tags Node
// @Summary 创建节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Node true "创建节点策略"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /NODE/createNode [post]
func (a *NODE) CreateNode(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Node
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceNode.CreateNode(ctx,&info)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteNode 删除节点策略
// @Tags Node
// @Summary 删除节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Node true "删除节点策略"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /NODE/deleteNode [delete]
func (a *NODE) DeleteNode(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := serviceNode.DeleteNode(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("删除成功", c)
}

// DeleteNodeByIds 批量删除节点策略
// @Tags Node
// @Summary 批量删除节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /NODE/deleteNodeByIds [delete]
func (a *NODE) DeleteNodeByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := serviceNode.DeleteNodeByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("批量删除成功", c)
}

// UpdateNode 更新节点策略
// @Tags Node
// @Summary 更新节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Node true "更新节点策略"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /NODE/updateNode [put]
func (a *NODE) UpdateNode(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var info model.Node
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceNode.UpdateNode(ctx,info)
    if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("更新成功", c)
}

// FindNode 用id查询节点策略
// @Tags Node
// @Summary 用id查询节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询节点策略"
// @Success 200 {object} response.Response{data=model.Node,msg=string} "查询成功"
// @Router /NODE/findNode [get]
func (a *NODE) FindNode(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reNODE, err := serviceNode.GetNode(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
    response.OkWithData(reNODE, c)
}
// GetNodeList 分页获取节点策略列表
// @Tags Node
// @Summary 分页获取节点策略列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.NodeSearch true "分页获取节点策略列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /NODE/getNodeList [get]
func (a *NODE) GetNodeList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo request.NodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceNode.GetNodeInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetNodePublic 不需要鉴权的节点策略接口
// @Tags Node
// @Summary 不需要鉴权的节点策略接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /NODE/getNodePublic [get]
func (a *NODE) GetNodePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serviceNode.GetNodePublic(ctx)
    response.OkWithDetailed(gin.H{"info": "不需要鉴权的节点策略接口信息"}, "获取成功", c)
}
