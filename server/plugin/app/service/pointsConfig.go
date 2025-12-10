package service

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var PointsCfg = new(PCFG)

type PCFG struct{}

func (s *PCFG) CreatePointsConfig(ctx context.Context, m *model.PointsConfig) (err error) {
    return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
        // 状态规范化：兼容字典返回的数值型字符串
        if m.Status != nil {
            switch *m.Status {
            case "1":
                v := "enabled"
                m.Status = &v
            case "0", "2":
                v := "disabled"
                m.Status = &v
            }
        }
        // 首单优惠唯一性
        if m.Type != nil && *m.Type == "first_order" {
            var cnt int64
            if e := tx.Model(&model.PointsConfig{}).Where("type = ?", "first_order").Count(&cnt).Error; e != nil {
                return e
            }
            if cnt > 0 {
                return errors.New("首单优惠配置已存在")
            }
        }
        return tx.Create(m).Error
    })
}

func (s *PCFG) DeletePointsConfig(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.PointsConfig{}, "id = ?", ID).Error
	return err
}

func (s *PCFG) DeletePointsConfigByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.PointsConfig{}, "id in ?", IDs).Error
	return err
}

func (s *PCFG) UpdatePointsConfig(ctx context.Context, m model.PointsConfig) (err error) {
    return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
        // 状态规范化
        if m.Status != nil {
            switch *m.Status {
            case "1":
                v := "enabled"
                m.Status = &v
            case "0", "2":
                v := "disabled"
                m.Status = &v
            }
        }
        if m.Type != nil && *m.Type == "first_order" {
            var cnt int64
            if e := tx.Model(&model.PointsConfig{}).Where("type = ? AND id <> ?", "first_order", m.ID).Count(&cnt).Error; e != nil {
                return e
            }
            if cnt > 0 {
                return errors.New("首单优惠配置已存在")
            }
        }
        return tx.Model(&model.PointsConfig{}).Where("id = ?", m.ID).Updates(&m).Error
    })
}

func (s *PCFG) GetPointsConfig(ctx context.Context, ID string) (m model.PointsConfig, err error) {
	err = global.GVA_DB.First(&m, "id = ?", ID).Error
	return
}

func (s *PCFG) GetPointsConfigInfoList(ctx context.Context, info request.PointsConfigSearch) (list []model.PointsConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.PointsConfig{})
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.Type != nil && *info.Type != "" {
		db = db.Where("type = ?", *info.Type)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.Limited != nil {
		db = db.Where("limited = ?", *info.Limited)
	}
	if len(info.StartAtRange) == 2 {
		db = db.Where("start_at BETWEEN ? AND ?", info.StartAtRange[0], info.StartAtRange[1])
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	// 排序
	order := "sort asc, id asc"
	if info.Sort != "" {
		order = info.Sort
		if info.Order == "descending" {
			order += " desc"
		}
	}
	err = db.Order(order).Find(&list).Error
	return
}

// 公共接口：返回已启用且在时间窗口内的配置列表
func (s *PCFG) GetPointsConfigPublic(ctx context.Context, info request.PointsConfigSearch) (list []model.PointsConfig, total int64, err error) {
	info.Status = strPtr("enabled")
	return s.GetPointsConfigInfoList(ctx, info)
}
