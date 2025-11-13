
package service

import (
    "context"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var ProductSpecOption = new(PSO)

type PSO struct {}
// CreateProductSpecOption 创建商品规格值记录
// Author [yourname](https://github.com/yourname)
func (s *PSO) CreateProductSpecOption(ctx context.Context, PSO *model.ProductSpecOption, merchantID int) (err error) {
    PSO.MerchantID = ptrInt64(int64(merchantID))
    err = global.GVA_DB.WithContext(ctx).Create(PSO).Error
    return err
}

// DeleteProductSpecOption 删除商品规格值记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PSO) DeleteProductSpecOption(ctx context.Context, ID string, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&model.ProductSpecOption{}, "id = ?", ID).Error
    return err
}

// DeleteProductSpecOptionByIds 批量删除商品规格值记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PSO) DeleteProductSpecOptionByIds(ctx context.Context, IDs []string, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&[]model.ProductSpecOption{}, "id in ?", IDs).Error
    return err
}

// UpdateProductSpecOption 更新商品规格值记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PSO) UpdateProductSpecOption(ctx context.Context, PSO model.ProductSpecOption, merchantID int) (err error) {
    err = global.GVA_DB.WithContext(ctx).Model(&model.ProductSpecOption{}).Where("id = ? AND merchant_id = ?", PSO.ID, merchantID).Updates(&PSO).Error
    return err
}

// GetProductSpecOption 根据ID获取商品规格值记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *PSO) GetProductSpecOption(ctx context.Context, ID string, merchantID int) (PSO model.ProductSpecOption, err error) {
    err = global.GVA_DB.WithContext(ctx).Where("id = ? AND merchant_id = ?", ID, merchantID).First(&PSO).Error
    return
}
// GetProductSpecOptionInfoList 分页获取商品规格值记录
// Author [yourname](https://github.com/yourname)
func (s *PSO) GetProductSpecOptionInfoList(ctx context.Context, info request.ProductSpecOptionSearch, merchantID int) (list []model.ProductSpecOption, total int64, err error) {
    limit := info.PageSize
    offset := info.PageSize * (info.Page - 1)
    // 创建db
    db := global.GVA_DB.WithContext(ctx).Model(&model.ProductSpecOption{})
    var PSOs []model.ProductSpecOption
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.ProductID != nil {
        db = db.Where("product_id = ?", *info.ProductID)
    }
    if info.SpecID != nil {
        db = db.Where("spec_id = ?", *info.SpecID)
    }
    if info.Value != nil && *info.Value != "" {
        db = db.Where("value LIKE ?", "%"+ *info.Value+"%")
    }
    if info.SortValue != nil {
        db = db.Where("sort = ?", *info.SortValue)
    }
    db = db.Where("merchant_id = ?", merchantID)
    err = db.Count(&total).Error
    if err!=nil {
        return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
    err = db.Find(&PSOs).Error
    return  PSOs, total, err
}
func (s *PSO)GetProductSpecOptionDataSource(ctx context.Context, merchantID int) (res map[string][]map[string]any, err error) {
    res = make(map[string][]map[string]any)
    productId := make([]map[string]any, 0)
    global.GVA_DB.Table("app_products").Where("deleted_at IS NULL AND merchant_id = ?", merchantID).Select("name as label,id as value").Scan(&productId)
    res["productId"] = productId
    specId := make([]map[string]any, 0)
    global.GVA_DB.Table("app_product_specs").Where("deleted_at IS NULL AND merchant_id = ?", merchantID).Select("name as label,id as value").Scan(&specId)
    res["specId"] = specId
    return
}

func (s *PSO)GetProductSpecOptionPublic(ctx context.Context) {

}
