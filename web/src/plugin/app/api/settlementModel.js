import service from '@/utils/request'
// @Tags Settlement
// @Summary 创建结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Settlement true "创建结算管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /settlement/createSettlement [post]
export const createSettlement = (data) => {
  return service({
    url: '/settlement/createSettlement',
    method: 'post',
    data
  })
}

// @Tags Settlement
// @Summary 删除结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Settlement true "删除结算管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /settlement/deleteSettlement [delete]
export const deleteSettlement = (params) => {
  return service({
    url: '/settlement/deleteSettlement',
    method: 'delete',
    params
  })
}

// @Tags Settlement
// @Summary 批量删除结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除结算管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /settlement/deleteSettlement [delete]
export const deleteSettlementByIds = (params) => {
  return service({
    url: '/settlement/deleteSettlementByIds',
    method: 'delete',
    params
  })
}

// @Tags Settlement
// @Summary 更新结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Settlement true "更新结算管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /settlement/updateSettlement [put]
export const updateSettlement = (data) => {
  return service({
    url: '/settlement/updateSettlement',
    method: 'put',
    data
  })
}

// @Tags Settlement
// @Summary 用id查询结算管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Settlement true "用id查询结算管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /settlement/findSettlement [get]
export const findSettlement = (params) => {
  return service({
    url: '/settlement/findSettlement',
    method: 'get',
    params
  })
}

// @Tags Settlement
// @Summary 分页获取结算管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取结算管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /settlement/getSettlementList [get]
export const getSettlementList = (params) => {
  return service({
    url: '/settlement/getSettlementList',
    method: 'get',
    params
  })
}
// @Tags Settlement
// @Summary 不需要鉴权的结算管理接口
// @Accept application/json
// @Produce application/json
// @Param data query request.SettlementSearch true "分页获取结算管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /settlement/getSettlementPublic [get]
export const getSettlementPublic = () => {
  return service({
    url: '/settlement/getSettlementPublic',
    method: 'get',
  })
}
