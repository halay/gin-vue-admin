package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var ProductSku = new(SKU)

type SKU struct {}

// Init 初始化 商品SKU（规格组合） 路由信息
func (r *SKU) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("SKU").Use(middleware.OperationRecord())
		group.POST("createProductSku", apiProductSku.CreateProductSku)   // 新建商品SKU（规格组合）
		group.DELETE("deleteProductSku", apiProductSku.DeleteProductSku) // 删除商品SKU（规格组合）
		group.DELETE("deleteProductSkuByIds", apiProductSku.DeleteProductSkuByIds) // 批量删除商品SKU（规格组合）
		group.PUT("updateProductSku", apiProductSku.UpdateProductSku)    // 更新商品SKU（规格组合）
	}
	{
	    group := private.Group("SKU")
		group.GET("findProductSku", apiProductSku.FindProductSku)        // 根据ID获取商品SKU（规格组合）
		group.GET("getProductSkuList", apiProductSku.GetProductSkuList)  // 获取商品SKU（规格组合）列表
	}
    {
        group := private.Group("SKU")
        group.GET("getProductSkuDataSource", apiProductSku.GetProductSkuDataSource)  // 获取商品SKU（规格组合）数据源（需鉴权）
    }
    {
        group := public.Group("SKU")
        group.GET("getProductSkuPublic", apiProductSku.GetProductSkuPublic)  // 商品SKU（规格组合）开放接口
    }
}
