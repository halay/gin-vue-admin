import service from '@/utils/request'
// @Tags MerchantAnnouncement
// @Summary 创建商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAnnouncement true "创建商户公告信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /MA/createMerchantAnnouncement [post]
export const createMerchantAnnouncement = (data) => {
  return service({
    url: '/MA/createMerchantAnnouncement',
    method: 'post',
    data
  })
}

// @Tags MerchantAnnouncement
// @Summary 删除商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAnnouncement true "删除商户公告信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MA/deleteMerchantAnnouncement [delete]
export const deleteMerchantAnnouncement = (params) => {
  return service({
    url: '/MA/deleteMerchantAnnouncement',
    method: 'delete',
    params
  })
}

// @Tags MerchantAnnouncement
// @Summary 批量删除商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商户公告信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MA/deleteMerchantAnnouncement [delete]
export const deleteMerchantAnnouncementByIds = (params) => {
  return service({
    url: '/MA/deleteMerchantAnnouncementByIds',
    method: 'delete',
    params
  })
}

// @Tags MerchantAnnouncement
// @Summary 更新商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MerchantAnnouncement true "更新商户公告信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MA/updateMerchantAnnouncement [put]
export const updateMerchantAnnouncement = (data) => {
  return service({
    url: '/MA/updateMerchantAnnouncement',
    method: 'put',
    data
  })
}

// @Tags MerchantAnnouncement
// @Summary 用id查询商户公告信息
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MerchantAnnouncement true "用id查询商户公告信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MA/findMerchantAnnouncement [get]
export const findMerchantAnnouncement = (params) => {
  return service({
    url: '/MA/findMerchantAnnouncement',
    method: 'get',
    params
  })
}

// @Tags MerchantAnnouncement
// @Summary 分页获取商户公告信息列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商户公告信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MA/getMerchantAnnouncementList [get]
export const getMerchantAnnouncementList = (params) => {
  return service({
    url: '/MA/getMerchantAnnouncementList',
    method: 'get',
    params
  })
}
// @Tags MerchantAnnouncement
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MA/findMerchantAnnouncementDataSource [get]
export const getMerchantAnnouncementDataSource = () => {
  return service({
    url: '/MA/getMerchantAnnouncementDataSource',
    method: 'get',
  })
}
// @Tags MerchantAnnouncement
// @Summary 不需要鉴权的商户公告信息接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MerchantAnnouncementSearch true "分页获取商户公告信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /MA/getMerchantAnnouncementPublic [get]
export const getMerchantAnnouncementPublic = () => {
  return service({
    url: '/MA/getMerchantAnnouncementPublic',
    method: 'get',
  })
}
