package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var ProductCategory = new(PC)

type PC struct {}

// Init 初始化 商户商品分类 路由信息
func (r *PC) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("PC").Use(middleware.OperationRecord())
		group.POST("createProductCategory", apiProductCategory.CreateProductCategory)   // 新建商户商品分类
		group.DELETE("deleteProductCategory", apiProductCategory.DeleteProductCategory) // 删除商户商品分类
		group.DELETE("deleteProductCategoryByIds", apiProductCategory.DeleteProductCategoryByIds) // 批量删除商户商品分类
		group.PUT("updateProductCategory", apiProductCategory.UpdateProductCategory)    // 更新商户商品分类
	}
	{
	    group := private.Group("PC")
		group.GET("findProductCategory", apiProductCategory.FindProductCategory)        // 根据ID获取商户商品分类
		group.GET("getProductCategoryList", apiProductCategory.GetProductCategoryList)  // 获取商户商品分类列表
	}
    {
        group := public.Group("PC")
        group.GET("getProductCategoryPublic", apiProductCategory.GetProductCategoryPublic)  // 商户商品分类开放接口
        group.GET("getCategoryTreePublic", apiProductCategory.GetCategoryTreePublic)
    }
}
