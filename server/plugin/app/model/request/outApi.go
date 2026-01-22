package request

import "encoding/json"

type YApiKLineRequest struct {
	Market   string `json:"market" form:"market" binding:"required"`       //USDT
	Token    string `json:"token" form:"token" binding:"required"`         //BTC
	TimeType string `json:"time_type" form:"time_type" binding:"required"` //时间粒度，单位：分钟；可取值：1、5、15、30、60、240、360、1440、10080、43200
}
type YApiDeepRequest struct {
	Market string `json:"market" form:"market" binding:"required"` //USDT
	Token  string `json:"token" form:"token" binding:"required"`   //BTC
	//Limit 默认100 最大300 非必填
	Limit int `json:"limit" form:"limit"`
}
type YApiMarketInfoRequest struct {
	Market string `json:"market" form:"market" binding:"required"` //USDT
	Token  string `json:"token" form:"token" binding:"required"`   //BTC
}
type XCgCoinOHLCRequest struct {
	Coins string `json:"coins" form:"coins" binding:"required"` //bitcoin
}

type BLCTYImagesRequest struct {
	Model          string   `json:"model" form:"model" binding:"required"`   //model
	Prompt         string   `json:"prompt" form:"prompt" binding:"required"` //prompt
	AspectRatio    string   `json:"aspect_ratio" form:"aspect_ratio"`        //aspect_ratio
	ResponseFormat string   `json:"response_format" form:"response_format"`  //response_format
	Image          []string `json:"image" form:"image"`                      //image
	ImageSize      string   `json:"image_size" form:"image_size"`            //image_size
}

type GetBLCTYImagesRequest struct {
	ID int `json:"id" form:"id" binding:"required"`
}

type DeleteBLCTYImagesRequest struct {
	ID     int   `json:"id" form:"id" binding:"required"`
	Indexs []int `json:"indexes" form:"indexes"`
}

/*
*
"bytes": 9139,
"created_at": 1769049294,
"file_name": "客户.png",
"id": "7598008397669695526"
*/
type CoreUploadData struct {
	Bytes    int    `json:"bytes"`
	Created  int    `json:"created_at"`
	FileName string `json:"file_name"`
	Id       string `json:"id"`
}

type CozeResult struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Detail  struct {
		Logid string `json:"logid"`
	} `json:"detail"`
	Data      json.RawMessage `json:"data"`
	ExecuteId string          `json:"execute_id"`
}

type CozeTaskRequest struct {
	WorkflowId string            `json:"workflow_id"` //工作流id
	Type       string            `json:"type"`        //类型
	Parameters map[string]string `json:"parameters"`  //参数
}

type CozeWorkflowRequest struct {
	WorkflowId      string            `json:"workflow_id"`
	IsAsync         bool              `json:"is_async"`
	Parameters      map[string]string `json:"parameters"`
	AppID           string            `json:"app_id,omitempty"`
	WorkflowVersion string            `json:"workflow_version,omitempty"`
}

func (r *CozeResult) Decode(v interface{}) error {
	return json.Unmarshal(r.Data, v)
}
