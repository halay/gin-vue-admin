import service from '@/utils/request'
// @Tags UserPointsAccount
// @Summary 创建用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsAccount true "创建用户积分账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /UPA/createUserPointsAccount [post]
export const createUserPointsAccount = (data) => {
  return service({
    url: '/UPA/createUserPointsAccount',
    method: 'post',
    data
  })
}

// @Tags UserPointsAccount
// @Summary 删除用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsAccount true "删除用户积分账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /UPA/deleteUserPointsAccount [delete]
export const deleteUserPointsAccount = (params) => {
  return service({
    url: '/UPA/deleteUserPointsAccount',
    method: 'delete',
    params
  })
}

// @Tags UserPointsAccount
// @Summary 批量删除用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户积分账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /UPA/deleteUserPointsAccount [delete]
export const deleteUserPointsAccountByIds = (params) => {
  return service({
    url: '/UPA/deleteUserPointsAccountByIds',
    method: 'delete',
    params
  })
}

// @Tags UserPointsAccount
// @Summary 更新用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.UserPointsAccount true "更新用户积分账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /UPA/updateUserPointsAccount [put]
export const updateUserPointsAccount = (data) => {
  return service({
    url: '/UPA/updateUserPointsAccount',
    method: 'put',
    data
  })
}

// @Tags UserPointsAccount
// @Summary 用id查询用户积分账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.UserPointsAccount true "用id查询用户积分账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /UPA/findUserPointsAccount [get]
export const findUserPointsAccount = (params) => {
  return service({
    url: '/UPA/findUserPointsAccount',
    method: 'get',
    params
  })
}

// @Tags UserPointsAccount
// @Summary 分页获取用户积分账户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户积分账户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /UPA/getUserPointsAccountList [get]
export const getUserPointsAccountList = (params) => {
  return service({
    url: '/UPA/getUserPointsAccountList',
    method: 'get',
    params
  })
}
// @Tags UserPointsAccount
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /UPA/findUserPointsAccountDataSource [get]
export const getUserPointsAccountDataSource = () => {
  return service({
    url: '/UPA/getUserPointsAccountDataSource',
    method: 'get',
  })
}
// @Tags UserPointsAccount
// @Summary 不需要鉴权的用户积分账户接口
// @Accept application/json
// @Produce application/json
// @Param data query request.UserPointsAccountSearch true "分页获取用户积分账户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /UPA/getUserPointsAccountPublic [get]
export const getUserPointsAccountPublic = () => {
  return service({
    url: '/UPA/getUserPointsAccountPublic',
    method: 'get',
  })
}
