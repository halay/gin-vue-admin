import service from '@/utils/request'
// @Tags MerchantCategory
// @Summary 创建商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantCategory true "创建商户分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /MCAT/createMerchantCategory [post]
export const createMerchantCategory = (data) => {
  return service({
    url: '/MCAT/createMerchantCategory',
    method: 'post',
    data
  })
}

// @Tags MerchantCategory
// @Summary 删除商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantCategory true "删除商户分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MCAT/deleteMerchantCategory [delete]
export const deleteMerchantCategory = (params) => {
  return service({
    url: '/MCAT/deleteMerchantCategory',
    method: 'delete',
    params
  })
}

// @Tags MerchantCategory
// @Summary 批量删除商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商户分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MCAT/deleteMerchantCategory [delete]
export const deleteMerchantCategoryByIds = (params) => {
  return service({
    url: '/MCAT/deleteMerchantCategoryByIds',
    method: 'delete',
    params
  })
}

// @Tags MerchantCategory
// @Summary 更新商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantCategory true "更新商户分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MCAT/updateMerchantCategory [put]
export const updateMerchantCategory = (data) => {
  return service({
    url: '/MCAT/updateMerchantCategory',
    method: 'put',
    data
  })
}

// @Tags MerchantCategory
// @Summary 用id查询商户分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MerchantCategory true "用id查询商户分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MCAT/findMerchantCategory [get]
export const findMerchantCategory = (params) => {
  return service({
    url: '/MCAT/findMerchantCategory',
    method: 'get',
    params
  })
}

// @Tags MerchantCategory
// @Summary 分页获取商户分类列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商户分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MCAT/getMerchantCategoryList [get]
export const getMerchantCategoryList = (params) => {
  return service({
    url: '/MCAT/getMerchantCategoryList',
    method: 'get',
    params
  })
}
// @Tags MerchantCategory
// @Summary 不需要鉴权的商户分类接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantCategorySearch true "分页获取商户分类列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MCAT/getMerchantCategoryPublic [get]
export const getMerchantCategoryPublic = () => {
  return service({
    url: '/MCAT/getMerchantCategoryPublic',
    method: 'get',
  })
}
