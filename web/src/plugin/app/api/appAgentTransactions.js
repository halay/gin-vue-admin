import service from '@/utils/request'
// @Tags AppAgentTransactions
// @Summary 创建appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppAgentTransactions true "创建appAgentTransactions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /appAgentTransactions/createAppAgentTransactions [post]
export const createAppAgentTransactions = (data) => {
  return service({
    url: '/appAgentTransactions/createAppAgentTransactions',
    method: 'post',
    data
  })
}

// @Tags AppAgentTransactions
// @Summary 删除appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppAgentTransactions true "删除appAgentTransactions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAgentTransactions/deleteAppAgentTransactions [delete]
export const deleteAppAgentTransactions = (params) => {
  return service({
    url: '/appAgentTransactions/deleteAppAgentTransactions',
    method: 'delete',
    params
  })
}

// @Tags AppAgentTransactions
// @Summary 批量删除appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除appAgentTransactions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appAgentTransactions/deleteAppAgentTransactions [delete]
export const deleteAppAgentTransactionsByIds = (params) => {
  return service({
    url: '/appAgentTransactions/deleteAppAgentTransactionsByIds',
    method: 'delete',
    params
  })
}

// @Tags AppAgentTransactions
// @Summary 更新appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppAgentTransactions true "更新appAgentTransactions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appAgentTransactions/updateAppAgentTransactions [put]
export const updateAppAgentTransactions = (data) => {
  return service({
    url: '/appAgentTransactions/updateAppAgentTransactions',
    method: 'put',
    data
  })
}

// @Tags AppAgentTransactions
// @Summary 用id查询appAgentTransactions表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AppAgentTransactions true "用id查询appAgentTransactions表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appAgentTransactions/findAppAgentTransactions [get]
export const findAppAgentTransactions = (params) => {
  return service({
    url: '/appAgentTransactions/findAppAgentTransactions',
    method: 'get',
    params
  })
}

// @Tags AppAgentTransactions
// @Summary 分页获取appAgentTransactions表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取appAgentTransactions表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appAgentTransactions/getAppAgentTransactionsList [get]
export const getAppAgentTransactionsList = (params) => {
  return service({
    url: '/appAgentTransactions/getAppAgentTransactionsList',
    method: 'get',
    params
  })
}
// @Tags AppAgentTransactions
// @Summary 不需要鉴权的appAgentTransactions表接口
// @Accept application/json
// @Produce application/json
// @Param data query request.AppAgentTransactionsSearch true "分页获取appAgentTransactions表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appAgentTransactions/getAppAgentTransactionsPublic [get]
export const getAppAgentTransactionsPublic = () => {
  return service({
    url: '/appAgentTransactions/getAppAgentTransactionsPublic',
    method: 'get',
  })
}
