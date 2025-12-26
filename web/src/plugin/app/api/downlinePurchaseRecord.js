import service from '@/utils/request'

// @Tags DownlinePurchaseRecord
// @Summary 分页获取下级购买记录列表 (后台/商户)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.DownlinePurchaseRecordSearch true "分页获取下级购买记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dpr/getDownlinePurchaseRecordList [get]
export const getDownlinePurchaseRecordList = (params) => {
  return service({
    url: '/dpr/getDownlinePurchaseRecordList',
    method: 'get',
    params
  })
}
