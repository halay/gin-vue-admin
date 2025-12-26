package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

type DPR struct{}

// GetDownlinePurchaseRecordList 分页获取下级购买记录列表 (后台/商户)
func (s *DPR) GetDownlinePurchaseRecordList(ctx context.Context, info request.DownlinePurchaseRecordSearch) (list []model.DownlinePurchaseRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.WithContext(ctx).Model(&model.DownlinePurchaseRecord{})

	// 查询条件
	if info.MerchantId != nil {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.OrderNo != nil && *info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+*info.OrderNo+"%")
	}
	if info.UplineUserId != nil {
		db = db.Where("upline_user_id = ?", info.UplineUserId)
	}
	if info.PurchaseUserId != nil {
		db = db.Where("purchase_user_id = ?", info.PurchaseUserId)
	}
	if info.IsDirect != nil {
		db = db.Where("is_direct = ?", info.IsDirect)
	}
	if info.PayStatus != nil && *info.PayStatus != "" {
		db = db.Where("pay_status = ?", *info.PayStatus)
	}
	if info.UplineUserEmail != nil && *info.UplineUserEmail != "" {
		db = db.Where("upline_user_email LIKE ?", "%"+*info.UplineUserEmail+"%")
	}
	if info.PurchaseUserEmail != nil && *info.PurchaseUserEmail != "" {
		db = db.Where("purchase_user_email LIKE ?", "%"+*info.PurchaseUserEmail+"%")
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&list).Error
	return
}

// GetMyDownlinePurchaseRecordList 分页获取我的下级购买记录列表 (APP端)
func (s *DPR) GetMyDownlinePurchaseRecordList(ctx context.Context, userID uint, info request.DownlinePurchaseRecordSearch) (list []model.DownlinePurchaseRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.WithContext(ctx).Model(&model.DownlinePurchaseRecord{})

	// 强制限制只能看自己的下级记录
	db = db.Where("upline_user_id = ?", userID)

	// 其他查询条件
	if info.OrderNo != nil && *info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+*info.OrderNo+"%")
	}
	if info.OrderStatus != nil && *info.OrderStatus != "" {
		db = db.Where("order_status = ?", *info.OrderStatus)
	}
	if info.PurchaseUserId != nil {
		db = db.Where("purchase_user_id = ?", info.PurchaseUserId)
	}
	if info.IsDirect != nil {
		db = db.Where("is_direct = ?", info.IsDirect)
	}
	if info.PayStatus != nil && *info.PayStatus != "" {
		db = db.Where("pay_status = ?", *info.PayStatus)
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&list).Error
	return
}

// CreateRecordsForOrder 为订单创建下级购买记录
func (s *DPR) CreateRecordsForOrder(ctx context.Context, order model.Order) error {
	if order.UserID == nil {
		return nil
	}

	// 1. 获取购买用户信息
	var user model.AppUsers
	// order.UserID is *int64, user.ID is uint
	if err := global.GVA_DB.WithContext(ctx).First(&user, uint(*order.UserID)).Error; err != nil {
		return err
	}

	//拼接上自己的uid放到第一个
	path := strconv.Itoa(int(user.ID)) + "/" + *user.InvitePath
	user.InvitePath = &path
	// 2. 检查是否有邀请路径
	if user.InvitePath == nil || *user.InvitePath == "" {
		return nil
	}

	// 3. 解析邀请路径
	// 路径格式如 "12/21/31"，最后一个是直属上级
	pathIds := strings.Split(*user.InvitePath, "/")
	if len(pathIds) == 0 {
		return nil
	}

	records := make([]model.DownlinePurchaseRecord, 0, len(pathIds))

	// 3.5 批量查询上级用户信息(主要是邮箱)
	uplineIDs := make([]uint, 0, len(pathIds))
	for _, uidStr := range pathIds {
		if uidStr == "" {
			continue
		}
		if uid, err := strconv.ParseUint(uidStr, 10, 64); err == nil {
			uplineIDs = append(uplineIDs, uint(uid))
		}
	}

	uplineEmailMap := make(map[uint]string)
	if len(uplineIDs) > 0 {
		var uplineUsers []model.AppUsers
		// Select id and email
		if err := global.GVA_DB.WithContext(ctx).Select("id, email").Where("id IN ?", uplineIDs).Find(&uplineUsers).Error; err == nil {
			for _, u := range uplineUsers {
				if u.Email != nil {
					uplineEmailMap[u.ID] = *u.Email
				}
			}
		}
	}

	// 4. 获取商户信息 (一次性)
	var merchantID *uint
	var merchantName *string
	if order.MerchantID != nil && *order.MerchantID > 0 {
		mid := uint(*order.MerchantID)
		merchantID = &mid
		var m model.Merchants
		if e := global.GVA_DB.WithContext(ctx).First(&m, mid).Error; e == nil {
			merchantName = m.MerchantName
		}
	}

	for i, uidStr := range pathIds {
		if uidStr == "" {
			continue
		}
		uid, err := strconv.ParseUint(uidStr, 10, 64)
		if err != nil {
			continue
		}
		uID := uint(uid)

		// 判断是否直属上级 (路径中最后一个ID)
		isDirect := false
		if i == len(pathIds)-1 {
			isDirect = true
		}

		// 计算层级深度 (倒数第几个)
		// 比如 12/21/31，31是1级(直属)，21是2级...
		inviteLevel := len(pathIds) - i

		oid := order.ID

		var uplineEmail *string
		if email, ok := uplineEmailMap[uID]; ok {
			uplineEmail = &email
		}

		record := model.DownlinePurchaseRecord{
			OrderId:              &oid,
			OrderNo:              order.OrderNo,
			PurchaseUserId:       &user.ID,
			PurchaseUserNickname: user.Nickname,
			PurchaseUserPhone:    user.Phone,
			PurchaseUserEmail:    user.Email,
			MerchantId:           merchantID,
			MerchantName:         merchantName,
			UplineUserId:         &uID,
			UplineUserEmail:      uplineEmail,
			IsDirect:             &isDirect,
			InviteLevel:          &inviteLevel,
			TotalAmount:          order.TotalAmount,
			TotalPoints:          order.TotalPoints,
			PayStatus:            order.PayStatus,
			OrderStatus:          order.Status,
		}

		records = append(records, record)
	}

	if len(records) > 0 {
		return global.GVA_DB.WithContext(ctx).Create(&records).Error
	}

	return nil
}

// UpdateStatus 更新记录状态
func (s *DPR) UpdateStatus(ctx context.Context, orderNo string, payStatus string, orderStatus string) error {
	return global.GVA_DB.WithContext(ctx).Model(&model.DownlinePurchaseRecord{}).
		Where("order_no = ?", orderNo).
		Updates(map[string]interface{}{
			"pay_status":   payStatus,
			"order_status": orderStatus,
		}).Error
}
