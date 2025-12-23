package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var AgentLevel = new(AL)

type AL struct{}

// CreateAgentLevel 创建营销方案设置记录
// Author [yourname](https://github.com/yourname)
func (s *AL) CreateAgentLevel(ctx context.Context, AL *model.AgentLevel) (err error) {
	err = global.GVA_DB.Create(AL).Error
	return err
}

// DeleteAgentLevel 删除营销方案设置记录
// Author [yourname](https://github.com/yourname)
func (s *AL) DeleteAgentLevel(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&model.AgentLevel{}, "id = ?", id).Error
	return err
}

// DeleteAgentLevelByIds 批量删除营销方案设置记录
// Author [yourname](https://github.com/yourname)
func (s *AL) DeleteAgentLevelByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.AgentLevel{}, "id in ?", ids).Error
	return err
}

// UpdateAgentLevel 更新营销方案设置记录
// Author [yourname](https://github.com/yourname)
func (s *AL) UpdateAgentLevel(ctx context.Context, AL model.AgentLevel) (err error) {
	err = global.GVA_DB.Model(&model.AgentLevel{}).Where("id = ?", AL.Id).Updates(&AL).Error
	return err
}

// GetAgentLevel 根据id获取营销方案设置记录
// Author [yourname](https://github.com/yourname)
func (s *AL) GetAgentLevel(ctx context.Context, id string) (AL model.AgentLevel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&AL).Error
	return
}

// GetAgentLevelInfoList 分页获取营销方案设置记录
// Author [yourname](https://github.com/yourname)
func (s *AL) GetAgentLevelInfoList(ctx context.Context, info request.AgentLevelSearch) (list []model.AgentLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.AgentLevel{})
	if info.MerchantId != nil && *info.MerchantId != 0 {
		db = db.Where("merchant_id = ?", *info.MerchantId)
	}
	var ALs []model.AgentLevel
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&ALs).Error
	return ALs, total, err
}
func (s *AL) GetAgentLevelDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	merchantId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_merchants").Where("deleted_at IS NULL").Select("merchant_name as label,id as value").Scan(&merchantId)
	res["merchantId"] = merchantId
	return
}

func (s *AL) GetAgentLevelPublic(ctx context.Context) {

}
