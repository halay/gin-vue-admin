package router

import (
	"github.com/gin-gonic/gin"
)

var YApi = new(yApi)

type yApi struct{}

// Init 初始化 商家信息 路由信息
func (r *yApi) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
		group := public.Group("yApi")
		group.GET("getKLine", apiYApi.GetYApiKLine)                         // K线
		group.GET("getDeep", apiYApi.GetYApiDeep)                           // 深度
		group.GET("getAllMarketInfo", apiYApi.GetYApiAllMarketInfo)         // 所有交易对行情
		group.GET("getMarketInfo", apiYApi.GetYApiMarketInfo)               // 指定交易对行情
		group.GET("getMerchantsMarketList", apiYApi.GetMerchantsMarketList) // 商户行情列表
		group.GET("getXCgKLine", apiYApi.GetXCgKLine)                       // XCgK线
		group.GET("getXCgKOhlc", apiYApi.GetXCgCoinsOHLC)                   // XCg Coins OHLC
		group.GET("getShopGoods", apiYApi.GetShopGoods)                     // GetShopGoods
		group.POST("createOrder", apiYApi.CreateOrder)                      // CreateOrder

		privateGroup := private.Group("yApi")
		privateGroup.POST("getBLCTYImages", apiYApi.GetBLCTYImages)          // GetBLCTYImages
		privateGroup.GET("getBLCTYImageResult", apiYApi.GetBLCTYImageResult) // 获取图片生成结果
		privateGroup.PUT("deleteBLCTYImage", apiYApi.DeleteBLCTYImageResult) // 删除已经生成的图片
		privateGroup.POST("uploadCozeFile", apiYApi.UploadCozeFile)          //上传文件到Coze
		privateGroup.POST("executeCozeTask", apiYApi.ExecuteCozeTask)        //执行Coze任务
		privateGroup.POST("getCozeTaskResult", apiYApi.GetCozeTaskResult)    //获取执行coze的任务
		privateGroup.POST("reExecuteCozeTask", apiYApi.ReExecuteCozeTask)    //获取执行coze的任务

	}
}
