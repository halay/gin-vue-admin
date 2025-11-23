
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var MerchantCategory = new(MCAT)

type MCAT struct {}
// CreateMerchantCategory 创建商户分类记录
// Author [yourname](https://github.com/yourname)
func (s *MCAT) CreateMerchantCategory(ctx context.Context, MCAT *model.MerchantCategory) (err error) {
	err = global.GVA_DB.Create(MCAT).Error
	return err
}

// DeleteMerchantCategory 删除商户分类记录
// Author [yourname](https://github.com/yourname)
func (s *MCAT) DeleteMerchantCategory(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.MerchantCategory{},"id = ?",ID).Error
	return err
}

// DeleteMerchantCategoryByIds 批量删除商户分类记录
// Author [yourname](https://github.com/yourname)
func (s *MCAT) DeleteMerchantCategoryByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.MerchantCategory{},"id in ?",IDs).Error
	return err
}

// UpdateMerchantCategory 更新商户分类记录
// Author [yourname](https://github.com/yourname)
func (s *MCAT) UpdateMerchantCategory(ctx context.Context, MCAT model.MerchantCategory) (err error) {
	err = global.GVA_DB.Model(&model.MerchantCategory{}).Where("id = ?",MCAT.ID).Updates(&MCAT).Error
	return err
}

// GetMerchantCategory 根据ID获取商户分类记录
// Author [yourname](https://github.com/yourname)
func (s *MCAT) GetMerchantCategory(ctx context.Context, ID string) (MCAT model.MerchantCategory, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&MCAT).Error
	return
}
// GetMerchantCategoryInfoList 分页获取商户分类记录
// Author [yourname](https://github.com/yourname)
func (s *MCAT) GetMerchantCategoryInfoList(ctx context.Context, info request.MerchantCategorySearch) (list []model.MerchantCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.MerchantCategory{})
    var MCATs []model.MerchantCategory
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
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["name"] = true
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
	err = db.Find(&MCATs).Error
	return  MCATs, total, err
}

func (s *MCAT)GetMerchantCategoryPublic(ctx context.Context) {

}
