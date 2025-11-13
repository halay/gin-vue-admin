import service from '@/utils/request'
// @Tags MerchantAdmin
// @Summary 创建后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAdmin true "创建后台用户与商户关联映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /MADM/createMerchantAdmin [post]
export const createMerchantAdmin = (data) => {
  return service({
    url: '/MADM/createMerchantAdmin',
    method: 'post',
    data
  })
}

// @Tags MerchantAdmin
// @Summary 删除后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAdmin true "删除后台用户与商户关联映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MADM/deleteMerchantAdmin [delete]
export const deleteMerchantAdmin = (params) => {
  return service({
    url: '/MADM/deleteMerchantAdmin',
    method: 'delete',
    params
  })
}

// @Tags MerchantAdmin
// @Summary 批量删除后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除后台用户与商户关联映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MADM/deleteMerchantAdmin [delete]
export const deleteMerchantAdminByIds = (params) => {
  return service({
    url: '/MADM/deleteMerchantAdminByIds',
    method: 'delete',
    params
  })
}

// @Tags MerchantAdmin
// @Summary 更新后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAdmin true "更新后台用户与商户关联映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MADM/updateMerchantAdmin [put]
export const updateMerchantAdmin = (data) => {
  return service({
    url: '/MADM/updateMerchantAdmin',
    method: 'put',
    data
  })
}

// @Tags MerchantAdmin
// @Summary 用id查询后台用户与商户关联映射
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MerchantAdmin true "用id查询后台用户与商户关联映射"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MADM/findMerchantAdmin [get]
export const findMerchantAdmin = (params) => {
  return service({
    url: '/MADM/findMerchantAdmin',
    method: 'get',
    params
  })
}

// @Tags MerchantAdmin
// @Summary 分页获取后台用户与商户关联映射列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取后台用户与商户关联映射列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MADM/getMerchantAdminList [get]
export const getMerchantAdminList = (params) => {
  return service({
    url: '/MADM/getMerchantAdminList',
    method: 'get',
    params
  })
}
// @Tags MerchantAdmin
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MADM/findMerchantAdminDataSource [get]
export const getMerchantAdminDataSource = () => {
  return service({
    url: '/MADM/getMerchantAdminDataSource',
    method: 'get',
  })
}
// @Tags MerchantAdmin
// @Summary 不需要鉴权的后台用户与商户关联映射接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantAdminSearch true "分页获取后台用户与商户关联映射列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MADM/getMerchantAdminPublic [get]
export const getMerchantAdminPublic = () => {
  return service({
    url: '/MADM/getMerchantAdminPublic',
    method: 'get',
  })
}
