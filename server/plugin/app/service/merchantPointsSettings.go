
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var MerchantPointsSettings = new(MPTS)

type MPTS struct {}
// CreateMerchantPointsSettings 创建商户积分营销配置记录
// Author [yourname](https://github.com/yourname)
func (s *MPTS) CreateMerchantPointsSettings(ctx context.Context, MPTS *model.MerchantPointsSettings) (err error) {
	err = global.GVA_DB.Create(MPTS).Error
	return err
}

// DeleteMerchantPointsSettings 删除商户积分营销配置记录
// Author [yourname](https://github.com/yourname)
func (s *MPTS) DeleteMerchantPointsSettings(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.MerchantPointsSettings{},"id = ?",ID).Error
	return err
}

// DeleteMerchantPointsSettingsByIds 批量删除商户积分营销配置记录
// Author [yourname](https://github.com/yourname)
func (s *MPTS) DeleteMerchantPointsSettingsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.MerchantPointsSettings{},"id in ?",IDs).Error
	return err
}

// UpdateMerchantPointsSettings 更新商户积分营销配置记录
// Author [yourname](https://github.com/yourname)
func (s *MPTS) UpdateMerchantPointsSettings(ctx context.Context, MPTS model.MerchantPointsSettings) (err error) {
	err = global.GVA_DB.Model(&model.MerchantPointsSettings{}).Where("id = ?",MPTS.ID).Updates(&MPTS).Error
	return err
}

// GetMerchantPointsSettings 根据ID获取商户积分营销配置记录
// Author [yourname](https://github.com/yourname)
func (s *MPTS) GetMerchantPointsSettings(ctx context.Context, ID string) (MPTS model.MerchantPointsSettings, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&MPTS).Error
	return
}
// GetMerchantPointsSettingsInfoList 分页获取商户积分营销配置记录
// Author [yourname](https://github.com/yourname)
func (s *MPTS) GetMerchantPointsSettingsInfoList(ctx context.Context, info request.MerchantPointsSettingsSearch) (list []model.MerchantPointsSettings, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.MerchantPointsSettings{})
    var MPTSs []model.MerchantPointsSettings
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.MerchantID != nil {
        db = db.Where("merchant_id = ?", *info.MerchantID)
    }
    if info.TokenName != nil && *info.TokenName != "" {
        db = db.Where("token_name LIKE ?", "%"+ *info.TokenName+"%")
    }
    if info.Symbol != nil && *info.Symbol != "" {
        db = db.Where("symbol LIKE ?", "%"+ *info.Symbol+"%")
    }
    if info.InviterReward != nil {
        db = db.Where("inviter_reward >= ?", *info.InviterReward)
    }
    if info.InviteeReward != nil {
        db = db.Where("invitee_reward >= ?", *info.InviteeReward)
    }
    if info.RegisterReward != nil {
        db = db.Where("register_reward >= ?", *info.RegisterReward)
    }
    if info.CoverageMode != nil && *info.CoverageMode != "" {
        db = db.Where("coverage_mode = ?", *info.CoverageMode)
    }
    if info.DailyLimit != nil {
        db = db.Where("daily_limit >= ?", *info.DailyLimit)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
    if info.Sort != nil {
        db = db.Where("sort = ?", *info.Sort)
    }
    if info.Remark != nil && *info.Remark != "" {
        db = db.Where("remark LIKE ?", "%"+ *info.Remark+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&MPTSs).Error
	return  MPTSs, total, err
}

func (s *MPTS)GetMerchantPointsSettingsPublic(ctx context.Context) {

}
