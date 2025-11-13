import service from '@/utils/request'
// @Tags OrderItem
// @Summary 创建订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "创建订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ORDI/createOrderItem [post]
export const createOrderItem = (data) => {
  return service({
    url: '/ORDI/createOrderItem',
    method: 'post',
    data
  })
}

// @Tags OrderItem
// @Summary 删除订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "删除订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ORDI/deleteOrderItem [delete]
export const deleteOrderItem = (params) => {
  return service({
    url: '/ORDI/deleteOrderItem',
    method: 'delete',
    params
  })
}

// @Tags OrderItem
// @Summary 批量删除订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ORDI/deleteOrderItem [delete]
export const deleteOrderItemByIds = (params) => {
  return service({
    url: '/ORDI/deleteOrderItemByIds',
    method: 'delete',
    params
  })
}

// @Tags OrderItem
// @Summary 更新订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.OrderItem true "更新订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ORDI/updateOrderItem [put]
export const updateOrderItem = (data) => {
  return service({
    url: '/ORDI/updateOrderItem',
    method: 'put',
    data
  })
}

// @Tags OrderItem
// @Summary 用id查询订单明细
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.OrderItem true "用id查询订单明细"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ORDI/findOrderItem [get]
export const findOrderItem = (params) => {
  return service({
    url: '/ORDI/findOrderItem',
    method: 'get',
    params
  })
}

// @Tags OrderItem
// @Summary 分页获取订单明细列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取订单明细列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ORDI/getOrderItemList [get]
export const getOrderItemList = (params) => {
  return service({
    url: '/ORDI/getOrderItemList',
    method: 'get',
    params
  })
}
// @Tags OrderItem
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ORDI/findOrderItemDataSource [get]
export const getOrderItemDataSource = () => {
  return service({
    url: '/ORDI/getOrderItemDataSource',
    method: 'get',
  })
}
// @Tags OrderItem
// @Summary 不需要鉴权的订单明细接口
// @Accept application/json
// @Produce application/json
// @Param data query request.OrderItemSearch true "分页获取订单明细列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ORDI/getOrderItemPublic [get]
export const getOrderItemPublic = () => {
  return service({
    url: '/ORDI/getOrderItemPublic',
    method: 'get',
  })
}
