
package service

import (
    "context"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var ProductCategory = new(PC)

type PC struct {}
// CreateProductCategory 创建商户商品分类记录
// Author [yourname](https://github.com/yourname)
func (s *PC) CreateProductCategory(ctx context.Context, PC *model.ProductCategory, merchantID int) (err error) {
    PC.MerchantID = ptrInt64(int64(merchantID))
    err = global.GVA_DB.WithContext(ctx).Create(PC).Error
    return err
}

// DeleteProductCategory 删除商户商品分类记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PC) DeleteProductCategory(ctx context.Context, ID string, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&model.ProductCategory{}, "id = ?", ID).Error
    return err
}

// DeleteProductCategoryByIds 批量删除商户商品分类记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PC) DeleteProductCategoryByIds(ctx context.Context, IDs []string, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&[]model.ProductCategory{}, "id in ?", IDs).Error
    return err
}

// UpdateProductCategory 更新商户商品分类记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PC) UpdateProductCategory(ctx context.Context, PC model.ProductCategory, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Model(&model.ProductCategory{}).Where("id = ? AND merchant_id = ?", PC.ID, merchantID).Updates(&PC).Error
    return err
}

// GetProductCategory 根据ID获取商户商品分类记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PC) GetProductCategory(ctx context.Context, ID string, merchantID int) (PC model.ProductCategory, err error) {
    err = global.GVA_DB.WithContext(ctx).Where("id = ? AND merchant_id = ?", ID, merchantID).First(&PC).Error
    return
}
// GetProductCategoryInfoList 分页获取商户商品分类记录
// Author [yourname](https://github.com/yourname)
func (s *PC) GetProductCategoryInfoList(ctx context.Context, info request.ProductCategorySearch, merchantID int) (list []model.ProductCategory, total int64, err error) {
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    // 创建db
    db := global.GVA_DB.WithContext(ctx).Model(&model.ProductCategory{})
    var PCs []model.ProductCategory
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
    if info.SortValue != nil {
        db = db.Where("sort = ?", *info.SortValue)
    }
    if info.Description != "" {
        // TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
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
        orderMap["status"] = true
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
    err = db.Find(&PCs).Error
    return  PCs, total, err
}

func (s *PC)GetProductCategoryPublic(ctx context.Context) {

}

func ptrInt64(v int64) *int64 { return &v }
