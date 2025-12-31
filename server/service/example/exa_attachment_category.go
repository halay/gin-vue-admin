package example

import (
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

type AttachmentCategoryService struct{}

// AddCategory 创建/更新的分类
func (a *AttachmentCategoryService) AddCategory(req *example.ExaAttachmentCategory) (err error) {
	// 检查是否已存在相同名称的分类
	if (!errors.Is(global.GVA_DB.Take(&example.ExaAttachmentCategory{}, "name = ? and pid = ? and user_id = ?", req.Name, req.Pid, req.UserID).Error, gorm.ErrRecordNotFound)) {
		return errors.New("分类名称已存在")
	}
	if req.ID > 0 {
		// 校验权限：只能修改自己的分类
		var old example.ExaAttachmentCategory
		if err = global.GVA_DB.First(&old, req.ID).Error; err != nil {
			return errors.New("分类不存在")
		}
		if old.UserID != req.UserID {
			return errors.New("无权修改此分类")
		}

		if err = global.GVA_DB.Model(&example.ExaAttachmentCategory{}).Where("id = ?", req.ID).Updates(&example.ExaAttachmentCategory{
			Name: req.Name,
			Pid:  req.Pid,
		}).Error; err != nil {
			return err
		}
	} else {
		if err = global.GVA_DB.Create(&example.ExaAttachmentCategory{
			Name:       req.Name,
			Pid:        req.Pid,
			UserID:     req.UserID,
			MerchantID: req.MerchantID,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

// DeleteCategory 删除分类
func (a *AttachmentCategoryService) DeleteCategory(id *int, userID uint) error {
	var childCount int64
	global.GVA_DB.Model(&example.ExaAttachmentCategory{}).Where("pid = ?", id).Count(&childCount)
	if childCount > 0 {
		return errors.New("请先删除子级")
	}
	if err := global.GVA_DB.Where("id = ? and user_id = ?", id, userID).Unscoped().Delete(&example.ExaAttachmentCategory{}).Error; err != nil {
		return err
	}
	return nil
}

// GetCategoryList 分类列表
func (a *AttachmentCategoryService) GetCategoryList(userID uint, merchantID *int64) (res []*example.ExaAttachmentCategory, err error) {
	var fileLists []example.ExaAttachmentCategory
	global.GVA_DB.Where("user_id = ? and merchant_id = ?", userID, merchantID).Find(&fileLists)
	if err != nil {
		return res, err
	}
	return a.getChildrenList(fileLists, 0), nil
}

// getChildrenList 子类
func (a *AttachmentCategoryService) getChildrenList(categories []example.ExaAttachmentCategory, parentID uint) []*example.ExaAttachmentCategory {
	var tree []*example.ExaAttachmentCategory
	for _, category := range categories {
		if category.Pid == parentID {
			category.Children = a.getChildrenList(categories, category.ID)
			tree = append(tree, &category)
		}
	}
	return tree
}
