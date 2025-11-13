import service from '@/utils/request'
// @Tags ProductSpec
// @Summary 创建商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpec true "创建商品规格键"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PS/createProductSpec [post]
export const createProductSpec = (data) => {
  return service({
    url: '/PS/createProductSpec',
    method: 'post',
    data
  })
}

// @Tags ProductSpec
// @Summary 删除商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpec true "删除商品规格键"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PS/deleteProductSpec [delete]
export const deleteProductSpec = (params) => {
  return service({
    url: '/PS/deleteProductSpec',
    method: 'delete',
    params
  })
}

// @Tags ProductSpec
// @Summary 批量删除商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品规格键"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PS/deleteProductSpec [delete]
export const deleteProductSpecByIds = (params) => {
  return service({
    url: '/PS/deleteProductSpecByIds',
    method: 'delete',
    params
  })
}

// @Tags ProductSpec
// @Summary 更新商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpec true "更新商品规格键"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PS/updateProductSpec [put]
export const updateProductSpec = (data) => {
  return service({
    url: '/PS/updateProductSpec',
    method: 'put',
    data
  })
}

// @Tags ProductSpec
// @Summary 用id查询商品规格键
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ProductSpec true "用id查询商品规格键"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PS/findProductSpec [get]
export const findProductSpec = (params) => {
  return service({
    url: '/PS/findProductSpec',
    method: 'get',
    params
  })
}

// @Tags ProductSpec
// @Summary 分页获取商品规格键列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品规格键列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PS/getProductSpecList [get]
export const getProductSpecList = (params) => {
  return service({
    url: '/PS/getProductSpecList',
    method: 'get',
    params
  })
}
// @Tags ProductSpec
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PS/findProductSpecDataSource [get]
export const getProductSpecDataSource = () => {
  return service({
    url: '/PS/getProductSpecDataSource',
    method: 'get',
  })
}
// @Tags ProductSpec
// @Summary 不需要鉴权的商品规格键接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductSpecSearch true "分页获取商品规格键列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PS/getProductSpecPublic [get]
export const getProductSpecPublic = () => {
  return service({
    url: '/PS/getProductSpecPublic',
    method: 'get',
  })
}
