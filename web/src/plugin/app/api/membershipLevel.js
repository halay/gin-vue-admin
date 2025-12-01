import service from '@/utils/request'
// @Tags MembershipLevel
// @Summary 创建会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MembershipLevel true "创建会员等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ML/createMembershipLevel [post]
export const createMembershipLevel = (data) => {
  return service({
    url: '/ML/createMembershipLevel',
    method: 'post',
    data
  })
}

// @Tags MembershipLevel
// @Summary 删除会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MembershipLevel true "删除会员等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ML/deleteMembershipLevel [delete]
export const deleteMembershipLevel = (params) => {
  return service({
    url: '/ML/deleteMembershipLevel',
    method: 'delete',
    params
  })
}

// @Tags MembershipLevel
// @Summary 批量删除会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除会员等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ML/deleteMembershipLevel [delete]
export const deleteMembershipLevelByIds = (params) => {
  return service({
    url: '/ML/deleteMembershipLevelByIds',
    method: 'delete',
    params
  })
}

// @Tags MembershipLevel
// @Summary 更新会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.MembershipLevel true "更新会员等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ML/updateMembershipLevel [put]
export const updateMembershipLevel = (data) => {
  return service({
    url: '/ML/updateMembershipLevel',
    method: 'put',
    data
  })
}

// @Tags MembershipLevel
// @Summary 用id查询会员等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.MembershipLevel true "用id查询会员等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ML/findMembershipLevel [get]
export const findMembershipLevel = (params) => {
  return service({
    url: '/ML/findMembershipLevel',
    method: 'get',
    params
  })
}

// @Tags MembershipLevel
// @Summary 分页获取会员等级列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取会员等级列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ML/getMembershipLevelList [get]
export const getMembershipLevelList = (params) => {
  return service({
    url: '/ML/getMembershipLevelList',
    method: 'get',
    params
  })
}
// @Tags MembershipLevel
// @Summary 不需要鉴权的会员等级接口
// @Accept application/json
// @Produce application/json
// @Param data query request.MembershipLevelSearch true "分页获取会员等级列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /ML/getMembershipLevelPublic [get]
export const getMembershipLevelPublic = () => {
  return service({
    url: '/ML/getMembershipLevelPublic',
    method: 'get',
  })
}
