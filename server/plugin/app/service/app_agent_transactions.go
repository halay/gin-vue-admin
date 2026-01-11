package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var AppAgentTransactions = new(appAgentTransactions)

type appAgentTransactions struct{}

// CreateAppAgentTransactions 创建appAgentTransactions表记录
// Author [yourname](https://github.com/yourname)
func (s *appAgentTransactions) CreateAppAgentTransactions(ctx context.Context, appAgentTransactions *model.AppAgentTransactions) (err error) {
	err = global.GVA_DB.Create(appAgentTransactions).Error
	return err
}

// DeleteAppAgentTransactions 删除appAgentTransactions表记录
// Author [yourname](https://github.com/yourname)
func (s *appAgentTransactions) DeleteAppAgentTransactions(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&model.AppAgentTransactions{}, "id = ?", id).Error
	return err
}

// DeleteAppAgentTransactionsByIds 批量删除appAgentTransactions表记录
// Author [yourname](https://github.com/yourname)
func (s *appAgentTransactions) DeleteAppAgentTransactionsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.AppAgentTransactions{}, "id in ?", ids).Error
	return err
}

// UpdateAppAgentTransactions 更新appAgentTransactions表记录
// Author [yourname](https://github.com/yourname)
func (s *appAgentTransactions) UpdateAppAgentTransactions(ctx context.Context, appAgentTransactions model.AppAgentTransactions) (err error) {
	err = global.GVA_DB.Model(&model.AppAgentTransactions{}).Where("id = ?", appAgentTransactions.Id).Updates(&appAgentTransactions).Error
	return err
}

// GetAppAgentTransactions 根据id获取appAgentTransactions表记录
// Author [yourname](https://github.com/yourname)
func (s *appAgentTransactions) GetAppAgentTransactions(ctx context.Context, id string) (appAgentTransactions model.AppAgentTransactions, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appAgentTransactions).Error
	return
}

// GetAppAgentTransactionsInfoList 分页获取appAgentTransactions表记录
// Author [yourname](https://github.com/yourname)
func (s *appAgentTransactions) GetAppAgentTransactionsInfoList(ctx context.Context, info request.AppAgentTransactionsSearch) (list []model.AppAgentTransactions, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AppAgentTransactions{})
	var appAgentTransactionss []model.AppAgentTransactions
	// 如果有条件搜索 下方会自动创建搜索语句
	// 查询条件
	if info.MerchantId != nil {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.BeneficiaryId != nil {
		db = db.Where("beneficiary_id = ?", info.BeneficiaryId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&appAgentTransactionss).Error
	return appAgentTransactionss, total, err
}

func (s *appAgentTransactions) GetAppAgentTransactionsPublic(ctx context.Context) {

}
