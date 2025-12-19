import service from '@/utils/request'
// @Tags MerchantPointsSettings
// @Summary 创建商户积分营销配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantPointsSettings true "创建商户积分营销配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /MPTS/createMerchantPointsSettings [post]
export const createMerchantPointsSettings = (data) => {
  return service({
    url: '/MPTS/createMerchantPointsSettings',
    method: 'post',
    data
  })
}

// @Tags MerchantPointsSettings
// @Summary 删除商户积分营销配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantPointsSettings true "删除商户积分营销配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MPTS/deleteMerchantPointsSettings [delete]
export const deleteMerchantPointsSettings = (params) => {
  return service({
    url: '/MPTS/deleteMerchantPointsSettings',
    method: 'delete',
    params
  })
}

// @Tags MerchantPointsSettings
// @Summary 批量删除商户积分营销配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商户积分营销配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MPTS/deleteMerchantPointsSettings [delete]
export const deleteMerchantPointsSettingsByIds = (params) => {
  return service({
    url: '/MPTS/deleteMerchantPointsSettingsByIds',
    method: 'delete',
    params
  })
}

// @Tags MerchantPointsSettings
// @Summary 更新商户积分营销配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantPointsSettings true "更新商户积分营销配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MPTS/updateMerchantPointsSettings [put]
export const updateMerchantPointsSettings = (data) => {
  return service({
    url: '/MPTS/updateMerchantPointsSettings',
    method: 'put',
    data
  })
}

// @Tags MerchantPointsSettings
// @Summary 用id查询商户积分营销配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MerchantPointsSettings true "用id查询商户积分营销配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MPTS/findMerchantPointsSettings [get]
export const findMerchantPointsSettings = (params) => {
  return service({
    url: '/MPTS/findMerchantPointsSettings',
    method: 'get',
    params
  })
}

// @Tags MerchantPointsSettings
// @Summary 分页获取商户积分营销配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商户积分营销配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MPTS/getMerchantPointsSettingsList [get]
export const getMerchantPointsSettingsList = (params) => {
  return service({
    url: '/MPTS/getMerchantPointsSettingsList',
    method: 'get',
    params
  })
}
// @Tags MerchantPointsSettings
// @Summary 不需要鉴权的商户积分营销配置接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantPointsSettingsSearch true "分页获取商户积分营销配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MPTS/getMerchantPointsSettingsPublic [get]
export const getMerchantPointsSettingsPublic = () => {
  return service({
    url: '/MPTS/getMerchantPointsSettingsPublic',
    method: 'get',
  })
}
