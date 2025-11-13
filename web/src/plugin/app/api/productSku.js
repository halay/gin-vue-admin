import service from '@/utils/request'
// @Tags ProductSku
// @Summary 创建商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSku true "创建商品SKU（规格组合）"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /SKU/createProductSku [post]
export const createProductSku = (data) => {
  return service({
    url: '/SKU/createProductSku',
    method: 'post',
    data
  })
}

// @Tags ProductSku
// @Summary 删除商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSku true "删除商品SKU（规格组合）"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SKU/deleteProductSku [delete]
export const deleteProductSku = (params) => {
  return service({
    url: '/SKU/deleteProductSku',
    method: 'delete',
    params
  })
}

// @Tags ProductSku
// @Summary 批量删除商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品SKU（规格组合）"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SKU/deleteProductSku [delete]
export const deleteProductSkuByIds = (params) => {
  return service({
    url: '/SKU/deleteProductSkuByIds',
    method: 'delete',
    params
  })
}

// @Tags ProductSku
// @Summary 更新商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSku true "更新商品SKU（规格组合）"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SKU/updateProductSku [put]
export const updateProductSku = (data) => {
  return service({
    url: '/SKU/updateProductSku',
    method: 'put',
    data
  })
}

// @Tags ProductSku
// @Summary 用id查询商品SKU（规格组合）
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ProductSku true "用id查询商品SKU（规格组合）"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SKU/findProductSku [get]
export const findProductSku = (params) => {
  return service({
    url: '/SKU/findProductSku',
    method: 'get',
    params
  })
}

// @Tags ProductSku
// @Summary 分页获取商品SKU（规格组合）列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品SKU（规格组合）列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SKU/getProductSkuList [get]
export const getProductSkuList = (params) => {
  return service({
    url: '/SKU/getProductSkuList',
    method: 'get',
    params
  })
}
// @Tags ProductSku
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SKU/findProductSkuDataSource [get]
export const getProductSkuDataSource = () => {
  return service({
    url: '/SKU/getProductSkuDataSource',
    method: 'get',
  })
}
// @Tags ProductSku
// @Summary 不需要鉴权的商品SKU（规格组合）接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductSkuSearch true "分页获取商品SKU（规格组合）列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /SKU/getProductSkuPublic [get]
export const getProductSkuPublic = () => {
  return service({
    url: '/SKU/getProductSkuPublic',
    method: 'get',
  })
}
