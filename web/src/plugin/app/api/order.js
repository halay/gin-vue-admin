import service from '@/utils/request'
// @Tags Order
// @Summary 创建订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "创建订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ORD/createOrder [post]
export const createOrder = (data) => {
  return service({
    url: '/ORD/createOrder',
    method: 'post',
    data
  })
}

// @Tags Order
// @Summary 删除订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ORD/deleteOrder [delete]
export const deleteOrder = (params) => {
  return service({
    url: '/ORD/deleteOrder',
    method: 'delete',
    params
  })
}

// @Tags Order
// @Summary 批量删除订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ORD/deleteOrder [delete]
export const deleteOrderByIds = (params) => {
  return service({
    url: '/ORD/deleteOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags Order
// @Summary 更新订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Order true "更新订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ORD/updateOrder [put]
export const updateOrder = (data) => {
  return service({
    url: '/ORD/updateOrder',
    method: 'put',
    data
  })
}

// @Tags Order
// @Summary 用id查询订单
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Order true "用id查询订单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ORD/findOrder [get]
export const findOrder = (params) => {
  return service({
    url: '/ORD/findOrder',
    method: 'get',
    params
  })
}

// @Tags Order
// @Summary 分页获取订单列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ORD/getOrderList [get]
export const getOrderList = (params) => {
  return service({
    url: '/ORD/getOrderList',
    method: 'get',
    params
  })
}
// @Tags Order
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ORD/findOrderDataSource [get]
export const getOrderDataSource = () => {
  return service({
    url: '/ORD/getOrderDataSource',
    method: 'get',
  })
}
// @Tags Order
// @Summary 不需要鉴权的订单接口
// @Accept application/json
// @Produce application/json
// @Param data query request.OrderSearch true "分页获取订单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ORD/getOrderPublic [get]
export const getOrderPublic = () => {
  return service({
    url: '/ORD/getOrderPublic',
    method: 'get',
  })
}
