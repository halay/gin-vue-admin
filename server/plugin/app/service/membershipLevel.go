package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var MembershipLevel = new(ML)

type ML struct{}

// CreateMembershipLevel 创建会员等级记录
// Author [yourname](https://github.com/yourname)
func (s *ML) CreateMembershipLevel(ctx context.Context, ML *model.MembershipLevel) (err error) {
	err = global.GVA_DB.Create(ML).Error
	return err
}

// DeleteMembershipLevel 删除会员等级记录
// Author [yourname](https://github.com/yourname)
func (s *ML) DeleteMembershipLevel(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.MembershipLevel{}, "id = ?", ID).Error
	return err
}

// DeleteMembershipLevelByIds 批量删除会员等级记录
// Author [yourname](https://github.com/yourname)
func (s *ML) DeleteMembershipLevelByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.MembershipLevel{}, "id in ?", IDs).Error
	return err
}

// UpdateMembershipLevel 更新会员等级记录
// Author [yourname](https://github.com/yourname)
func (s *ML) UpdateMembershipLevel(ctx context.Context, ML model.MembershipLevel) (err error) {
	err = global.GVA_DB.Model(&model.MembershipLevel{}).Where("id = ?", ML.ID).Updates(&ML).Error
	return err
}

// GetMembershipLevel 根据ID获取会员等级记录
// Author [yourname](https://github.com/yourname)
func (s *ML) GetMembershipLevel(ctx context.Context, ID string) (ML model.MembershipLevel, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ML).Error
	return
}

// GetMembershipLevelInfoByCode
func (s *ML) GetMembershipLevelInfoByCode(ctx context.Context, code string) (ML model.MembershipLevel, err error) {
	err = global.GVA_DB.Where("code = ?", code).First(&ML).Error
	return
}

// GetMembershipLevelInfoList 分页获取会员等级记录
// Author [yourname](https://github.com/yourname)
func (s *ML) GetMembershipLevelInfoList(ctx context.Context, info request.MembershipLevelSearch) (list []model.MembershipLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.MembershipLevel{})
	var MLs []model.MembershipLevel
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Name != nil && *info.Name != "" {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}
	if info.Code != nil && *info.Code != "" {
		db = db.Where("code LIKE ?", "%"+*info.Code+"%")
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.SortValue != nil {
		db = db.Where("sort = ?", *info.SortValue)
	}
	if info.Requirement != nil && *info.Requirement != "" {
		db = db.Where("requirement LIKE ?", "%"+*info.Requirement+"%")
	}
	if info.RequirementValue != nil {
		db = db.Where("requirement_value >= ?", *info.RequirementValue)
	}
	if info.Benefits != "" {
		// TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
	}
	err = db.Count(&total).Error
	if err != nil {
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
	err = db.Find(&MLs).Error
	return MLs, total, err
}

func (s *ML) GetMembershipLevelPublic(ctx context.Context) {

}

func (s *ML) ListEnabled(ctx context.Context) (list []model.MembershipLevel, err error) {
	err = global.GVA_DB.Where("status IN (?)", []string{"1", "enabled"}).Order("sort asc, id desc").Find(&list).Error
	return
}
