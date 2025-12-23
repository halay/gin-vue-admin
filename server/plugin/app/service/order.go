package service

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"gorm.io/gorm"
)

var Order = new(ORD)

type ORD struct{}

// CreateOrder 创建订单记录
// Author [yourname](https://github.com/yourname)
func (s *ORD) CreateOrder(ctx context.Context, ORD *model.Order, merchantID int) (err error) {
	ORD.MerchantID = ptrInt64(int64(merchantID))
	err = global.GVA_DB.WithContext(ctx).Create(ORD).Error
	return err
}

// DeleteOrder 删除订单记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *ORD) DeleteOrder(ctx context.Context, ID string, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&model.Order{}, "id = ?", ID).Error
	return err
}

// DeleteOrderByIds 批量删除订单记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *ORD) DeleteOrderByIds(ctx context.Context, IDs []string, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Delete(&[]model.Order{}, "id in ?", IDs).Error
	return err
}

// UpdateOrder 更新订单记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *ORD) UpdateOrder(ctx context.Context, ORD model.Order, merchantID int) (err error) {
	err = global.GVA_DB.WithContext(ctx).Model(&model.Order{}).Where("id = ? AND merchant_id = ?", ORD.ID, merchantID).Updates(&ORD).Error
	return err
}

// GetOrder 根据ID获取订单记录（限定商户范围）
// Author [yourname](https://github.com/yourname)
func (s *ORD) GetOrder(ctx context.Context, ID string, merchantID int) (ORD model.Order, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ? AND merchant_id = ?", ID, merchantID).First(&ORD).Error
	return
}

// GetOrderInfoList 分页获取订单记录
// Author [yourname](https://github.com/yourname)
func (s *ORD) GetOrderInfoList(ctx context.Context, info request.OrderSearch, merchantID int) (list []model.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.WithContext(ctx).Model(&model.Order{})
	var ORDs []model.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.OrderNo != nil && *info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+*info.OrderNo+"%")
	}
	if info.UserID != nil {
		db = db.Where("user_id = ?", *info.UserID)
	}
	db = db.Where("merchant_id = ?", merchantID)
	if info.TotalAmount != nil {
		db = db.Where("total_amount >= ?", *info.TotalAmount)
	}
	if info.TotalPoints != nil {
		db = db.Where("total_points >= ?", *info.TotalPoints)
	}
	if info.PayMethod != nil && *info.PayMethod != "" {
		db = db.Where("pay_method = ?", *info.PayMethod)
	}
	if info.PayStatus != nil && *info.PayStatus != "" {
		db = db.Where("pay_status = ?", *info.PayStatus)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.ConsigneeName != nil && *info.ConsigneeName != "" {
		db = db.Where("consignee_name LIKE ?", "%"+*info.ConsigneeName+"%")
	}
	if info.ConsigneePhone != nil && *info.ConsigneePhone != "" {
		db = db.Where("consignee_phone LIKE ?", "%"+*info.ConsigneePhone+"%")
	}
	if info.Address != nil && *info.Address != "" {
		db = db.Where("address LIKE ?", "%"+*info.Address+"%")
	}
	if info.Country != nil && *info.Country != "" {
		db = db.Where("country LIKE ?", "%"+*info.Country+"%")
	}
	if info.PostalCode != nil && *info.PostalCode != "" {
		db = db.Where("postal_code LIKE ?", "%"+*info.PostalCode+"%")
	}
	if info.DeliveryStatus != nil && *info.DeliveryStatus != "" {
		db = db.Where("delivery_status = ?", *info.DeliveryStatus)
	}
	if info.ExpressName != nil && *info.ExpressName != "" {
		db = db.Where("express_name LIKE ?", "%"+*info.ExpressName+"%")
	}
	if info.ExpressNo != nil && *info.ExpressNo != "" {
		db = db.Where("express_no LIKE ?", "%"+*info.ExpressNo+"%")
	}
	if info.Remark != nil && *info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+*info.Remark+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["created_at"] = true
	orderMap["order_no"] = true
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
	err = db.Find(&ORDs).Error
	return ORDs, total, err
}
func (s *ORD) GetOrderDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	userId := make([]map[string]any, 0)
	global.GVA_DB.Table("app_users").Where("deleted_at IS NULL").Select("email as label,id as value").Scan(&userId)
	res["userId"] = userId
	return
}

func (s *ORD) GetOrderPublic(ctx context.Context) {

}

// GetOrderInfoListByUser 用户维度查询订单列表（跨商户）
func (s *ORD) GetOrderInfoListByUser(ctx context.Context, userID int64, info request.OrderSearch) (list []model.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.WithContext(ctx).Model(&model.Order{}).Where("user_id = ?", userID)
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}
	if info.OrderNo != nil && *info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+*info.OrderNo+"%")
	}
	if info.PayStatus != nil && *info.PayStatus != "" {
		db = db.Where("pay_status = ?", *info.PayStatus)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.Country != nil && *info.Country != "" {
		db = db.Where("country LIKE ?", "%"+*info.Country+"%")
	}
	if info.PostalCode != nil && *info.PostalCode != "" {
		db = db.Where("postal_code LIKE ?", "%"+*info.PostalCode+"%")
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

// GetOrderDetailByUser 用户维度查询订单详情与明细
func (s *ORD) GetOrderDetailByUser(ctx context.Context, userID int64, orderNo string) (ord model.Order, items []model.OrderItem, err error) {
	if err = global.GVA_DB.WithContext(ctx).Where("order_no = ? AND user_id = ?", orderNo, userID).First(&ord).Error; err != nil {
		return
	}
	id64 := int64(ord.ID)
	err = global.GVA_DB.WithContext(ctx).Where("order_id = ?", id64).Find(&items).Error
	return
}

// CreateOrderByPoints 生成订单并创建明细（仅积分支付），返回订单与明细
func (s *ORD) CreateOrderByPoints(ctx context.Context, userID int64, sku model.ProductSku, qty int64, consigneeName, consigneePhone, address string, province, city, district string, country string, postalCode string, payMethod string) (ord model.Order, item model.OrderItem, err error) {
	orderNo := fmt.Sprintf("P%v%04d", time.Now().Unix(), time.Now().Nanosecond()%10000)
	// 计算金额或积分
	totalPoints := int64(0)
	totalAmount := 0.0
	if payMethod == "card" && sku.Price != nil {
		totalAmount = *sku.Price * float64(qty)
	} else {
		if sku.Points != nil {
			totalPoints = *sku.Points * qty
		}
	}
	merchantID := int64(0)
	if sku.MerchantID != nil {
		merchantID = *sku.MerchantID
	}

	// 读取商品名称与封面快照
	var prod model.Product
	_ = global.GVA_DB.WithContext(ctx).First(&prod, sku.ProductID).Error
	var prodName string
	if prod.Name != nil {
		prodName = *prod.Name
	}
	cover := sku.Image
	if cover == "" && prod.CoverImage != "" {
		cover = prod.CoverImage
	}

	ord = model.Order{
		OrderNo:        &orderNo,
		UserID:         &userID,
		MerchantID:     &merchantID,
		TotalAmount:    ptrFloat64(totalAmount),
		TotalPoints:    &totalPoints,
		PayMethod:      ptrStr(payMethod),
		PayStatus:      ptrStr("unpaid"),
		Status:         ptrStr("created"),
		ConsigneeName:  nilIfEmpty(consigneeName),
		ConsigneePhone: nilIfEmpty(consigneePhone),
		Address:        nilIfEmpty(address),
		Country:        nilIfEmpty(country),
		PostalCode:     nilIfEmpty(postalCode),
		Province:       nilIfEmpty(province),
		City:           nilIfEmpty(city),
		District:       nilIfEmpty(district),
		DeliveryStatus: ptrStr("none"),
	}
	item = model.OrderItem{
		ProductID:   sku.ProductID,
		SkuID:       uToI64Ptr(sku.ID),
		ProductName: ptrStr(prodName),
		SkuAttrs:    sku.Attrs,
		Price: func() *float64 {
			if payMethod == "card" && sku.Price != nil {
				return ptrFloat64(*sku.Price)
			}
			return ptrFloat64(0)
		}(),
		Points: func() *int64 {
			if payMethod != "card" {
				return sku.Points
			}
			return nil
		}(),
		Quantity:    &qty,
		TotalAmount: ptrFloat64(totalAmount),
		TotalPoints: &totalPoints,
		CoverImage:  cover,
		MerchantID:  sku.MerchantID,
	}

	err = global.GVA_DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if e := tx.Create(&ord).Error; e != nil {
			return e
		}
		// 绑定明细到订单
		id64 := int64(ord.ID)
		item.OrderID = &id64
		if e := tx.Create(&item).Error; e != nil {
			return e
		}
		return nil
	})
	return
}

// PayOrderByPoints 积分支付订单（仅未支付->已支付），原子扣减积分并写流水
func (s *ORD) PayOrderByPoints(ctx context.Context, userID int64, orderNo string) error {
	var ord model.Order
	if err := global.GVA_DB.WithContext(ctx).Where("order_no = ? AND user_id = ?", orderNo, userID).First(&ord).Error; err != nil {
		return err
	}
	if ord.PayStatus != nil && *ord.PayStatus == "paid" {
		return nil
	}
	// 扣减积分
	need := int64(0)
	if ord.TotalPoints != nil {
		need = *ord.TotalPoints
	}
	after, err := UserPointsAccount.DeductPoints(ctx, userID, need)
	if err != nil {
		return err
	}
	// 写流水
	_ = UserPointsAccount.AddLog(ctx, userID, -need, after, "订单积分支付", orderNo, "")
	// 更新订单状态
	return global.GVA_DB.WithContext(ctx).Model(&model.Order{}).
		Where("id = ?", ord.ID).
		Updates(map[string]any{"pay_status": "paid", "status": "paid"}).Error
}

// CreateOrderPaymentIntent 创建Stripe支付意图（银行卡支付）
func (s *ORD) CreateOrderPaymentIntent(ctx context.Context, userID uint, orderNo string) (piID string, clientSecret string, err error) {
	// 查订单（归属校验）
	var ord model.Order
	if err = global.GVA_DB.WithContext(ctx).Where("order_no = ? AND user_id = ?", orderNo, int64(userID)).First(&ord).Error; err != nil {
		return
	}
	// 金额（分）
	amount := 0.0
	if ord.TotalAmount != nil {
		amount = *ord.TotalAmount
	}
	amountCents := int64(amount * 100)
	if amountCents <= 0 {
		err = fmt.Errorf("订单金额无效")
		return
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
	// 创建支付意图，标记为商品订单
	piID, clientSecret, err = Stripe.CreatePaymentIntent(amountCents, "cny", "card", getStr(user.StripeCustomerID), map[string]string{
		"order_no": orderNo,
		"user_id":  fmt.Sprintf("%d", userID),
		"pay_type": "order",
	})
	// 更新订单支付方式为card（可选）
	if err == nil {
		_ = global.GVA_DB.WithContext(ctx).Model(&model.Order{}).Where("id = ?", ord.ID).Update("pay_method", "card").Error
	}
	return
}

// PayCallbackCard 商品订单支付回调
func (s *ORD) PayCallbackCard(ctx context.Context, orderNo string, paySuccess bool, transactionID string) error {
	var ord model.Order
	if err := global.GVA_DB.WithContext(ctx).Where("order_no = ?", orderNo).First(&ord).Error; err != nil {
		return err
	}
	if ord.PayStatus != nil && *ord.PayStatus == "paid" {
		return nil
	}
	newStatus := "failed"
	if paySuccess {
		newStatus = "paid"
	}
	return global.GVA_DB.WithContext(ctx).Model(&model.Order{}).
		Where("id = ?", ord.ID).
		Updates(map[string]any{"pay_status": newStatus, "status": newStatus}).Error
}

func ptrStr(s string) *string       { return &s }
func ptrFloat64(f float64) *float64 { return &f }
func uToI64Ptr(u uint) *int64       { v := int64(u); return &v }
func nilIfEmpty(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}
