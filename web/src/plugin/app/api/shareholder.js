import service from '@/utils/request'
// @Tags ShareholderProfit
// @Summary 创建股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ShareholderProfit true "创建股东分润"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /spr/createShareholderProfit [post]
export const createShareholderProfit = (data) => {
  return service({
    url: '/spr/createShareholderProfit',
    method: 'post',
    data
  })
}

// @Tags ShareholderProfit
// @Summary 删除股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ShareholderProfit true "删除股东分润"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /spr/deleteShareholderProfit [delete]
export const deleteShareholderProfit = (params) => {
  return service({
    url: '/spr/deleteShareholderProfit',
    method: 'delete',
    params
  })
}

// @Tags ShareholderProfit
// @Summary 批量删除股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除股东分润"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /spr/deleteShareholderProfit [delete]
export const deleteShareholderProfitByIds = (params) => {
  return service({
    url: '/spr/deleteShareholderProfitByIds',
    method: 'delete',
    params
  })
}

// @Tags ShareholderProfit
// @Summary 更新股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ShareholderProfit true "更新股东分润"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /spr/updateShareholderProfit [put]
export const updateShareholderProfit = (data) => {
  return service({
    url: '/spr/updateShareholderProfit',
    method: 'put',
    data
  })
}

// @Tags ShareholderProfit
// @Summary 用id查询股东分润
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ShareholderProfit true "用id查询股东分润"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /spr/findShareholderProfit [get]
export const findShareholderProfit = (params) => {
  return service({
    url: '/spr/findShareholderProfit',
    method: 'get',
    params
  })
}

// @Tags ShareholderProfit
// @Summary 分页获取股东分润列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取股东分润列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spr/getShareholderProfitList [get]
export const getShareholderProfitList = (params) => {
  return service({
    url: '/spr/getShareholderProfitList',
    method: 'get',
    params
  })
}
// @Tags ShareholderProfit
// @Summary 不需要鉴权的股东分润接口
// @Accept application/json
// @Produce application/json
// @Param data query request.ShareholderProfitSearch true "分页获取股东分润列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /spr/getShareholderProfitPublic [get]
export const getShareholderProfitPublic = () => {
  return service({
    url: '/spr/getShareholderProfitPublic',
    method: 'get',
  })
}
