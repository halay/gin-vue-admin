package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var Banner = new(banner)

type banner struct{}

// CreateBanner 创建Banner图片记录
// Author [yourname](https://github.com/yourname)
func (s *banner) CreateBanner(ctx context.Context, banner *model.Banner) (err error) {
	err = global.GVA_DB.Create(banner).Error
	return err
}

// DeleteBanner 删除Banner图片记录
// Author [yourname](https://github.com/yourname)
func (s *banner) DeleteBanner(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Banner{}, "id = ?", ID).Error
	return err
}

// DeleteBannerByIds 批量删除Banner图片记录
// Author [yourname](https://github.com/yourname)
func (s *banner) DeleteBannerByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Banner{}, "id in ?", IDs).Error
	return err
}

// UpdateBanner 更新Banner图片记录
// Author [yourname](https://github.com/yourname)
func (s *banner) UpdateBanner(ctx context.Context, banner model.Banner) (err error) {
	err = global.GVA_DB.Model(&model.Banner{}).Where("id = ?", banner.ID).Updates(&banner).Error
	return err
}

// GetBanner 根据ID获取Banner图片记录
// Author [yourname](https://github.com/yourname)
func (s *banner) GetBanner(ctx context.Context, ID string) (banner model.Banner, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&banner).Error
	return
}

// GetBannerInfoList 分页获取Banner图片记录
// Author [yourname](https://github.com/yourname)
func (s *banner) GetBannerInfoList(ctx context.Context, info request.BannerSearch) (list []model.Banner, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Banner{})
	var banners []model.Banner
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Position != nil && *info.Position != "" {
		db = db.Where("position = ?", *info.Position)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if len(info.StartTimeRange) == 2 {
		db = db.Where("start_time BETWEEN ? AND ? ", info.StartTimeRange[0], info.StartTimeRange[1])
	}
	if len(info.EndTimeRange) == 2 {
		db = db.Where("end_time BETWEEN ? AND ? ", info.EndTimeRange[0], info.EndTimeRange[1])
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["sort"] = true
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
	err = db.Find(&banners).Error
	return banners, total, err
}

func (s *banner) GetBannerPublic(ctx context.Context, position string) (list []model.Banner, err error) {
	db := global.GVA_DB.WithContext(ctx).Model(&model.Banner{})
	now := global.GVA_DB.NowFunc()
	db = db.Where("status = ?", "enabled")
	db = db.Where("(start_time IS NULL OR start_time <= ?)", now)
	db = db.Where("(end_time IS NULL OR end_time >= ?)", now)
	if position != "" {
		db = db.Where("position = ?", position)
	}
	db = db.Order("sort asc").Order("created_at desc")
	err = db.Find(&list).Error
	return
}
