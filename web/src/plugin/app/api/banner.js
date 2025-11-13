import service from '@/utils/request'
// @Tags Banner
// @Summary 创建Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Banner true "创建Banner图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /banner/createBanner [post]
export const createBanner = (data) => {
  return service({
    url: '/banner/createBanner',
    method: 'post',
    data
  })
}

// @Tags Banner
// @Summary 删除Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Banner true "删除Banner图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banner/deleteBanner [delete]
export const deleteBanner = (params) => {
  return service({
    url: '/banner/deleteBanner',
    method: 'delete',
    params
  })
}

// @Tags Banner
// @Summary 批量删除Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Banner图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banner/deleteBanner [delete]
export const deleteBannerByIds = (params) => {
  return service({
    url: '/banner/deleteBannerByIds',
    method: 'delete',
    params
  })
}

// @Tags Banner
// @Summary 更新Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Banner true "更新Banner图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /banner/updateBanner [put]
export const updateBanner = (data) => {
  return service({
    url: '/banner/updateBanner',
    method: 'put',
    data
  })
}

// @Tags Banner
// @Summary 用id查询Banner图片
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Banner true "用id查询Banner图片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /banner/findBanner [get]
export const findBanner = (params) => {
  return service({
    url: '/banner/findBanner',
    method: 'get',
    params
  })
}

// @Tags Banner
// @Summary 分页获取Banner图片列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Banner图片列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /banner/getBannerList [get]
export const getBannerList = (params) => {
  return service({
    url: '/banner/getBannerList',
    method: 'get',
    params
  })
}
// @Tags Banner
// @Summary 不需要鉴权的Banner图片接口
// @Accept application/json
// @Produce application/json
// @Param data query request.BannerSearch true "分页获取Banner图片列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /banner/getBannerPublic [get]
export const getBannerPublic = () => {
  return service({
    url: '/banner/getBannerPublic',
    method: 'get',
  })
}
