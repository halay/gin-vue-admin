package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var Product = new(P)

type P struct{}

// CreateProduct 创建商户商品记录
// Author [yourname](https://github.com/yourname)
func (s *P) CreateProduct(ctx context.Context, P *model.Product, merchantID int) (err error) {
	// 校验分类归属
	if P.CategoryID != nil {
		var count int64
		global.GVA_DB.WithContext(ctx).Table("app_product_categories").Where("id = ? AND merchant_id = ? AND deleted_at IS NULL", *P.CategoryID, merchantID).Count(&count)
		if count == 0 {
			return global.GVA_DB.Error
		}
	}
	P.MerchantID = ptrInt64(int64(merchantID))
	err = global.GVA_DB.WithContext(ctx).Create(P).Error
	return err
}

// DeleteProduct 删除商户商品记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *P) DeleteProduct(ctx context.Context, ID string, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&model.Product{}, "id = ?", ID).Error
	return err
}

// DeleteProductByIds 批量删除商户商品记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *P) DeleteProductByIds(ctx context.Context, IDs []string, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&[]model.Product{}, "id in ?", IDs).Error
	return err
}

// UpdateProduct 更新商户商品记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *P) UpdateProduct(ctx context.Context, P model.Product, merchantID int) (err error) {
	if P.CategoryID != nil {
		var count int64
		global.GVA_DB.WithContext(ctx).Table("app_product_categories").Where("id = ? AND merchant_id = ? AND deleted_at IS NULL", *P.CategoryID, merchantID).Count(&count)
		if count == 0 {
			return global.GVA_DB.Error
		}
	}
	err = global.GVA_DB.WithContext(ctx).Model(&model.Product{}).Where("id = ? AND merchant_id = ?", P.ID, merchantID).Updates(&P).Error
	return err
}

// GetProduct 根据ID获取商户商品记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *P) GetProduct(ctx context.Context, ID string, merchantID int) (P model.Product, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ? AND merchant_id = ?", ID, merchantID).First(&P).Error
	return
}

// GetProductInfoList 分页获取商户商品记录
// Author [yourname](https://github.com/yourname)
func (s *P) GetProductInfoList(ctx context.Context, info request.ProductSearch, merchantID int) (list []model.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&model.Product{})
	var Ps []model.Product
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.Type != nil && *info.Type != "" {
		db = db.Where("type = ?", *info.Type)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.CategoryID != nil {
		db = db.Where("category_id = ?", *info.CategoryID)
	}
	if info.CoverImage != "" {
		// TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
	}
	if info.CarouselImages != "" {
		// TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
	}
	if info.Video != "" {
		// TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
	}
	if info.Detail != "" {
		// TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
	}
	if info.HasVariants != nil {
		db = db.Where("has_variants = ?", *info.HasVariants)
	}
	if info.EnablePoints != nil {
		db = db.Where("enable_points = ?", *info.EnablePoints)
	}
	if info.Price != nil {
		db = db.Where("price >= ?", *info.Price)
	}
	if info.Points != nil {
		db = db.Where("points >= ?", *info.Points)
	}
	if info.Stock != nil {
		db = db.Where("stock >= ?", *info.Stock)
	}
	db = db.Where("merchant_id = ?", merchantID)
	db = db.Order("id desc")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["name"] = true
	orderMap["type"] = true
	orderMap["status"] = true
	orderMap["category_id"] = true
	orderMap["has_variants"] = true
	orderMap["enable_points"] = true
	orderMap["price"] = true
	orderMap["points"] = true
	orderMap["stock"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&Ps).Error
	return Ps, total, err
}
func (s *P) GetProductDataSource(ctx context.Context, merchantID int) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	categoryId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_product_categories").Where("deleted_at IS NULL AND merchant_id = ?", merchantID).Select("name as label,id as value").Scan(&categoryId)
	res["categoryId"] = categoryId
	return
}

func (s *P) GetProductPublic(ctx context.Context) {

}
