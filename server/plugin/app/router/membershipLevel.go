package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var MembershipLevel = new(ML)

type ML struct {}

// Init 初始化 会员等级 路由信息
func (r *ML) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("ML").Use(middleware.OperationRecord())
		group.POST("createMembershipLevel", apiMembershipLevel.CreateMembershipLevel)   // 新建会员等级
		group.DELETE("deleteMembershipLevel", apiMembershipLevel.DeleteMembershipLevel) // 删除会员等级
		group.DELETE("deleteMembershipLevelByIds", apiMembershipLevel.DeleteMembershipLevelByIds) // 批量删除会员等级
		group.PUT("updateMembershipLevel", apiMembershipLevel.UpdateMembershipLevel)    // 更新会员等级
	}
	{
	    group := private.Group("ML")
		group.GET("findMembershipLevel", apiMembershipLevel.FindMembershipLevel)        // 根据ID获取会员等级
		group.GET("getMembershipLevelList", apiMembershipLevel.GetMembershipLevelList)  // 获取会员等级列表
	}
	{
	    group := public.Group("ML")
	    group.GET("getMembershipLevelPublic", apiMembershipLevel.GetMembershipLevelPublic)  // 会员等级开放接口
	}
}
