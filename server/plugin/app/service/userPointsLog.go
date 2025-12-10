package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var UserPointsLog = new(UPL)

type UPL struct{}

// CreateUserPointsLog 创建用户积分流水记录
// Author [yourname](https://github.com/yourname)
func (s *UPL) CreateUserPointsLog(ctx context.Context, UPL *model.UserPointsLog) (err error) {
	err = global.GVA_DB.Create(UPL).Error
	return err
}

// DeleteUserPointsLog 删除用户积分流水记录
// Author [yourname](https://github.com/yourname)
func (s *UPL) DeleteUserPointsLog(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.UserPointsLog{}, "id = ?", ID).Error
	return err
}

// DeleteUserPointsLogByIds 批量删除用户积分流水记录
// Author [yourname](https://github.com/yourname)
func (s *UPL) DeleteUserPointsLogByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.UserPointsLog{}, "id in ?", IDs).Error
	return err
}

// UpdateUserPointsLog 更新用户积分流水记录
// Author [yourname](https://github.com/yourname)
func (s *UPL) UpdateUserPointsLog(ctx context.Context, UPL model.UserPointsLog) (err error) {
	err = global.GVA_DB.Model(&model.UserPointsLog{}).Where("id = ?", UPL.ID).Updates(&UPL).Error
	return err
}

// GetUserPointsLog 根据ID获取用户积分流水记录
// Author [yourname](https://github.com/yourname)
func (s *UPL) GetUserPointsLog(ctx context.Context, ID string) (UPL model.UserPointsLog, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&UPL).Error
	return
}

// GetUserPointsLogInfoList 分页获取用户积分流水记录
// Author [yourname](https://github.com/yourname)
func (s *UPL) GetUserPointsLogInfoList(ctx context.Context, info request.UserPointsLogSearch) (list []model.UserPointsLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.UserPointsLog{})
	var UPLs []model.UserPointsLog
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.UserID != nil {
		db = db.Where("user_id = ?", *info.UserID)
	}
	if info.Change != nil {
		db = db.Where("change >= ?", *info.Change)
	}
	if info.BalanceAfter != nil {
		db = db.Where("balance_after >= ?", *info.BalanceAfter)
	}
	if info.Type != nil && *info.Type != "" {
		db = db.Where("type = ?", *info.Type)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.Reason != nil && *info.Reason != "" {
		db = db.Where("reason LIKE ?", "%"+*info.Reason+"%")
	}
	if info.OrderNo != nil && *info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+*info.OrderNo+"%")
	}
	if info.Remark != nil && *info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+*info.Remark+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&UPLs).Error
	return UPLs, total, err
}
func (s *UPL) GetUserPointsLogDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_users").Where("deleted_at IS NULL").Select("email as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}

func (s *UPL) GetUserPointsLogPublic(ctx context.Context) {

}
