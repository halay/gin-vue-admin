package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var Merchants = new(mc)

type mc struct{}

// CreateMerchants 创建商家信息记录
// Author [yourname](https://github.com/yourname)
func (s *mc) CreateMerchants(ctx context.Context, mc *model.Merchants) (err error) {
	err = global.GVA_DB.Create(mc).Error
	return err
}

// DeleteMerchants 删除商家信息记录
// Author [yourname](https://github.com/yourname)
func (s *mc) DeleteMerchants(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Merchants{}, "id = ?", ID).Error
	return err
}

// DeleteMerchantsByIds 批量删除商家信息记录
// Author [yourname](https://github.com/yourname)
func (s *mc) DeleteMerchantsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Merchants{}, "id in ?", IDs).Error
	return err
}

// UpdateMerchants 更新商家信息记录
// Author [yourname](https://github.com/yourname)
func (s *mc) UpdateMerchants(ctx context.Context, mc model.Merchants) (err error) {
	err = global.GVA_DB.Model(&model.Merchants{}).Where("id = ?", mc.ID).Updates(&mc).Error
	return err
}

// GetMerchants 根据ID获取商家信息记录
// Author [yourname](https://github.com/yourname)
func (s *mc) GetMerchants(ctx context.Context, ID string) (mc model.Merchants, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mc).Error
	return
}

// GetMerchantsInfoList 分页获取商家信息记录
// Author [yourname](https://github.com/yourname)
func (s *mc) GetMerchantsInfoList(ctx context.Context, info request.MerchantsSearch) (list []model.Merchants, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Merchants{})
	var mcs []model.Merchants
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.MerchantName != nil && *info.MerchantName != "" {
		db = db.Where("merchant_name LIKE ?", "%"+*info.MerchantName+"%")
	}
	if info.BusinessScope != nil && *info.BusinessScope != "" {
		db = db.Where("business_scope LIKE ?", "%"+*info.BusinessScope+"%")
	}
	if info.ContactPerson != nil && *info.ContactPerson != "" {
		db = db.Where("contact_person LIKE ?", "%"+*info.ContactPerson+"%")
	}
	if info.ContactPhone != nil && *info.ContactPhone != "" {
		db = db.Where("contact_phone LIKE ?", "%"+*info.ContactPhone+"%")
	}
	if info.Email != nil && *info.Email != "" {
		db = db.Where("email LIKE ?", "%"+*info.Email+"%")
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.Founder != nil && *info.Founder != "" {
		db = db.Where("founder LIKE ?", "%"+*info.Founder+"%")
	}
	if info.FounderPhone != nil && *info.FounderPhone != "" {
		db = db.Where("founder_phone LIKE ?", "%"+*info.FounderPhone+"%")
	}
	if info.DigitalAssetName != nil && *info.DigitalAssetName != "" {
		db = db.Where("digital_asset_name LIKE ?", "%"+*info.DigitalAssetName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["merchant_name"] = true
	orderMap["rating"] = true
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
	err = db.Find(&mcs).Error
	return mcs, total, err
}

func (s *mc) GetMerchantsPublic(ctx context.Context) {

}
