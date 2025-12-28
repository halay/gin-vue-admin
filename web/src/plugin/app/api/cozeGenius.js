import service from '@/utils/request'
// @Tags getBLCTYImages
// @Summary 创建营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentLevel true "创建营销方案设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /AL/createAgentLevel [post]
export const getBLCTYImages = (data) => {
  if (!(data instanceof FormData)) {
      // 兼容旧的调用方式（虽然 cozeGenius.vue 改了，但为了健壮性）
      // 这里其实 cozeGenius.vue 传过来的已经是 FormData 了，或者我们可以要求调用方自己构建 FormData
      // 但为了方便，我们假设 cozeGenius.vue 会传 FormData
  }
  return service({
    url: '/yApi/getBLCTYImages',
    // 移除 Content-Type: application/json，让浏览器自动设置 multipart/form-data
    method: 'post',
    data
  })
}

