import service from '@/utils/request'
// @Tags AppRelease
// @Summary 创建App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppRelease true "创建App客户端版本升级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /AR/createAppRelease [post]
export const createAppRelease = (data) => {
  return service({
    url: '/AR/createAppRelease',
    method: 'post',
    data
  })
}

// @Tags AppRelease
// @Summary 删除App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppRelease true "删除App客户端版本升级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AR/deleteAppRelease [delete]
export const deleteAppRelease = (params) => {
  return service({
    url: '/AR/deleteAppRelease',
    method: 'delete',
    params
  })
}

// @Tags AppRelease
// @Summary 批量删除App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除App客户端版本升级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /AR/deleteAppRelease [delete]
export const deleteAppReleaseByIds = (params) => {
  return service({
    url: '/AR/deleteAppReleaseByIds',
    method: 'delete',
    params
  })
}

// @Tags AppRelease
// @Summary 更新App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AppRelease true "更新App客户端版本升级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /AR/updateAppRelease [put]
export const updateAppRelease = (data) => {
  return service({
    url: '/AR/updateAppRelease',
    method: 'put',
    data
  })
}

// @Tags AppRelease
// @Summary 用id查询App客户端版本升级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AppRelease true "用id查询App客户端版本升级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /AR/findAppRelease [get]
export const findAppRelease = (params) => {
  return service({
    url: '/AR/findAppRelease',
    method: 'get',
    params
  })
}

// @Tags AppRelease
// @Summary 分页获取App客户端版本升级列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取App客户端版本升级列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /AR/getAppReleaseList [get]
export const getAppReleaseList = (params) => {
  return service({
    url: '/AR/getAppReleaseList',
    method: 'get',
    params
  })
}
// @Tags AppRelease
// @Summary 不需要鉴权的App客户端版本升级接口
// @Accept application/json
// @Produce application/json
// @Param data query request.AppReleaseSearch true "分页获取App客户端版本升级列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /AR/getAppReleasePublic [get]
export const getAppReleasePublic = () => {
  return service({
    url: '/AR/getAppReleasePublic',
    method: 'get',
  })
}
