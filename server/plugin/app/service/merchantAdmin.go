package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var MerchantAdmin = new(MADM)

type MADM struct{}

// CreateMerchantAdmin 创建后台用户与商户关联映射记录
// Author [yourname](https://github.com/yourname)
func (s *MADM) CreateMerchantAdmin(ctx context.Context, MADM *model.MerchantAdmin) (err error) {
	err = global.GVA_DB.Create(MADM).Error
	return err
}

// DeleteMerchantAdmin 删除后台用户与商户关联映射记录
// Author [yourname](https://github.com/yourname)
func (s *MADM) DeleteMerchantAdmin(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.MerchantAdmin{}, "id = ?", ID).Error
	return err
}

// DeleteMerchantAdminByIds 批量删除后台用户与商户关联映射记录
// Author [yourname](https://github.com/yourname)
func (s *MADM) DeleteMerchantAdminByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.MerchantAdmin{}, "id in ?", IDs).Error
	return err
}

// UpdateMerchantAdmin 更新后台用户与商户关联映射记录
// Author [yourname](https://github.com/yourname)
func (s *MADM) UpdateMerchantAdmin(ctx context.Context, MADM model.MerchantAdmin) (err error) {
	err = global.GVA_DB.Model(&model.MerchantAdmin{}).Where("id = ?", MADM.ID).Updates(&MADM).Error
	return err
}

// GetMerchantAdmin 根据ID获取后台用户与商户关联映射记录
// Author [yourname](https://github.com/yourname)
func (s *MADM) GetMerchantAdmin(ctx context.Context, ID string) (MADM model.MerchantAdmin, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&MADM).Error
	return
}

// GetMerchantIDByUserID 根据后台用户ID获取绑定的商户ID
func (s *MADM) GetMerchantIDByUserID(ctx context.Context, userID uint) (merchantID *int, err error) {
	var rel model.MerchantAdmin
	err = global.GVA_DB.WithContext(ctx).Where("user_id = ?", userID).First(&rel).Error
	if err != nil {
		return nil, err
	}
	return rel.MerchantID, nil
}

// GetMerchantAdminInfoList 分页获取后台用户与商户关联映射记录
// Author [yourname](https://github.com/yourname)
func (s *MADM) GetMerchantAdminInfoList(ctx context.Context, info request.MerchantAdminSearch) (list []model.MerchantAdmin, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.MerchantAdmin{})
	var MADMs []model.MerchantAdmin
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.UserID != nil {
		db = db.Where("user_id = ?", *info.UserID)
	}
	if info.MerchantID != nil {
		db = db.Where("merchant_id = ?", *info.MerchantID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["user_id"] = true
	orderMap["merchant_id"] = true
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
	err = db.Find(&MADMs).Error
	return MADMs, total, err
}
func (s *MADM) GetMerchantAdminDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	merchantId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_merchants").Where("deleted_at IS NULL").Select("merchant_name as label,id as value").Scan(&merchantId)
	res["merchantId"] = merchantId
	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_users").Where("deleted_at IS NULL").Select("username as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}

func (s *MADM) GetMerchantAdminPublic(ctx context.Context) {

}
