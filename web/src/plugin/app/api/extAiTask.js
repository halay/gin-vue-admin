import service from '@/utils/request'
// @Tags ExtAiTask
// @Summary 创建extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ExtAiTask true "创建extAiTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /extAiTask/createExtAiTask [post]
export const createExtAiTask = (data) => {
  return service({
    url: '/extAiTask/createExtAiTask',
    method: 'post',
    data
  })
}

// @Tags ExtAiTask
// @Summary 删除extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ExtAiTask true "删除extAiTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /extAiTask/deleteExtAiTask [delete]
export const deleteExtAiTask = (params) => {
  return service({
    url: '/extAiTask/deleteExtAiTask',
    method: 'delete',
    params
  })
}

// @Tags ExtAiTask
// @Summary 批量删除extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除extAiTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /extAiTask/deleteExtAiTask [delete]
export const deleteExtAiTaskByIds = (params) => {
  return service({
    url: '/extAiTask/deleteExtAiTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags ExtAiTask
// @Summary 更新extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ExtAiTask true "更新extAiTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /extAiTask/updateExtAiTask [put]
export const updateExtAiTask = (data) => {
  return service({
    url: '/extAiTask/updateExtAiTask',
    method: 'put',
    data
  })
}

// @Tags ExtAiTask
// @Summary 用id查询extAiTask表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ExtAiTask true "用id查询extAiTask表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /extAiTask/findExtAiTask [get]
export const findExtAiTask = (params) => {
  return service({
    url: '/extAiTask/findExtAiTask',
    method: 'get',
    params
  })
}

// @Tags ExtAiTask
// @Summary 分页获取extAiTask表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取extAiTask表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /extAiTask/getExtAiTaskList [get]
export const getExtAiTaskList = (params) => {
  return service({
    url: '/extAiTask/getExtAiTaskList',
    method: 'get',
    params
  })
}
// @Tags ExtAiTask
// @Summary 不需要鉴权的extAiTask表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ExtAiTaskSearch true "分页获取extAiTask表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /extAiTask/getExtAiTaskPublic [get]
export const getExtAiTaskPublic = () => {
  return service({
    url: '/extAiTask/getExtAiTaskPublic',
    method: 'get',
  })
}
