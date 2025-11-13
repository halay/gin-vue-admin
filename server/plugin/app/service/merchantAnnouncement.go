package service

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var MerchantAnnouncement = new(MA)

type MA struct{}

// CreateMerchantAnnouncement 创建商户公告信息记录
// Author [yourname](https://github.com/yourname)
func (s *MA) CreateMerchantAnnouncement(ctx context.Context, MA *model.MerchantAnnouncement, merchantID int) (err error) {
	MA.MerchantID = &merchantID
	err = global.GVA_DB.WithContext(ctx).Create(MA).Error
	return err
}

// DeleteMerchantAnnouncement 删除商户公告信息记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *MA) DeleteMerchantAnnouncement(ctx context.Context, ID string, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&model.MerchantAnnouncement{}, "id = ?", ID).Error
	return err
}

// DeleteMerchantAnnouncementByIds 批量删除商户公告信息记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *MA) DeleteMerchantAnnouncementByIds(ctx context.Context, IDs []string, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&[]model.MerchantAnnouncement{}, "id in ?", IDs).Error
	return err
}

// UpdateMerchantAnnouncement 更新商户公告信息记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *MA) UpdateMerchantAnnouncement(ctx context.Context, MA model.MerchantAnnouncement, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&model.MerchantAnnouncement{}).Where("id = ? AND merchant_id = ?", MA.ID, merchantID).Updates(&MA).Error
	return err
}

// GetMerchantAnnouncement 根据ID获取商户公告信息记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *MA) GetMerchantAnnouncement(ctx context.Context, ID string, merchantID int) (MA model.MerchantAnnouncement, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ? AND merchant_id = ?", ID, merchantID).First(&MA).Error
	return
}

// GetMerchantAnnouncementInfoList 分页获取商户公告信息记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *MA) GetMerchantAnnouncementInfoList(ctx context.Context, info request.MerchantAnnouncementSearch, merchantID int) (list []model.MerchantAnnouncement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&model.MerchantAnnouncement{})
	var MAs []model.MerchantAnnouncement
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	if len(info.PublishAtRange) == 2 {
		db = db.Where("publish_at BETWEEN ? AND ? ", info.PublishAtRange[0], info.PublishAtRange[1])
	}
	db = db.Where("merchant_id = ?", merchantID)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["title"] = true
	orderMap["status"] = true
	orderMap["publish_at"] = true
	orderMap["expire_at"] = true
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
	err = db.Find(&MAs).Error
	return MAs, total, err
}
func (s *MA) GetMerchantAnnouncementDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	merchantId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_merchants").Where("deleted_at IS NULL").Select("merchant_name as label,id as value").Scan(&merchantId)
	res["merchantId"] = merchantId
	return
}

func (s *MA) GetMerchantAnnouncementPublic(ctx context.Context) {

}
