package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var UserAddress = new(UA)

type UA struct{}

// Init 初始化 用户收货地址 路由信息
func (r *UA) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("UA").Use(middleware.OperationRecord())
		group.POST("createUserAddress", apiUserAddress.CreateUserAddress)             // 新建用户收货地址
		group.DELETE("deleteUserAddress", apiUserAddress.DeleteUserAddress)           // 删除用户收货地址
		group.DELETE("deleteUserAddressByIds", apiUserAddress.DeleteUserAddressByIds) // 批量删除用户收货地址
		group.PUT("updateUserAddress", apiUserAddress.UpdateUserAddress)              // 更新用户收货地址
	}
	{
		group := private.Group("UA")
		group.GET("findUserAddress", apiUserAddress.FindUserAddress)       // 根据ID获取用户收货地址
		group.GET("getUserAddressList", apiUserAddress.GetUserAddressList) // 获取用户收货地址列表
	}
	{
		group := public.Group("UA")
		group.GET("getUserAddressPublic", apiUserAddress.GetUserAddressPublic) // 用户收货地址开放接口
	}
}
