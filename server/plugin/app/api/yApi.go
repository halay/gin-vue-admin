package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sort"
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
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	// 查询所有含交易对的商户
	type mrow struct {
		ID           int64
		MerchantName string
		Logo         string
		TradingPair  string
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
	client := &http.Client{Timeout: time.Duration(plugin.Config.YApiTimeout) * time.Second}
	base := plugin.Config.YApiUrl + "/openapi/v1/market/getMarketInfo"

	type item struct {
		ID                int64   `json:"id"`
		MerchantName      string  `json:"merchantName"`
		Logo              string  `json:"logo"`
		Market            string  `json:"market"`
		Token             string  `json:"token"`
		LatestPrice       string  `json:"latest_price"`
		PercentChange     float64 `json:"percent_change"`
		CirculateQuantity string  `json:"circulate_quantity"`
	}
	list := make([]item, 0, len(rows))
	for _, r := range rows {
		parts := strings.Split(r.TradingPair, "/")
		if len(parts) != 2 {
			continue
		}
		token := strings.TrimSpace(parts[0])
		market := strings.TrimSpace(parts[1])
		// 请求行情
		url := base + "?market=" + market + "&token=" + token
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		// 解析通用响应
		var resStruct struct {
			Code int `json:"code"`
			Data struct {
				LatestPrice       string `json:"latest_price"`
				PercentChange     string `json:"percent_change"`
				CirculateQuantity string `json:"circulate_quantity"`
				Market            string `json:"market"`
				Token             string `json:"token"`
			} `json:"data"`
		}
		if err := json.Unmarshal(b, &resStruct); err != nil || resStruct.Code != 0 {
			continue
		}
		pc, _ := strconv.ParseFloat(resStruct.Data.PercentChange, 64)
		list = append(list, item{
			ID:                r.ID,
			MerchantName:      r.MerchantName,
			Logo:              r.Logo,
			Market:            market,
			Token:             token,
			LatestPrice:       resStruct.Data.LatestPrice,
			PercentChange:     pc,
			CirculateQuantity: resStruct.Data.CirculateQuantity,
		})
	}
	// 排序：percent_change 降序
	sort.Slice(list, func(i, j int) bool { return list[i].PercentChange > list[j].PercentChange })
	total := len(list)
	start := (page - 1) * size
	if start > total {
		start = total
	}
	end := start + size
	if end > total {
		end = total
	}
	pageList := list[start:end]

	response.OkWithDetailed(response.PageResult{
		List:     pageList,
		Total:    int64(total),
		Page:     page,
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
func (y *yApi) GetBLCTYImages(c *gin.Context) {
	var reqs request.BLCTYImagesRequest
	reqs.ResponseFormat = "b64_json"
	reqs.Model = "nano-banana-2"
	// 尝试从multipart form中绑定普通字段
	if err := c.ShouldBind(&reqs); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reqs.Model = "nano-banana-2"
	reqs.ImageSize = "4K"
	reqs.Prompt = reqs.Prompt + " 所有内容用中文显示"
	// 处理上传的文件
	file, _, err := c.Request.FormFile("file")
	if err == nil {
		defer file.Close()
		// 读取文件内容
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			response.FailWithMessage("读取文件失败: "+err.Error(), c)
			return
		}
		// 转Base64
		imgBase64 := base64.StdEncoding.EncodeToString(fileBytes)
		reqs.Image = []string{imgBase64}
	}
	var apiURL = plugin.Config.BltcyApiUrl + "/v1/images/generations"
	var apiKey = plugin.Config.BltcyApiKey
	client := &http.Client{}
	type blctyResp struct {
		Data []struct {
			URL     string `json:"url"`
			B64Json string `json:"b64_json"`
		} `json:"data"`
	}
	payload, _ := json.Marshal(reqs)
	//uploadSvc := service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
	uploaded := make([]string, 0, 4)
	target := 4
	maxAttempts := 10
	attempts := 0
	imgBaseUrl := plugin.Config.HrefUrl
	for attempts < maxAttempts && len(uploaded) < target {
		attempts++
		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
		if err != nil {
			continue
		}
		req.Header.Add("Authorization", "Bearer "+apiKey)
		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			continue
		}
		b, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			continue
		}
		var obj blctyResp
		if err := json.Unmarshal(b, &obj); err != nil || len(obj.Data) == 0 {
			continue
		}
		//raw := strings.TrimSpace(strings.Trim(obj.Data[0].B64Json, "`\" "))
		raw := obj.Data[0].B64Json
		var payloadB64 string
		ext := "jpg"
		if strings.HasPrefix(raw, "data:") {
			parts := strings.SplitN(raw, ",", 2)
			if len(parts) < 2 {
				continue
			}
			h := parts[0]
			payloadB64 = parts[1]
			if strings.HasPrefix(h, "data:image/") {
				t := strings.TrimPrefix(h, "data:image/")
				if i := strings.Index(t, ";"); i >= 0 {
					t = t[:i]
				}
				switch strings.ToLower(t) {
				case "jpeg", "jpg":
					ext = "jpg"
				case "png":
					ext = "png"
				case "webp":
					ext = "webp"
				case "gif":
					ext = "gif"
				}
			}
		} else {
			payloadB64 = raw
		}
		imgBytes, err := base64.StdEncoding.DecodeString(payloadB64)
		if err != nil || len(imgBytes) == 0 {
			continue
		}
		filename := fmt.Sprintf("blcty_%d.%s", time.Now().UnixNano(), ext)
		/*var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)
		part, _ := writer.CreateFormFile("file", filename)
		if _, err := part.Write(imgBytes); err != nil {
			continue
		}
		writer.Close()
		upReq := &http.Request{Header: make(http.Header), Body: io.NopCloser(&buf)}
		upReq.Header.Set("Content-Type", writer.FormDataContentType())
		_ = upReq.ParseMultipartForm(int64(len(imgBytes)) + 1024)
		_, header, err := upReq.FormFile("file")
		if err != nil {
			continue
		}
		fileRec, err := uploadSvc.UploadFile(header, "0", 0)
		if err != nil {
			continue
		}*/
		filePath := filepath.Join(global.GVA_CONFIG.Local.Path, filename)
		if err := ioutil.WriteFile(filePath, imgBytes, 0644); err != nil {
			fmt.Println(err)
		}
		uploaded = append(uploaded, imgBaseUrl+filePath)
	}
	response.OkWithData(gin.H{"urls": uploaded}, c)
}
