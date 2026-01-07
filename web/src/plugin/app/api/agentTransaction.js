import service from '@/utils/request'

// @Tags AgentTransaction
// @Summary 分页获取代理分润记录列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.AgentTransactionSearch true "分页获取代理分润记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /agentTransaction/getAgentTransactionList [get]
export const getAgentTransactionList = (params) => {
  return service({
    url: '/agentTransaction/getAgentTransactionList',
    method: 'get',
    params
  })
}
