import service from '@/utils/request'
// @Tags PointsSettings
// @Summary 创建平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PointsSettings true "创建平台积分基础设置与奖励规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PTS/createPointsSettings [post]
export const createPointsSettings = (data) => {
  return service({
    url: '/PTS/createPointsSettings',
    method: 'post',
    data
  })
}

// @Tags PointsSettings
// @Summary 删除平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PointsSettings true "删除平台积分基础设置与奖励规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PTS/deletePointsSettings [delete]
export const deletePointsSettings = (params) => {
  return service({
    url: '/PTS/deletePointsSettings',
    method: 'delete',
    params
  })
}

// @Tags PointsSettings
// @Summary 批量删除平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除平台积分基础设置与奖励规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PTS/deletePointsSettings [delete]
export const deletePointsSettingsByIds = (params) => {
  return service({
    url: '/PTS/deletePointsSettingsByIds',
    method: 'delete',
    params
  })
}

// @Tags PointsSettings
// @Summary 更新平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.PointsSettings true "更新平台积分基础设置与奖励规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PTS/updatePointsSettings [put]
export const updatePointsSettings = (data) => {
  return service({
    url: '/PTS/updatePointsSettings',
    method: 'put',
    data
  })
}

// @Tags PointsSettings
// @Summary 用id查询平台积分基础设置与奖励规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.PointsSettings true "用id查询平台积分基础设置与奖励规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PTS/findPointsSettings [get]
export const findPointsSettings = (params) => {
  return service({
    url: '/PTS/findPointsSettings',
    method: 'get',
    params
  })
}

// @Tags PointsSettings
// @Summary 分页获取平台积分基础设置与奖励规则列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取平台积分基础设置与奖励规则列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PTS/getPointsSettingsList [get]
export const getPointsSettingsList = (params) => {
  return service({
    url: '/PTS/getPointsSettingsList',
    method: 'get',
    params
  })
}
// @Tags PointsSettings
// @Summary 不需要鉴权的平台积分基础设置与奖励规则接口
// @Accept application/json
// @Produce application/json
// @Param data query request.PointsSettingsSearch true "分页获取平台积分基础设置与奖励规则列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PTS/getPointsSettingsPublic [get]
export const getPointsSettingsPublic = () => {
  return service({
    url: '/PTS/getPointsSettingsPublic',
    method: 'get',
  })
}
