package service

import (
	"context"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var UserAddress = new(UA)

type UA struct{}

// CreateUserAddress 创建用户收货地址记录
// Author [yourname](https://github.com/yourname)
func (s *UA) CreateUserAddress(ctx context.Context, UA *model.UserAddress) (err error) {
	// 若设置为默认，取消当前用户其他默认地址
	if UA.IsDefault != nil && *UA.IsDefault && UA.UserID != nil {
		err = global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			if e := tx.Model(&model.UserAddress{}).
				Where("user_id = ?", *UA.UserID).
				Update("is_default", false).Error; e != nil {
				return e
			}
			return tx.Create(UA).Error
		})
		return err
	}
	err = global.GVA_DB.WithContext(ctx).Create(UA).Error
	return err
}

// DeleteUserAddress 删除用户收货地址记录
// Author [yourname](https://github.com/yourname)
func (s *UA) DeleteUserAddress(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.UserAddress{}, "id = ?", ID).Error
	return err
}

// DeleteUserAddressByIds 批量删除用户收货地址记录
// Author [yourname](https://github.com/yourname)
func (s *UA) DeleteUserAddressByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.UserAddress{}, "id in ?", IDs).Error
	return err
}

// UpdateUserAddress 更新用户收货地址记录
// Author [yourname](https://github.com/yourname)
func (s *UA) UpdateUserAddress(ctx context.Context, UA model.UserAddress) (err error) {
	// 若设置为默认，取消当前用户其他默认地址
	if UA.IsDefault != nil && *UA.IsDefault && UA.UserID != nil {
		err = global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			if e := tx.Model(&model.UserAddress{}).
				Where("user_id = ? AND id <> ?", *UA.UserID, UA.ID).
				Update("is_default", false).Error; e != nil {
				return e
			}
			return tx.Model(&model.UserAddress{}).Where("id = ?", UA.ID).Updates(&UA).Error
		})
		return err
	}
	err = global.GVA_DB.WithContext(ctx).Model(&model.UserAddress{}).Where("id = ?", UA.ID).Updates(&UA).Error
	return err
}

// GetUserAddress 根据ID获取用户收货地址记录
// Author [yourname](https://github.com/yourname)
func (s *UA) GetUserAddress(ctx context.Context, ID string) (UA model.UserAddress, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&UA).Error
	return
}

// GetUserAddressInfoList 分页获取用户收货地址记录
// Author [yourname](https://github.com/yourname)
func (s *UA) GetUserAddressInfoList(ctx context.Context, info request.UserAddressSearch) (list []model.UserAddress, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.UserAddress{})
	var UAs []model.UserAddress
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.UserID != nil {
		db = db.Where("user_id = ?", *info.UserID)
	}
	if info.Receiver != nil && *info.Receiver != "" {
		db = db.Where("receiver LIKE ?", "%"+*info.Receiver+"%")
	}
	if info.Phone != nil && *info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+*info.Phone+"%")
	}
	if info.Country != nil && *info.Country != "" {
		db = db.Where("country LIKE ?", "%"+*info.Country+"%")
	}
	if info.Province != nil && *info.Province != "" {
		db = db.Where("province LIKE ?", "%"+*info.Province+"%")
	}
	if info.City != nil && *info.City != "" {
		db = db.Where("city LIKE ?", "%"+*info.City+"%")
	}
	if info.District != nil && *info.District != "" {
		db = db.Where("district LIKE ?", "%"+*info.District+"%")
	}
	if info.Address != nil && *info.Address != "" {
		db = db.Where("address LIKE ?", "%"+*info.Address+"%")
	}
	if info.Zipcode != nil && *info.Zipcode != "" {
		db = db.Where("zipcode LIKE ?", "%"+*info.Zipcode+"%")
	}
	if info.IsDefault != nil {
		db = db.Where("is_default = ?", *info.IsDefault)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
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
	err = db.Find(&UAs).Error
	return UAs, total, err
}

func (s *UA) GetUserAddressPublic(ctx context.Context) {

}
