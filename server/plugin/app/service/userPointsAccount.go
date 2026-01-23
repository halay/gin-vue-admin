package service

import (
	"context"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var UserPointsAccount = new(UPA)

type UPA struct{}

// CreateUserPointsAccount 创建用户积分账户记录
// Author [yourname](https://github.com/yourname)
func (s *UPA) CreateUserPointsAccount(ctx context.Context, UPA *model.UserPointsAccount) (err error) {
	err = global.GVA_DB.Create(UPA).Error
	return err
}

// DeleteUserPointsAccount 删除用户积分账户记录
// Author [yourname](https://github.com/yourname)
func (s *UPA) DeleteUserPointsAccount(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.UserPointsAccount{}, "id = ?", ID).Error
	return err
}

// DeleteUserPointsAccountByIds 批量删除用户积分账户记录
// Author [yourname](https://github.com/yourname)
func (s *UPA) DeleteUserPointsAccountByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.UserPointsAccount{}, "id in ?", IDs).Error
	return err
}

// UpdateUserPointsAccount 更新用户积分账户记录
// Author [yourname](https://github.com/yourname)
func (s *UPA) UpdateUserPointsAccount(ctx context.Context, UPA model.UserPointsAccount) (err error) {
	err = global.GVA_DB.Model(&model.UserPointsAccount{}).Where("id = ?", UPA.ID).Updates(&UPA).Error
	return err
}

// GetUserPointsAccount 根据ID获取用户积分账户记录
// Author [yourname](https://github.com/yourname)
func (s *UPA) GetUserPointsAccount(ctx context.Context, ID string) (UPA model.UserPointsAccount, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&UPA).Error
	return
}

// GetUserAllPointsAccounts 获取用户所有积分账户列表（包括平台和各商户）
func (s *UPA) GetUserAllPointsAccounts(ctx context.Context, userID int64) (list []model.UserPointsAccount, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("user_id = ?", userID).Find(&list).Error
	return
}

func (s *UPA) GetUserPointsAccountList(ctx context.Context, info request.UserPointsAccountSearch) (list []model.UserPointsAccount, total int64, err error) {
	// GetUserPointsAccountInfoList 分页获取用户积分账户记录
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.UserPointsAccount{})
	var UPAs []model.UserPointsAccount
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.UserID != nil {
		db = db.Where("user_id = ?", *info.UserID)
	}
	if info.Balance != nil {
		db = db.Where("balance >= ?", *info.Balance)
	}
	if info.MerchantID != nil {
		db = db.Where("merchant_id = ?", *info.MerchantID)
	}
	db = db.Order("created_at desc")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&UPAs).Error
	return UPAs, total, err
}
func (s *UPA) GetUserPointsAccountDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_users").Where("deleted_at IS NULL").Select("email as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}

func (s *UPA) GetUserPointsAccountPublic(ctx context.Context) {

}

// EnsureAccount 确保用户积分账户存在
func (s *UPA) EnsureAccount(ctx context.Context, userID int64, merchantID int64) (acc model.UserPointsAccount, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("user_id = ? AND merchant_id = ?", userID, merchantID).First(&acc).Error
	if err == gorm.ErrRecordNotFound {
		acc = model.UserPointsAccount{
			UserID:     &userID,
			MerchantID: &merchantID,
		}
		// 获取TokenName和Symbol
		if merchantID == 0 {
			var ps model.PointsSettings
			if e := global.GVA_DB.WithContext(ctx).First(&ps).Error; e == nil {
				acc.TokenName = ps.TokenName
				acc.Symbol = ps.Symbol
			}
		} else {
			var mps model.MerchantPointsSettings
			if e := global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).First(&mps).Error; e == nil {
				acc.TokenName = mps.TokenName
				acc.Symbol = mps.Symbol
			}
		}
		zero := int64(0)
		acc.Balance = &zero
		if e := global.GVA_DB.WithContext(ctx).Create(&acc).Error; e != nil {
			return acc, e
		}
		err = nil
	}
	return
}

// AddPoints 增加积分（原子操作），返回加后余额
func (s *UPA) AddPoints(ctx context.Context, userID int64, points int64, merchantID int64) (after int64, err error) {
	// 先确保存在账户
	if _, err = s.EnsureAccount(ctx, userID, merchantID); err != nil {
		return 0, err
	}
	tx := global.GVA_DB.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	res := tx.Model(&model.UserPointsAccount{}).
		Where("user_id = ? AND merchant_id = ?", userID, merchantID).
		UpdateColumn("balance", gorm.Expr("balance + ?", points))
	if res.Error != nil {
		err = res.Error
		return
	}
	var acc model.UserPointsAccount
	if e := tx.Where("user_id = ? AND merchant_id = ?", userID, merchantID).First(&acc).Error; e != nil {
		err = e
		return
	}
	if e := tx.Commit().Error; e != nil {
		err = e
		return
	}
	if acc.Balance != nil {
		after = *acc.Balance
	}
	return
}

// DeductPoints 扣减积分（原子操作），返回扣后余额
func (s *UPA) DeductPoints(ctx context.Context, userID int64, points int64, merchantID int64) (after int64, err error) {
	// 先确保存在账户
	if _, err = s.EnsureAccount(ctx, userID, merchantID); err != nil {
		return 0, err
	}
	tx := global.GVA_DB.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 乐观扣减
	res := tx.Model(&model.UserPointsAccount{}).
		Where("user_id = ? AND merchant_id = ? AND balance >= ?", userID, merchantID, points).
		UpdateColumn("balance", gorm.Expr("balance - ?", points))
	if res.Error != nil {
		err = res.Error
		return
	}
	if res.RowsAffected == 0 {
		err = gorm.ErrInvalidData
		return
	}
	// 查询余额
	var acc model.UserPointsAccount
	if e := tx.Where("user_id = ? AND merchant_id = ?", userID, merchantID).First(&acc).Error; e != nil {
		err = e
		return
	}
	if e := tx.Commit().Error; e != nil {
		err = e
		return
	}
	if acc.Balance != nil {
		after = *acc.Balance
	}
	return
}

// AddLog 记录积分流水
func (s *UPA) AddLog(ctx context.Context, userID int64, change int64, balanceAfter int64, reason, orderNo, remark string, merchantID int64) error {
	l := model.UserPointsLog{
		UserID:       &userID,
		Change:       &change,
		BalanceAfter: &balanceAfter,
		Reason:       &reason,
		OrderNo:      &orderNo,
		Remark:       &remark,
		MerchantID:   &merchantID,
	}

	// 补充Token信息
	if merchantID == 0 {
		var ps model.PointsSettings
		if e := global.GVA_DB.WithContext(ctx).First(&ps).Error; e == nil {
			l.TokenName = ps.TokenName
			l.Symbol = ps.Symbol
		}
	} else {
		var mps model.MerchantPointsSettings
		if e := global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).First(&mps).Error; e == nil {
			l.TokenName = mps.TokenName
			l.Symbol = mps.Symbol
		}
	}

	return global.GVA_DB.WithContext(ctx).Create(&l).Error
}

// AddLogDetailed 记录积分流水（含类型、状态、关联ID）
func (s *UPA) AddLogDetailed(ctx context.Context, userID int64, change int64, balanceAfter int64, reason, orderNo, remark, logType, status string, relatedID *int64, merchantID *int64, inviterUserID *int64, inviterEmail *string, inviteeUserID *int64, inviteeEmail *string, isRecharge *bool) error {
	mID := int64(0)
	if merchantID != nil {
		mID = *merchantID
	}

	l := model.UserPointsLog{
		UserID:        &userID,
		Change:        &change,
		BalanceAfter:  &balanceAfter,
		Type:          &logType,
		Status:        &status,
		RelatedID:     relatedID,
		MerchantID:    merchantID,
		Reason:        &reason,
		OrderNo:       &orderNo,
		Remark:        &remark,
		InviterUserID: inviterUserID,
		InviterEmail:  inviterEmail,
		InviteeUserID: inviteeUserID,
		InviteeEmail:  inviteeEmail,
		IsRecharge:    isRecharge,
	}

	// 补充Token信息
	if mID == 0 {
		var ps model.PointsSettings
		if e := global.GVA_DB.WithContext(ctx).First(&ps).Error; e == nil {
			l.TokenName = ps.TokenName
			l.Symbol = ps.Symbol
		}
	} else {
		var mps model.MerchantPointsSettings
		if e := global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", mID).First(&mps).Error; e == nil {
			l.TokenName = mps.TokenName
			l.Symbol = mps.Symbol
		}
	}

	return global.GVA_DB.WithContext(ctx).Create(&l).Error
}
