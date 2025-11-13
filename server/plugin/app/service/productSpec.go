
package service

import (
    "context"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var ProductSpec = new(PS)

type PS struct {}
// CreateProductSpec 创建商品规格键记录
// Author [yourname](https://github.com/yourname)
func (s *PS) CreateProductSpec(ctx context.Context, PS *model.ProductSpec, merchantID int) (err error) {
    PS.MerchantID = ptrInt64(int64(merchantID))
    err = global.GVA_DB.WithContext(ctx).Create(PS).Error
    return err
}

// DeleteProductSpec 删除商品规格键记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PS) DeleteProductSpec(ctx context.Context, ID string, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&model.ProductSpec{}, "id = ?", ID).Error
    return err
}

// DeleteProductSpecByIds 批量删除商品规格键记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PS) DeleteProductSpecByIds(ctx context.Context, IDs []string, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&[]model.ProductSpec{}, "id in ?", IDs).Error
    return err
}

// UpdateProductSpec 更新商品规格键记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PS) UpdateProductSpec(ctx context.Context, PS model.ProductSpec, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Model(&model.ProductSpec{}).Where("id = ? AND merchant_id = ?", PS.ID, merchantID).Updates(&PS).Error
    return err
}

// GetProductSpec 根据ID获取商品规格键记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PS) GetProductSpec(ctx context.Context, ID string, merchantID int) (PS model.ProductSpec, err error) {
    err = global.GVA_DB.WithContext(ctx).Where("id = ? AND merchant_id = ?", ID, merchantID).First(&PS).Error
    return
}
// GetProductSpecInfoList 分页获取商品规格键记录
// Author [yourname](https://github.com/yourname)
func (s *PS) GetProductSpecInfoList(ctx context.Context, info request.ProductSpecSearch, merchantID int) (list []model.ProductSpec, total int64, err error) {
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    // 创建db
    db := global.GVA_DB.WithContext(ctx).Model(&model.ProductSpec{})
    var PSs []model.ProductSpec
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.ProductID != nil {
        db = db.Where("product_id = ?", *info.ProductID)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.SortValue != nil {
        db = db.Where("sort = ?", *info.SortValue)
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
        orderMap["name"] = true
        orderMap["sort"] = true
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
    err = db.Find(&PSs).Error
    return  PSs, total, err
}
func (s *PS)GetProductSpecDataSource(ctx context.Context, merchantID int) (res map[string][]map[string]any, err error) {
    res = make(map[string][]map[string]any)
    productId := make([]map[string]any, 0)
    global.GVA_DB.Table("app_products").Where("deleted_at IS NULL AND merchant_id = ?", merchantID).Select("name as label,id as value").Scan(&productId)
    res["productId"] = productId
    return
}

func (s *PS)GetProductSpecPublic(ctx context.Context) {

}
