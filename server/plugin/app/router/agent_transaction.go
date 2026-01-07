package router

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/api"
)

type AgentTransactionRouter struct{}

func (s *AgentTransactionRouter) InitAgentTransactionRouter(Router *gin.RouterGroup) {
	atRouter := Router.Group("agentTransaction")
	atApi := api.Api.AgentTransaction
	{
		atRouter.GET("getAgentTransactionList", atApi.GetAgentTransactionList) // 获取列表
	}
}
