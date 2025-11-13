package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var ProductSpec = new(PS)

type PS struct {}

// Init 初始化 商品规格键 路由信息
func (r *PS) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("PS").Use(middleware.OperationRecord())
		group.POST("createProductSpec", apiProductSpec.CreateProductSpec)   // 新建商品规格键
		group.DELETE("deleteProductSpec", apiProductSpec.DeleteProductSpec) // 删除商品规格键
		group.DELETE("deleteProductSpecByIds", apiProductSpec.DeleteProductSpecByIds) // 批量删除商品规格键
		group.PUT("updateProductSpec", apiProductSpec.UpdateProductSpec)    // 更新商品规格键
	}
	{
	    group := private.Group("PS")
		group.GET("findProductSpec", apiProductSpec.FindProductSpec)        // 根据ID获取商品规格键
		group.GET("getProductSpecList", apiProductSpec.GetProductSpecList)  // 获取商品规格键列表
	}
    {
        group := private.Group("PS")
        group.GET("getProductSpecDataSource", apiProductSpec.GetProductSpecDataSource)  // 获取商品规格键数据源（需鉴权）
    }
    {
        group := public.Group("PS")
        group.GET("getProductSpecPublic", apiProductSpec.GetProductSpecPublic)  // 商品规格键开放接口
    }
}
