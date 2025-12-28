import service from '@/utils/request'
// @Tags AgentTransactionDetail
// @Summary 创建代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentTransactionDetail true "创建代理交易明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /atd/createAgentTransactionDetail [post]
export const createAgentTransactionDetail = (data) => {
  return service({
    url: '/atd/createAgentTransactionDetail',
    method: 'post',
    data
  })
}

// @Tags AgentTransactionDetail
// @Summary 删除代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentTransactionDetail true "删除代理交易明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /atd/deleteAgentTransactionDetail [delete]
export const deleteAgentTransactionDetail = (params) => {
  return service({
    url: '/atd/deleteAgentTransactionDetail',
    method: 'delete',
    params
  })
}

// @Tags AgentTransactionDetail
// @Summary 批量删除代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除代理交易明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /atd/deleteAgentTransactionDetail [delete]
export const deleteAgentTransactionDetailByIds = (params) => {
  return service({
    url: '/atd/deleteAgentTransactionDetailByIds',
    method: 'delete',
    params
  })
}

// @Tags AgentTransactionDetail
// @Summary 更新代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentTransactionDetail true "更新代理交易明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /atd/updateAgentTransactionDetail [put]
export const updateAgentTransactionDetail = (data) => {
  return service({
    url: '/atd/updateAgentTransactionDetail',
    method: 'put',
    data
  })
}

// @Tags AgentTransactionDetail
// @Summary 用id查询代理交易明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AgentTransactionDetail true "用id查询代理交易明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /atd/findAgentTransactionDetail [get]
export const findAgentTransactionDetail = (params) => {
  return service({
    url: '/atd/findAgentTransactionDetail',
    method: 'get',
    params
  })
}

// @Tags AgentTransactionDetail
// @Summary 分页获取代理交易明细列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取代理交易明细列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /atd/getAgentTransactionDetailList [get]
export const getAgentTransactionDetailList = (params) => {
  return service({
    url: '/atd/getAgentTransactionDetailList',
    method: 'get',
    params
  })
}
// @Tags AgentTransactionDetail
// @Summary 不需要鉴权的代理交易明细接口
// @Accept application/json
// @Produce application/json
// @Param data query request.AgentTransactionDetailSearch true "分页获取代理交易明细列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /atd/getAgentTransactionDetailPublic [get]
export const getAgentTransactionDetailPublic = () => {
  return service({
    url: '/atd/getAgentTransactionDetailPublic',
    method: 'get',
  })
}
