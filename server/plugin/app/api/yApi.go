package api

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/plugin"
)

var YApi = new(yApi)

type yApi struct{}

func (y *yApi) GetYApiKLine(c *gin.Context) {
	var req request.YApiKLineRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var url = plugin.Config.YApiUrl + "/openapi/v1/market/kline"
	global.GVA_LOG.Info("请求K线url:" + url)
	client := &http.Client{
		Timeout: time.Duration(plugin.Config.YApiTimeout) * time.Second,
	}
	url = url + "?market=" + req.Market + "&token=" + req.Token + "&time_type=" + req.TimeType
	res, err := client.Get(url)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		global.GVA_LOG.Error("请求K线失败!", zap.Error(err))
		response.FailWithMessage("请求K线失败"+err.Error(), c)
		return
	}
	response.OkWithData(b, c)
}

func (y *yApi) GetYApiDeep(c *gin.Context) {
	var req request.YApiDeepRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var url = plugin.Config.YApiUrl + "/openapi/v1/market/books"
	global.GVA_LOG.Info("请求深度url:" + url)
	client := &http.Client{
		Timeout: time.Duration(plugin.Config.YApiTimeout) * time.Second,
	}
	//判断req.Limit默认为100，最大300
	if req.Limit > 300 {
		req.Limit = 300
	}
	if req.Limit <= 0 {
		req.Limit = 100
	}
	url = url + "?market=" + req.Market + "&token=" + req.Token + "&limit=" + strconv.Itoa(req.Limit)
	res, err := client.Get(url)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		global.GVA_LOG.Error("请求深度失败!", zap.Error(err))
		response.FailWithMessage("请求深度失败"+err.Error(), c)
		return
	}
	response.OkWithData(b, c)
}
func (y *yApi) GetYApiAllMarketInfo(c *gin.Context) {
	var url = plugin.Config.YApiUrl + "/openapi/v1/market/getAllMarketInfo"
	global.GVA_LOG.Info("请求所有交易对行情url:" + url)
	client := &http.Client{
		Timeout: time.Duration(plugin.Config.YApiTimeout) * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		global.GVA_LOG.Error("请求所有交易对行情失败!", zap.Error(err))
		response.FailWithMessage("请求所有交易对行情失败"+err.Error(), c)
		return
	}
	response.OkWithData(b, c)
}
func (y *yApi) GetYApiMarketInfo(c *gin.Context) {
	var req request.YApiMarketInfoRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var url = plugin.Config.YApiUrl + "/openapi/v1/market/getMarketInfo"
	global.GVA_LOG.Info("请求指定交易对行情url:" + url)
	client := &http.Client{
		Timeout: time.Duration(plugin.Config.YApiTimeout) * time.Second,
	}
	url = url + "?market=" + req.Market + "&token=" + req.Token
	res, err := client.Get(url)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		global.GVA_LOG.Error("请求指定交易对行情失败!", zap.Error(err))
		response.FailWithMessage("请求指定交易对行情失败"+err.Error(), c)
		return
	}
	response.OkWithData(b, c)
}
