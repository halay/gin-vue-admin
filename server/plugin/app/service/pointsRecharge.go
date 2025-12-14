package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var PointsRecharge = new(PR)

type PR struct{}

// CreateRechargeOrder 创建积分充值订单（pending）
func (s *PR) CreateRechargeOrder(ctx context.Context, userID int64, configID uint, payMethod string) (ord model.PointsRechargeOrder, cfg model.PointsConfig, err error) {
	if err = global.GVA_DB.WithContext(ctx).First(&cfg, "id = ? AND status = ?", configID, "enabled").Error; err != nil {
		return
	}
	// 限时判断
	if cfg.Limited != nil && *cfg.Limited {
		now := time.Now()
		if cfg.StartAt != nil && now.Before(*cfg.StartAt) {
			return ord, cfg, errors.New("当前不在活动开始时间")
		}
		if cfg.EndAt != nil && now.After(*cfg.EndAt) {
			return ord, cfg, errors.New("当前已过活动结束时间")
		}
	}
	// 首单优惠唯一性（用户维度）
	if cfg.Type != nil && *cfg.Type == "first_order" {
		var cnt int64
		if e := global.GVA_DB.WithContext(ctx).
			Table("app_points_recharge_orders o").
			Joins("JOIN app_points_configs c ON c.id = o.config_id").
			Where("o.user_id = ? AND c.type = ? AND (o.status = ? OR o.status = ?)", userID, "first_order", "success", "pending").
			Count(&cnt).Error; e != nil {
			return ord, cfg, e
		}
		if cnt > 0 {
			return ord, cfg, errors.New("首次充值已经完成，请选择其他充值类型")
		}
	}
	orderNo := fmt.Sprintf("R%v%04d", time.Now().Unix(), time.Now().Nanosecond()%10000)
	points := int64(0)
	bonus := int64(0)
	amount := 0.0
	if cfg.Points != nil {
		points = *cfg.Points
	}
	if cfg.BonusPoints != nil {
		bonus = *cfg.BonusPoints
	}
	if cfg.CurrentPrice != nil {
		amount = *cfg.CurrentPrice
	}
	// 确保Stripe客户
	var user model.AppUsers
	if e := global.GVA_DB.WithContext(ctx).First(&user, userID).Error; e == nil {
		if user.StripeCustomerID == nil || *user.StripeCustomerID == "" {
			custID, e2 := Stripe.EnsureCustomer(getStr(user.Email), user.ID)
			if e2 == nil {
				user.StripeCustomerID = &custID
				_ = global.GVA_DB.WithContext(ctx).Model(&model.AppUsers{}).Where("id = ?", user.ID).Update("stripe_customer_id", custID).Error
			}
		}
	}
	// 创建支付意图（金额单位：分）
	amountCents := int64(amount * 100)
	if amountCents < 0 {
		amountCents = 0
	}
	piID, clientSecret, err := Stripe.CreatePaymentIntent(amountCents, "CNY", payMethod, getStr(user.StripeCustomerID), map[string]string{"order_no": orderNo, "user_id": fmt.Sprintf("%d", userID)})

	if err != nil {
		return
	}
	ord = model.PointsRechargeOrder{
		OrderNo:            &orderNo,
		UserID:             &userID,
		ConfigID:           &configID,
		Points:             &points,
		BonusPoints:        &bonus,
		Amount:             &amount,
		PayMethod:          &payMethod,
		Status:             ptrStr("pending"),
		TransactionID:      &piID,
		StripeClientSecret: &clientSecret,
	}
	err = global.GVA_DB.WithContext(ctx).Create(&ord).Error
	return
}

// PayCallback 支付回调（成功逻辑幂等）
func (s *PR) PayCallback(ctx context.Context, orderNo string, paySuccess bool, transactionID string) (err error) {
	var ord model.PointsRechargeOrder
	if err = global.GVA_DB.WithContext(ctx).Where("order_no = ?", orderNo).First(&ord).Error; err != nil {
		return err
	}
	if ord.Status != nil && *ord.Status == "success" {
		return nil
	}
	return global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新订单状态
		now := time.Now()
		newStatus := "failed"
		if paySuccess {
			newStatus = "success"
		}
		if e := tx.Model(&model.PointsRechargeOrder{}).Where("id = ?", ord.ID).Updates(map[string]any{"status": newStatus, "transaction_id": transactionID, "paid_at": &now}).Error; e != nil {
			return e
		}
		if !paySuccess {
			return nil
		}
		// 更新账户余额
		uid := int64(0)
		if ord.UserID != nil {
			uid = *ord.UserID
		}
		acc, e := UserPointsAccount.EnsureAccount(ctx, uid)
		if e != nil {
			return e
		}
		add := int64(0)
		if ord.Points != nil {
			add += *ord.Points
		}
		if ord.BonusPoints != nil {
			add += *ord.BonusPoints
		}
		// 原子加积分
		res := tx.Model(&model.UserPointsAccount{}).Where("id = ?", acc.ID).UpdateColumn("balance", gorm.Expr("balance + ?", add))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return gorm.ErrInvalidData
		}
		// 获取新余额
		if e = tx.First(&acc, acc.ID).Error; e != nil {
			return e
		}
		after := int64(0)
		if acc.Balance != nil {
			after = *acc.Balance
		}
		// 写流水：充值购买
		relatedID := int64(ord.ID)
		_ = UserPointsAccount.AddLogDetailed(ctx, uid, add, after, "积分充值", orderNo, "", "recharge_purchase", "success", &relatedID)
		// 如有赠送，单独记录赠送
		if ord.BonusPoints != nil && *ord.BonusPoints > 0 {
			_ = UserPointsAccount.AddLogDetailed(ctx, uid, *ord.BonusPoints, after, "积分赠送", orderNo, "", "bonus", "success", &relatedID)
		}
		return nil
	})
}

func getStr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// GetRechargeOrdersByUser 用户维度查询充值订单列表
func (s *PR) GetRechargeOrdersByUser(ctx context.Context, userID int64, info request.PointsRechargeSearch) (list []model.PointsRechargeOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.WithContext(ctx).Model(&model.PointsRechargeOrder{}).Where("user_id = ?", userID)
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.PayMethod != nil && *info.PayMethod != "" {
		db = db.Where("pay_method = ?", *info.PayMethod)
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

// GetRechargeOrderByUser 用户维度查询单条充值订单
func (s *PR) GetRechargeOrderByUser(ctx context.Context, userID int64, orderNo string, id string) (ord model.PointsRechargeOrder, err error) {
	q := global.GVA_DB.WithContext(ctx).Model(&model.PointsRechargeOrder{}).Where("user_id = ?", userID)
	if orderNo != "" {
		err = q.Where("order_no = ?", orderNo).First(&ord).Error
		return
	}
	if id != "" {
		err = q.Where("id = ?", id).First(&ord).Error
		return
	}
	err = gorm.ErrRecordNotFound
	return
}
