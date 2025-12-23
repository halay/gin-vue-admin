import service from '@/utils/request'
// @Tags AgentLevel
// @Summary 创建营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentLevel true "创建营销方案设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /AL/createAgentLevel [post]
export const createAgentLevel = (data) => {
  return service({
    url: '/AL/createAgentLevel',
    method: 'post',
    data
  })
}

// @Tags AgentLevel
// @Summary 删除营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentLevel true "删除营销方案设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AL/deleteAgentLevel [delete]
export const deleteAgentLevel = (params) => {
  return service({
    url: '/AL/deleteAgentLevel',
    method: 'delete',
    params
  })
}

// @Tags AgentLevel
// @Summary 批量删除营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除营销方案设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AL/deleteAgentLevel [delete]
export const deleteAgentLevelByIds = (params) => {
  return service({
    url: '/AL/deleteAgentLevelByIds',
    method: 'delete',
    params
  })
}

// @Tags AgentLevel
// @Summary 更新营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentLevel true "更新营销方案设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /AL/updateAgentLevel [put]
export const updateAgentLevel = (data) => {
  return service({
    url: '/AL/updateAgentLevel',
    method: 'put',
    data
  })
}

// @Tags AgentLevel
// @Summary 用id查询营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AgentLevel true "用id查询营销方案设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /AL/findAgentLevel [get]
export const findAgentLevel = (params) => {
  return service({
    url: '/AL/findAgentLevel',
    method: 'get',
    params
  })
}

// @Tags AgentLevel
// @Summary 分页获取营销方案设置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取营销方案设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /AL/getAgentLevelList [get]
export const getAgentLevelList = (params) => {
  return service({
    url: '/AL/getAgentLevelList',
    method: 'get',
    params
  })
}
// @Tags AgentLevel
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /AL/findAgentLevelDataSource [get]
export const getAgentLevelDataSource = () => {
  return service({
    url: '/AL/getAgentLevelDataSource',
    method: 'get',
  })
}
// @Tags AgentLevel
// @Summary 不需要鉴权的营销方案设置接口
// @Accept application/json
// @Produce application/json
// @Param data query request.AgentLevelSearch true "分页获取营销方案设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /AL/getAgentLevelPublic [get]
export const getAgentLevelPublic = () => {
  return service({
    url: '/AL/getAgentLevelPublic',
    method: 'get',
  })
}
