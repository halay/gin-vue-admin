package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var Product = new(P)

type P struct {}

// Init 初始化 商户商品 路由信息
func (r *P) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("P").Use(middleware.OperationRecord())
		group.POST("createProduct", apiProduct.CreateProduct)   // 新建商户商品
		group.DELETE("deleteProduct", apiProduct.DeleteProduct) // 删除商户商品
		group.DELETE("deleteProductByIds", apiProduct.DeleteProductByIds) // 批量删除商户商品
		group.PUT("updateProduct", apiProduct.UpdateProduct)    // 更新商户商品
	}
	{
	    group := private.Group("P")
		group.GET("findProduct", apiProduct.FindProduct)        // 根据ID获取商户商品
		group.GET("getProductList", apiProduct.GetProductList)  // 获取商户商品列表
	}
    {
        group := private.Group("P")
        group.GET("getProductDataSource", apiProduct.GetProductDataSource)  // 获取商户商品数据源（需鉴权）
    }
    {
        group := public.Group("P")
        group.GET("getProductPublic", apiProduct.GetProductPublic)  // 商户商品开放接口
    }
}
