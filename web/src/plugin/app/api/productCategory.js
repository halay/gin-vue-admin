import service from '@/utils/request'
// @Tags ProductCategory
// @Summary 创建商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductCategory true "创建商户商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PC/createProductCategory [post]
export const createProductCategory = (data) => {
  return service({
    url: '/PC/createProductCategory',
    method: 'post',
    data
  })
}

// @Tags ProductCategory
// @Summary 删除商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductCategory true "删除商户商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PC/deleteProductCategory [delete]
export const deleteProductCategory = (params) => {
  return service({
    url: '/PC/deleteProductCategory',
    method: 'delete',
    params
  })
}

// @Tags ProductCategory
// @Summary 批量删除商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商户商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PC/deleteProductCategory [delete]
export const deleteProductCategoryByIds = (params) => {
  return service({
    url: '/PC/deleteProductCategoryByIds',
    method: 'delete',
    params
  })
}

// @Tags ProductCategory
// @Summary 更新商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductCategory true "更新商户商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PC/updateProductCategory [put]
export const updateProductCategory = (data) => {
  return service({
    url: '/PC/updateProductCategory',
    method: 'put',
    data
  })
}

// @Tags ProductCategory
// @Summary 用id查询商户商品分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ProductCategory true "用id查询商户商品分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PC/findProductCategory [get]
export const findProductCategory = (params) => {
  return service({
    url: '/PC/findProductCategory',
    method: 'get',
    params
  })
}

// @Tags ProductCategory
// @Summary 分页获取商户商品分类列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商户商品分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PC/getProductCategoryList [get]
export const getProductCategoryList = (params) => {
  return service({
    url: '/PC/getProductCategoryList',
    method: 'get',
    params
  })
}
// @Tags ProductCategory
// @Summary 不需要鉴权的商户商品分类接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductCategorySearch true "分页获取商户商品分类列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PC/getProductCategoryPublic [get]
export const getProductCategoryPublic = () => {
  return service({
    url: '/PC/getProductCategoryPublic',
    method: 'get',
  })
}
