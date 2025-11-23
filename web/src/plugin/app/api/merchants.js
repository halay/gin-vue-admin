import service from '@/utils/request'
// @Tags Merchants
// @Summary 创建商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Merchants true "创建商家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mc/createMerchants [post]
export const createMerchants = (data) => {
  return service({
    url: '/mc/createMerchants',
    method: 'post',
    data
  })
}

// @Tags Merchants
// @Summary 删除商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Merchants true "删除商家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mc/deleteMerchants [delete]
export const deleteMerchants = (params) => {
  return service({
    url: '/mc/deleteMerchants',
    method: 'delete',
    params
  })
}

// @Tags Merchants
// @Summary 批量删除商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mc/deleteMerchants [delete]
export const deleteMerchantsByIds = (params) => {
  return service({
    url: '/mc/deleteMerchantsByIds',
    method: 'delete',
    params
  })
}

// @Tags Merchants
// @Summary 更新商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Merchants true "更新商家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mc/updateMerchants [put]
export const updateMerchants = (data) => {
  return service({
    url: '/mc/updateMerchants',
    method: 'put',
    data
  })
}

// @Tags Merchants
// @Summary 用id查询商家信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Merchants true "用id查询商家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mc/findMerchants [get]
export const findMerchants = (params) => {
  return service({
    url: '/mc/findMerchants',
    method: 'get',
    params
  })
}

// @Tags Merchants
// @Summary 分页获取商家信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商家信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mc/getMerchantsList [get]
export const getMerchantsList = (params) => {
  return service({
    url: '/mc/getMerchantsList',
    method: 'get',
    params
  })
}
// @Tags Merchants
// @Summary 不需要鉴权的商家信息接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantsSearch true "分页获取商家信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /mc/getMerchantsPublic [get]
export const getMerchantsPublic = () => {
  return service({
    url: '/mc/getMerchantsPublic',
    method: 'get',
  })
}

export const getRecommendedMerchants = () => {
  return service({ url: '/mc/getRecommendedMerchants', method: 'get' })
}

export const getMerchantCategoryList = () => {
  return service({ url: '/mc/getMerchantCategoryList', method: 'get' })
}

export const getMerchantsByCategory = (params) => {
  return service({ url: '/mc/getMerchantsByCategory', method: 'get', params })
}
