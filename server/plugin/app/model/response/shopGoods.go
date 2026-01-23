package response

type ShopResponse struct {
	Name  string          `json:"name"`
	ID    int64           `json:"id"`
	Logo  string          `json:"logo"`
	Desc  string          `json:"desc"`
	Goods []GoodsResponse `json:"goods"`
}
type GoodsResponse struct {
	Name    string  `json:"name"`    // 商品名称
	ID      int64   `json:"id"`      // 商品ID
	SkuCode string  `json:"skuCode"` // 商品SKU码
	Price   float64 `json:"price"`   // 商品价格
	Desc    string  `json:"desc"`    // 商品描述
	Image   string  `json:"image"`   // 商品图片
}
