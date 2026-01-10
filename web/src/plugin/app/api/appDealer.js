import service from '@/utils/request'
// @Tags AppDealer
// @Summary 创建经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppDealer true "创建经销商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ADL/createAppDealer [post]
export const createAppDealer = (data) => {
  return service({
    url: '/ADL/createAppDealer',
    method: 'post',
    data
  })
}

// @Tags AppDealer
// @Summary 删除经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppDealer true "删除经销商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ADL/deleteAppDealer [delete]
export const deleteAppDealer = (params) => {
  return service({
    url: '/ADL/deleteAppDealer',
    method: 'delete',
    params
  })
}

// @Tags AppDealer
// @Summary 批量删除经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除经销商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ADL/deleteAppDealer [delete]
export const deleteAppDealerByIds = (params) => {
  return service({
    url: '/ADL/deleteAppDealerByIds',
    method: 'delete',
    params
  })
}

// @Tags AppDealer
// @Summary 更新经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppDealer true "更新经销商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ADL/updateAppDealer [put]
export const updateAppDealer = (data) => {
  return service({
    url: '/ADL/updateAppDealer',
    method: 'put',
    data
  })
}

// @Tags AppDealer
// @Summary 用id查询经销商管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AppDealer true "用id查询经销商管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ADL/findAppDealer [get]
export const findAppDealer = (params) => {
  return service({
    url: '/ADL/findAppDealer',
    method: 'get',
    params
  })
}

// @Tags AppDealer
// @Summary 分页获取经销商管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取经销商管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ADL/getAppDealerList [get]
export const getAppDealerList = (params) => {
  return service({
    url: '/ADL/getAppDealerList',
    method: 'get',
    params
  })
}
// @Tags AppDealer
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ADL/findAppDealerDataSource [get]
export const getAppDealerDataSource = () => {
  return service({
    url: '/ADL/getAppDealerDataSource',
    method: 'get',
  })
}
// @Tags AppDealer
// @Summary 不需要鉴权的经销商管理接口
// @Accept application/json
// @Produce application/json
// @Param data query request.AppDealerSearch true "分页获取经销商管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ADL/getAppDealerPublic [get]
export const getAppDealerPublic = () => {
  return service({
    url: '/ADL/getAppDealerPublic',
    method: 'get',
  })
}
