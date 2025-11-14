
package service

import (
    "context"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var ProductSku = new(SKU)

type SKU struct {}
// CreateProductSku 创建商品SKU（规格组合）记录
// Author [yourname](https://github.com/yourname)
func (s *SKU) CreateProductSku(ctx context.Context, SKU *model.ProductSku, merchantID int) (err error) {
    var cnt int64
    global.GVA_DB.WithContext(ctx).Table("app_products").Where("id = ? AND merchant_id = ? AND deleted_at IS NULL", SKU.ProductID, merchantID).Count(&cnt)
    if cnt == 0 {
        return global.GVA_DB.Error
    }
    SKU.MerchantID = ptrInt64(int64(merchantID))
    err = global.GVA_DB.WithContext(ctx).Create(SKU).Error
    if err == nil {
        _ = ProductSkuOptionSvc.UpsertSkuOptions(ctx, SKU.ID, merchantID, SKU.Attrs)
    }
    return err
}

// DeleteProductSku 删除商品SKU（规格组合）记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *SKU) DeleteProductSku(ctx context.Context, ID string, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&model.ProductSku{}, "id = ?", ID).Error
    return err
}

// DeleteProductSkuByIds 批量删除商品SKU（规格组合）记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *SKU) DeleteProductSkuByIds(ctx context.Context, IDs []string, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&[]model.ProductSku{}, "id in ?", IDs).Error
    return err
}

// UpdateProductSku 更新商品SKU（规格组合）记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *SKU) UpdateProductSku(ctx context.Context, SKU model.ProductSku, merchantID int) (err error) {
    if SKU.ProductID != nil {
        var cnt int64
        global.GVA_DB.WithContext(ctx).Table("app_products").Where("id = ? AND merchant_id = ? AND deleted_at IS NULL", *SKU.ProductID, merchantID).Count(&cnt)
        if cnt == 0 {
            return global.GVA_DB.Error
        }
    }
    err = global.GVA_DB.WithContext(ctx).Model(&model.ProductSku{}).Where("id = ? AND merchant_id = ?", SKU.ID, merchantID).Updates(&SKU).Error
    if err == nil {
        _ = ProductSkuOptionSvc.UpsertSkuOptions(ctx, SKU.ID, merchantID, SKU.Attrs)
    }
    return err
}

// GetProductSku 根据ID获取商品SKU（规格组合）记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *SKU) GetProductSku(ctx context.Context, ID string, merchantID int) (SKU model.ProductSku, err error) {
    err = global.GVA_DB.WithContext(ctx).Where("id = ? AND merchant_id = ?", ID, merchantID).First(&SKU).Error
    return
}
// GetProductSkuInfoList 分页获取商品SKU（规格组合）记录
// Author [yourname](https://github.com/yourname)
func (s *SKU) GetProductSkuInfoList(ctx context.Context, info request.ProductSkuSearch, merchantID int) (list []model.ProductSku, total int64, err error) {
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    // 创建db
    db := global.GVA_DB.WithContext(ctx).Model(&model.ProductSku{})
    var SKUs []model.ProductSku
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.ProductID != nil {
        db = db.Where("product_id = ?", *info.ProductID)
    }
    if info.SKUCode != nil && *info.SKUCode != "" {
        db = db.Where("sku_code LIKE ?", "%"+ *info.SKUCode+"%")
    }
    if info.Attrs != "" {
        // TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
    }
    if info.Price != nil {
        db = db.Where("price >= ?", *info.Price)
    }
    if info.Points != nil {
        db = db.Where("points >= ?", *info.Points)
    }
    if info.Image != "" {
        // TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
    }
    if info.Stock != nil {
        db = db.Where("stock >= ?", *info.Stock)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
    db = db.Where("merchant_id = ?", merchantID)
    err = db.Count(&total).Error
    if err!=nil {
        return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["product_id"] = true
        orderMap["sku_code"] = true
        orderMap["price"] = true
        orderMap["points"] = true
        orderMap["stock"] = true
        orderMap["status"] = true
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
    err = db.Find(&SKUs).Error
    return  SKUs, total, err
}
func (s *SKU)GetProductSkuDataSource(ctx context.Context, merchantID int) (res map[string][]map[string]any, err error) {
    res = make(map[string][]map[string]any)
    productId := make([]map[string]any, 0)
    global.GVA_DB.Table("app_products").Where("deleted_at IS NULL AND merchant_id = ?", merchantID).Select("name as label,id as value").Scan(&productId)
    res["productId"] = productId
    return
}

func (s *SKU)GetProductSkuPublic(ctx context.Context) {

}
