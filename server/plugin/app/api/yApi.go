package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
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
	var resStruct response.Response
	err = json.Unmarshal(b, &resStruct)
	if err != nil {
		global.GVA_LOG.Error("请求K线失败!", zap.Error(err))
		response.FailWithMessage("请求K线失败"+err.Error(), c)
		return
	}

	if resStruct.Code != 0 {
		global.GVA_LOG.Error("请求K线失败!"+resStruct.Msg, zap.Error(err))
		response.FailWithMessage("请求K线失败"+resStruct.Msg, c)
		return
	}
	response.OkWithData(resStruct.Data, c)
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
	var resStruct response.Response
	err = json.Unmarshal(b, &resStruct)
	if err != nil {
		global.GVA_LOG.Error("请求深度失败!", zap.Error(err))
		response.FailWithMessage("请求深度失败"+err.Error(), c)
		return
	}

	if resStruct.Code != 0 {
		global.GVA_LOG.Error("请求深度失败!"+resStruct.Msg, zap.Error(err))
		response.FailWithMessage("请求深度失败"+resStruct.Msg, c)
		return
	}
	response.OkWithData(resStruct.Data, c)
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
	var resStruct response.Response
	err = json.Unmarshal(b, &resStruct)
	if err != nil {
		global.GVA_LOG.Error("请求所有交易对行情失败!", zap.Error(err))
		response.FailWithMessage("请求所有交易对行情失败"+err.Error(), c)
		return
	}

	if resStruct.Code != 0 {
		global.GVA_LOG.Error("请求所有交易对行情失败!"+resStruct.Msg, zap.Error(err))
		response.FailWithMessage("请求所有交易对行情失败"+resStruct.Msg, c)
		return
	}
	response.OkWithData(resStruct.Data, c)
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
	var resStruct response.Response
	err = json.Unmarshal(b, &resStruct)
	if err != nil {
		global.GVA_LOG.Error("请求指定交易对行情失败!", zap.Error(err))
		response.FailWithMessage("请求指定交易对行情失败"+err.Error(), c)
		return
	}

	if resStruct.Code != 0 {
		global.GVA_LOG.Error("请求指定交易对行情失败!"+resStruct.Msg, zap.Error(err))
		response.FailWithMessage("请求指定交易对行情失败"+resStruct.Msg, c)
		return
	}
	response.OkWithData(resStruct.Data, c)
}
func (y *yApi) GetXCgKLine(c *gin.Context) {
	var url = plugin.Config.XCgProApiUrl + "/coins/markets?vs_currency=usd&price_change_percentage=1h,24h,7d"
	var key = plugin.Config.XCgProApiKey
	var ids = plugin.Config.XCgApiIds //把ids转换成逗号隔开的字符串放到url后面&连接
	url = url + "&ids=" + strings.Join(ids, ",")
	global.GVA_LOG.Info("请求K线url:" + url)

	// 创建一个新的GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		global.GVA_LOG.Error("创建请求失败", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	// 添加Header头
	req.Header.Add("x-cg-pro-api-key", key)
	// 创建一个http客户端
	client := &http.Client{
		Timeout: time.Duration(plugin.Config.YApiTimeout) * time.Second,
	}
	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		global.GVA_LOG.Error("请求失败", zap.Error(err))
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		global.GVA_LOG.Error("请求K线失败!", zap.Error(err))
		response.FailWithMessage("请求K线失败"+err.Error(), c)
		return
	}
	response.OkWithData(json.RawMessage(body), c)
}
func (y *yApi) GetXCgCoinsOHLC(c *gin.Context) {
	var reqs request.XCgCoinOHLCRequest
	if err := c.ShouldBind(&reqs); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var url = plugin.Config.XCgProApiUrl + "/coins/" + reqs.Coins + "/ohlc?vs_currency=usd&days=7"
	var key = plugin.Config.XCgProApiKey
	global.GVA_LOG.Info("请求Coins OHLC url:" + url)

	// 创建一个新的GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		global.GVA_LOG.Error("创建请求失败", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	// 添加Header头
	req.Header.Add("x-cg-pro-api-key", key)
	// 创建一个http客户端
	client := &http.Client{
		Timeout: time.Duration(plugin.Config.YApiTimeout) * time.Second,
	}
	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		response.FailWithMessage("获取OHLC失败"+err.Error(), c)
		global.GVA_LOG.Error("请求OHLC失败", zap.Error(err))
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		global.GVA_LOG.Error("请求OHLC失败!", zap.Error(err))
		response.FailWithMessage("请求OHLC失败"+err.Error(), c)
		return
	}
	response.OkWithData(json.RawMessage(body), c)
}
