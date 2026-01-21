package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var OrderItem = new(ORDI)

type ORDI struct{}

// CreateOrderItem 创建订单明细记录
// Author [yourname](https://github.com/yourname)
func (s *ORDI) CreateOrderItem(ctx context.Context, ORDI *model.OrderItem) (err error) {
	err = global.GVA_DB.Create(ORDI).Error
	return err
}

// DeleteOrderItem 删除订单明细记录
// Author [yourname](https://github.com/yourname)
func (s *ORDI) DeleteOrderItem(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.OrderItem{}, "id = ?", ID).Error
	return err
}

// DeleteOrderItemByIds 批量删除订单明细记录
// Author [yourname](https://github.com/yourname)
func (s *ORDI) DeleteOrderItemByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.OrderItem{}, "id in ?", IDs).Error
	return err
}

// UpdateOrderItem 更新订单明细记录
// Author [yourname](https://github.com/yourname)
func (s *ORDI) UpdateOrderItem(ctx context.Context, ORDI model.OrderItem) (err error) {
	err = global.GVA_DB.Model(&model.OrderItem{}).Where("id = ?", ORDI.ID).Updates(&ORDI).Error
	return err
}

// GetOrderItem 根据ID获取订单明细记录
// Author [yourname](https://github.com/yourname)
func (s *ORDI) GetOrderItem(ctx context.Context, ID string) (ORDI model.OrderItem, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ORDI).Error
	return
}

// GetOrderItemInfoList 分页获取订单明细记录
// Author [yourname](https://github.com/yourname)
func (s *ORDI) GetOrderItemInfoList(ctx context.Context, info request.OrderItemSearch) (list []model.OrderItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.OrderItem{}).
		Select("app_order_items.*, app_users.email as user_email").
		Joins("LEFT JOIN app_orders ON app_order_items.order_id = app_orders.id").
		Joins("LEFT JOIN app_users ON app_orders.user_id = app_users.id")
	var ORDIs []model.OrderItem
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.OrderID != nil {
		db = db.Where("order_id = ?", *info.OrderID)
	}
	if info.ProductID != nil {
		db = db.Where("product_id = ?", *info.ProductID)
	}
	if info.SkuID != nil {
		db = db.Where("sku_id = ?", *info.SkuID)
	}
	if info.ProductName != nil && *info.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+*info.ProductName+"%")
	}
	if info.SkuAttrs != "" {
		// TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
	}
	if info.Price != nil {
		db = db.Where("price >= ?", *info.Price)
	}
	if info.Points != nil {
		db = db.Where("points >= ?", *info.Points)
	}
	if info.Quantity != nil {
		db = db.Where("quantity >= ?", *info.Quantity)
	}
	if info.TotalAmount != nil {
		db = db.Where("total_amount >= ?", *info.TotalAmount)
	}
	if info.TotalPoints != nil {
		db = db.Where("total_points >= ?", *info.TotalPoints)
	}
	if info.CoverImage != "" {
		// TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
	}
	if info.MerchantID != nil {
		db = db.Where("app_order_items.merchant_id = ?", *info.MerchantID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&ORDIs).Error
	return ORDIs, total, err
}
func (s *ORDI) GetOrderItemDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	orderId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_orders").Where("deleted_at IS NULL").Select("order_no as label,id as value").Scan(&orderId)
	res["orderId"] = orderId
	productId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_products").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&productId)
	res["productId"] = productId
	skuId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_product_skus").Where("deleted_at IS NULL").Select("sku_code as label,id as value").Scan(&skuId)
	res["skuId"] = skuId
	return
}

func (s *ORDI) GetOrderItemPublic(ctx context.Context) {

}
