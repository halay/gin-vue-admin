package request

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
