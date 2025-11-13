import service from '@/utils/request'
// @Tags UserPointsLog
// @Summary 创建用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsLog true "创建用户积分流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /UPL/createUserPointsLog [post]
export const createUserPointsLog = (data) => {
  return service({
    url: '/UPL/createUserPointsLog',
    method: 'post',
    data
  })
}

// @Tags UserPointsLog
// @Summary 删除用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsLog true "删除用户积分流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /UPL/deleteUserPointsLog [delete]
export const deleteUserPointsLog = (params) => {
  return service({
    url: '/UPL/deleteUserPointsLog',
    method: 'delete',
    params
  })
}

// @Tags UserPointsLog
// @Summary 批量删除用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户积分流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /UPL/deleteUserPointsLog [delete]
export const deleteUserPointsLogByIds = (params) => {
  return service({
    url: '/UPL/deleteUserPointsLogByIds',
    method: 'delete',
    params
  })
}

// @Tags UserPointsLog
// @Summary 更新用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsLog true "更新用户积分流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /UPL/updateUserPointsLog [put]
export const updateUserPointsLog = (data) => {
  return service({
    url: '/UPL/updateUserPointsLog',
    method: 'put',
    data
  })
}

// @Tags UserPointsLog
// @Summary 用id查询用户积分流水
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserPointsLog true "用id查询用户积分流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /UPL/findUserPointsLog [get]
export const findUserPointsLog = (params) => {
  return service({
    url: '/UPL/findUserPointsLog',
    method: 'get',
    params
  })
}

// @Tags UserPointsLog
// @Summary 分页获取用户积分流水列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户积分流水列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /UPL/getUserPointsLogList [get]
export const getUserPointsLogList = (params) => {
  return service({
    url: '/UPL/getUserPointsLogList',
    method: 'get',
    params
  })
}
// @Tags UserPointsLog
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /UPL/findUserPointsLogDataSource [get]
export const getUserPointsLogDataSource = () => {
  return service({
    url: '/UPL/getUserPointsLogDataSource',
    method: 'get',
  })
}
// @Tags UserPointsLog
// @Summary 不需要鉴权的用户积分流水接口
// @Accept application/json
// @Produce application/json
// @Param data query request.UserPointsLogSearch true "分页获取用户积分流水列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /UPL/getUserPointsLogPublic [get]
export const getUserPointsLogPublic = () => {
  return service({
    url: '/UPL/getUserPointsLogPublic',
    method: 'get',
  })
}
