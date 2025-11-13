import service from '@/utils/request'
// @Tags Product
// @Summary 创建商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "创建商户商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /P/createProduct [post]
export const createProduct = (data) => {
  return service({
    url: '/P/createProduct',
    method: 'post',
    data
  })
}

// @Tags Product
// @Summary 删除商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "删除商户商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /P/deleteProduct [delete]
export const deleteProduct = (params) => {
  return service({
    url: '/P/deleteProduct',
    method: 'delete',
    params
  })
}

// @Tags Product
// @Summary 批量删除商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商户商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /P/deleteProduct [delete]
export const deleteProductByIds = (params) => {
  return service({
    url: '/P/deleteProductByIds',
    method: 'delete',
    params
  })
}

// @Tags Product
// @Summary 更新商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Product true "更新商户商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /P/updateProduct [put]
export const updateProduct = (data) => {
  return service({
    url: '/P/updateProduct',
    method: 'put',
    data
  })
}

// @Tags Product
// @Summary 用id查询商户商品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Product true "用id查询商户商品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /P/findProduct [get]
export const findProduct = (params) => {
  return service({
    url: '/P/findProduct',
    method: 'get',
    params
  })
}

// @Tags Product
// @Summary 分页获取商户商品列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商户商品列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /P/getProductList [get]
export const getProductList = (params) => {
  return service({
    url: '/P/getProductList',
    method: 'get',
    params
  })
}
// @Tags Product
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /P/findProductDataSource [get]
export const getProductDataSource = () => {
  return service({
    url: '/P/getProductDataSource',
    method: 'get',
  })
}
// @Tags Product
// @Summary 不需要鉴权的商户商品接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductSearch true "分页获取商户商品列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /P/getProductPublic [get]
export const getProductPublic = () => {
  return service({
    url: '/P/getProductPublic',
    method: 'get',
  })
}
