
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var PointsSettings = new(PTS)

type PTS struct {}
// CreatePointsSettings 创建平台积分基础设置与奖励规则记录
// Author [yourname](https://github.com/yourname)
func (s *PTS) CreatePointsSettings(ctx context.Context, PTS *model.PointsSettings) (err error) {
	err = global.GVA_DB.Create(PTS).Error
	return err
}

// DeletePointsSettings 删除平台积分基础设置与奖励规则记录
// Author [yourname](https://github.com/yourname)
func (s *PTS) DeletePointsSettings(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.PointsSettings{},"id = ?",ID).Error
	return err
}

// DeletePointsSettingsByIds 批量删除平台积分基础设置与奖励规则记录
// Author [yourname](https://github.com/yourname)
func (s *PTS) DeletePointsSettingsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.PointsSettings{},"id in ?",IDs).Error
	return err
}

// UpdatePointsSettings 更新平台积分基础设置与奖励规则记录
// Author [yourname](https://github.com/yourname)
func (s *PTS) UpdatePointsSettings(ctx context.Context, PTS model.PointsSettings) (err error) {
	err = global.GVA_DB.Model(&model.PointsSettings{}).Where("id = ?",PTS.ID).Updates(&PTS).Error
	return err
}

// GetPointsSettings 根据ID获取平台积分基础设置与奖励规则记录
// Author [yourname](https://github.com/yourname)
func (s *PTS) GetPointsSettings(ctx context.Context, ID string) (PTS model.PointsSettings, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&PTS).Error
	return
}
// GetPointsSettingsInfoList 分页获取平台积分基础设置与奖励规则记录
// Author [yourname](https://github.com/yourname)
func (s *PTS) GetPointsSettingsInfoList(ctx context.Context, info request.PointsSettingsSearch) (list []model.PointsSettings, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PointsSettings{})
    var PTSs []model.PointsSettings
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.TokenName != nil && *info.TokenName != "" {
        db = db.Where("token_name LIKE ?", "%"+ *info.TokenName+"%")
    }
    if info.Symbol != nil && *info.Symbol != "" {
        db = db.Where("symbol LIKE ?", "%"+ *info.Symbol+"%")
    }
    if info.RegisterReward != nil {
        db = db.Where("register_reward >= ?", *info.RegisterReward)
    }
    if info.InviterReward != nil {
        db = db.Where("inviter_reward >= ?", *info.InviterReward)
    }
    if info.InviteeReward != nil {
        db = db.Where("invitee_reward >= ?", *info.InviteeReward)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
    if info.SortValue != nil {
        db = db.Where("sort = ?", *info.SortValue)
    }
    if info.Remark != nil && *info.Remark != "" {
        db = db.Where("remark LIKE ?", "%"+ *info.Remark+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["token_name"] = true
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
	err = db.Find(&PTSs).Error
	return  PTSs, total, err
}

func (s *PTS)GetPointsSettingsPublic(ctx context.Context) {

}
