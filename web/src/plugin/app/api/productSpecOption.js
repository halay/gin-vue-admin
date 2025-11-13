import service from '@/utils/request'
// @Tags ProductSpecOption
// @Summary 创建商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpecOption true "创建商品规格值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PSO/createProductSpecOption [post]
export const createProductSpecOption = (data) => {
  return service({
    url: '/PSO/createProductSpecOption',
    method: 'post',
    data
  })
}

// @Tags ProductSpecOption
// @Summary 删除商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpecOption true "删除商品规格值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PSO/deleteProductSpecOption [delete]
export const deleteProductSpecOption = (params) => {
  return service({
    url: '/PSO/deleteProductSpecOption',
    method: 'delete',
    params
  })
}

// @Tags ProductSpecOption
// @Summary 批量删除商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商品规格值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PSO/deleteProductSpecOption [delete]
export const deleteProductSpecOptionByIds = (params) => {
  return service({
    url: '/PSO/deleteProductSpecOptionByIds',
    method: 'delete',
    params
  })
}

// @Tags ProductSpecOption
// @Summary 更新商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ProductSpecOption true "更新商品规格值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PSO/updateProductSpecOption [put]
export const updateProductSpecOption = (data) => {
  return service({
    url: '/PSO/updateProductSpecOption',
    method: 'put',
    data
  })
}

// @Tags ProductSpecOption
// @Summary 用id查询商品规格值
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ProductSpecOption true "用id查询商品规格值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PSO/findProductSpecOption [get]
export const findProductSpecOption = (params) => {
  return service({
    url: '/PSO/findProductSpecOption',
    method: 'get',
    params
  })
}

// @Tags ProductSpecOption
// @Summary 分页获取商品规格值列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商品规格值列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PSO/getProductSpecOptionList [get]
export const getProductSpecOptionList = (params) => {
  return service({
    url: '/PSO/getProductSpecOptionList',
    method: 'get',
    params
  })
}
// @Tags ProductSpecOption
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PSO/findProductSpecOptionDataSource [get]
export const getProductSpecOptionDataSource = () => {
  return service({
    url: '/PSO/getProductSpecOptionDataSource',
    method: 'get',
  })
}
// @Tags ProductSpecOption
// @Summary 不需要鉴权的商品规格值接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ProductSpecOptionSearch true "分页获取商品规格值列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PSO/getProductSpecOptionPublic [get]
export const getProductSpecOptionPublic = () => {
  return service({
    url: '/PSO/getProductSpecOptionPublic',
    method: 'get',
  })
}
