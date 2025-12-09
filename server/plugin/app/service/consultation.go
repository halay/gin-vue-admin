package service

import (
	"context"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/response"
)

var Consultation = new(CN)

type CN struct{}

func (s *CN) CreateConsultation(ctx context.Context, m *model.AppConsultation) (err error) {
	err = global.GVA_DB.WithContext(ctx).Create(m).Error
	return err
}

func (s *CN) DeleteConsultation(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&model.AppConsultation{}, "id = ?", ID).Error
	return err
}

func (s *CN) DeleteConsultationByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.WithContext(ctx).Delete(&[]model.AppConsultation{}, "id in ?", IDs).Error
	return err
}

func (s *CN) UpdateConsultation(ctx context.Context, m model.AppConsultation) (err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&model.AppConsultation{}).Where("id = ?", m.ID).Updates(&m).Error
	return err
}

func (s *CN) GetConsultation(ctx context.Context, ID string) (m model.AppConsultation, err error) {
	err = global.GVA_DB.WithContext(ctx).First(&m, "id = ?", ID).Error
	return
}

func (s *CN) GetConsultationInfoList(ctx context.Context, info request.ConsultationSearch) (list []response.ConsultationResponse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.WithContext(ctx).Model(&model.AppConsultation{})
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Category != nil && *info.Category != "" {
		db = db.Where("category = ?", *info.Category)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if len(info.PublishAtRange) == 2 {
		db = db.Where("publish_at BETWEEN ? AND ?", info.PublishAtRange[0], info.PublishAtRange[1])
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := map[string]bool{"id": true, "created_at": true, "publish_at": true, "sort": true, "clicks": true, "title": true, "status": true}
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr += " desc"
		}
		db = db.Order(OrderStr)
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&list).Error
	return
}

func (s *CN) GetConsultationPublicList(ctx context.Context, info request.ConsultationSearch) (list []response.ConsultationResponse, total int64, err error) {
	info.Status = strPtr("enabled")
	if info.Category == nil {
		info.Category = strPtr("1")
	}
	return s.GetConsultationInfoList(ctx, info)
}

func (s *CN) GetConsultationPublicDetail(ctx context.Context, ID string) (m model.AppConsultation, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ? AND status = ?", ID, "enabled").First(&m).Error
	if err == nil {
		_ = global.GVA_DB.WithContext(ctx).Model(&model.AppConsultation{}).Where("id = ?", ID).UpdateColumn("clicks", gorm.Expr("clicks + 1")).Error
	}
	return
}

func strPtr(s string) *string { return &s }
