import service from '@/utils/request'
// @Tags Node
// @Summary 创建节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Node true "创建节点策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /NODE/createNode [post]
export const createNode = (data) => {
  return service({
    url: '/NODE/createNode',
    method: 'post',
    data
  })
}

// @Tags Node
// @Summary 删除节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Node true "删除节点策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NODE/deleteNode [delete]
export const deleteNode = (params) => {
  return service({
    url: '/NODE/deleteNode',
    method: 'delete',
    params
  })
}

// @Tags Node
// @Summary 批量删除节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除节点策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NODE/deleteNode [delete]
export const deleteNodeByIds = (params) => {
  return service({
    url: '/NODE/deleteNodeByIds',
    method: 'delete',
    params
  })
}

// @Tags Node
// @Summary 更新节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Node true "更新节点策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NODE/updateNode [put]
export const updateNode = (data) => {
  return service({
    url: '/NODE/updateNode',
    method: 'put',
    data
  })
}

// @Tags Node
// @Summary 用id查询节点策略
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Node true "用id查询节点策略"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /NODE/findNode [get]
export const findNode = (params) => {
  return service({
    url: '/NODE/findNode',
    method: 'get',
    params
  })
}

// @Tags Node
// @Summary 分页获取节点策略列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取节点策略列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NODE/getNodeList [get]
export const getNodeList = (params) => {
  return service({
    url: '/NODE/getNodeList',
    method: 'get',
    params
  })
}
// @Tags Node
// @Summary 不需要鉴权的节点策略接口
// @Accept application/json
// @Produce application/json
// @Param data query request.NodeSearch true "分页获取节点策略列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /NODE/getNodePublic [get]
export const getNodePublic = () => {
  return service({
    url: '/NODE/getNodePublic',
    method: 'get',
  })
}
