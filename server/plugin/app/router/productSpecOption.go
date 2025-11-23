package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var ProductSpecOption = new(PSO)

type PSO struct{}

// Init 初始化 商品规格值 路由信息
func (r *PSO) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := private.Group("PSO").Use(middleware.OperationRecord())
		group.POST("createProductSpecOption", apiProductSpecOption.CreateProductSpecOption)             // 新建商品规格值
		group.DELETE("deleteProductSpecOption", apiProductSpecOption.DeleteProductSpecOption)           // 删除商品规格值
		group.DELETE("deleteProductSpecOptionByIds", apiProductSpecOption.DeleteProductSpecOptionByIds) // 批量删除商品规格值
		group.PUT("updateProductSpecOption", apiProductSpecOption.UpdateProductSpecOption)              // 更新商品规格值
	}
	{
		group := private.Group("PSO")
		group.GET("findProductSpecOption", apiProductSpecOption.FindProductSpecOption)       // 根据ID获取商品规格值
		group.GET("getProductSpecOptionList", apiProductSpecOption.GetProductSpecOptionList) // 获取商品规格值列表
	}
	{
		group := private.Group("PSO")
		group.GET("getProductSpecOptionDataSource", apiProductSpecOption.GetProductSpecOptionDataSource) // 获取商品规格值数据源（需鉴权）
	}
	{
		group := public.Group("PSO")
		group.GET("getProductSpecOptionPublic", apiProductSpecOption.GetProductSpecOptionPublic) // 商品规格值开放接口
	}
}
