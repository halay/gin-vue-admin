import service from '@/utils/request'
// @Tags getBLCTYImages
// @Summary 创建营销方案设置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AgentLevel true "创建营销方案设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /AL/createAgentLevel [post]
export const getBLCTYImages = (params) => {
  const data = JSON.stringify({
    model: 'nano-banana',
    ...params
  })
  return service({
    url: '/yApi/getBLCTYImages',
     headers: {
      'Content-Type': 'application/json'
    },
    method: 'post',
    data
  })
}

