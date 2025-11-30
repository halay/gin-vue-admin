package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"sort"

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

// GetMerchantsMarketList 获取所有含交易对的商户行情（公共）
// @Tags YApi
// @Summary 商户行情列表（公共）
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string}
// @Router /yApi/getMerchantsMarketList [get]
func (y *yApi) GetMerchantsMarketList(c *gin.Context) {
    ctx := c.Request.Context()
    // 读取分页参数
    pageStr := c.Query("page")
    sizeStr := c.Query("pageSize")
    page, _ := strconv.Atoi(pageStr)
    size, _ := strconv.Atoi(sizeStr)
    if page <= 0 { page = 1 }
    if size <= 0 { size = 10 }

    // 查询所有含交易对的商户
    type mrow struct{
        ID int64
        MerchantName string
        Logo string
        TradingPair string
    }
    var rows []mrow
    if err := global.GVA_DB.WithContext(ctx).Table("app_merchants").
        Where("trading_pair IS NOT NULL AND trading_pair <> ''").
        Select("id, merchant_name, logo, trading_pair").
        Scan(&rows).Error; err != nil {
        response.FailWithMessage("查询商户失败:"+err.Error(), c)
        return
    }
    // 调行情
    client := &http.Client{ Timeout: time.Duration(plugin.Config.YApiTimeout) * time.Second }
    base := plugin.Config.YApiUrl + "/openapi/v1/market/getMarketInfo"

    type item struct{
        ID int64 `json:"id"`
        MerchantName string `json:"merchantName"`
        Logo string `json:"logo"`
        Market string `json:"market"`
        Token string `json:"token"`
        LatestPrice string `json:"latest_price"`
        PercentChange float64 `json:"percent_change"`
        CirculateQuantity string `json:"circulate_quantity"`
    }
    list := make([]item, 0, len(rows))
    for _, r := range rows {
        parts := strings.Split(r.TradingPair, "/")
        if len(parts) != 2 { continue }
        token := strings.TrimSpace(parts[0])
        market := strings.TrimSpace(parts[1])
        // 请求行情
        url := base + "?market=" + market + "&token=" + token
        resp, err := client.Get(url)
        if err != nil { continue }
        b, _ := io.ReadAll(resp.Body); resp.Body.Close()
        // 解析通用响应
        var resStruct struct{
            Code int `json:"code"`
            Data struct{
                LatestPrice string `json:"latest_price"`
                PercentChange string `json:"percent_change"`
                CirculateQuantity string `json:"circulate_quantity"`
                Market string `json:"market"`
                Token string `json:"token"`
            } `json:"data"`
        }
        if err := json.Unmarshal(b, &resStruct); err != nil || resStruct.Code != 0 { continue }
        pc, _ := strconv.ParseFloat(resStruct.Data.PercentChange, 64)
        list = append(list, item{
            ID: r.ID,
            MerchantName: r.MerchantName,
            Logo: r.Logo,
            Market: market,
            Token: token,
            LatestPrice: resStruct.Data.LatestPrice,
            PercentChange: pc,
            CirculateQuantity: resStruct.Data.CirculateQuantity,
        })
    }
    // 排序：percent_change 降序
    sort.Slice(list, func(i, j int) bool { return list[i].PercentChange > list[j].PercentChange })
    total := len(list)
    start := (page-1)*size
    if start > total { start = total }
    end := start + size
    if end > total { end = total }
    pageList := list[start:end]

    response.OkWithDetailed(response.PageResult{
        List: pageList,
        Total: int64(total),
        Page: page,
        PageSize: size,
    }, "获取成功", c)
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
