package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var ShareholderProfit = new(spr)

type spr struct{}

// CreateShareholderProfit 创建股东分润记录
// Author [yourname](https://github.com/yourname)
func (s *spr) CreateShareholderProfit(ctx context.Context, spr *model.ShareholderProfit, merchantID int) (err error) {
	spr.MerchantID = ptrInt64(int64(merchantID))
	err = global.GVA_DB.WithContext(ctx).Create(spr).Error
	return err
}

// DeleteShareholderProfit 删除股东分润记录
// Author [yourname](https://github.com/yourname)
func (s *spr) DeleteShareholderProfit(ctx context.Context, ID string, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&model.ShareholderProfit{}, "id = ?", ID).Error
	return err
}

// DeleteShareholderProfitByIds 批量删除股东分润记录
// Author [yourname](https://github.com/yourname)
func (s *spr) DeleteShareholderProfitByIds(ctx context.Context, IDs []string, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&[]model.ShareholderProfit{}, "id in ?", IDs).Error
	return err
}

// UpdateShareholderProfit 更新股东分润记录
// Author [yourname](https://github.com/yourname)
func (s *spr) UpdateShareholderProfit(ctx context.Context, spr model.ShareholderProfit, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&model.ShareholderProfit{}).Where("id = ? AND merchant_id = ?", spr.ID, merchantID).Updates(&spr).Error
	return err
}

// GetShareholderProfit 根据ID获取股东分润记录
// Author [yourname](https://github.com/yourname)
func (s *spr) GetShareholderProfit(ctx context.Context, ID string, merchantID int) (spr model.ShareholderProfit, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ? AND merchant_id = ?", ID, merchantID).First(&spr).Error
	return
}

// GetShareholderProfitInfoList 分页获取股东分润记录
// Author [yourname](https://github.com/yourname)
func (s *spr) GetShareholderProfitInfoList(ctx context.Context, info request.ShareholderProfitSearch, merchantID int) (list []model.ShareholderProfit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&model.ShareholderProfit{})
	var sprs []model.ShareholderProfit
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	db = db.Where("merchant_id = ?", merchantID)

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["name"] = true
	orderMap["share_ratio"] = true
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
	err = db.Find(&sprs).Error
	return sprs, total, err
}

func (s *spr) GetShareholderProfitPublic(ctx context.Context) {

}
