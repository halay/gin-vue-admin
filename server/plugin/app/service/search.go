package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
	"time"
)

var Search = new(SRCH)

type SRCH struct{}

func (s *SRCH) GetSearchPublic(ctx context.Context, info request.SearchPublic) (list []model.Product, total int64, guess []model.Product, hot []model.Product, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.WithContext(ctx).Model(&model.Product{})
	if info.MerchantID != nil {
		db = db.Where("merchant_id = ?", *info.MerchantID)
	}
	if info.Keyword != "" {
		kw := "%" + info.Keyword + "%"
		db = db.Where("name LIKE ? OR detail LIKE ?", kw, kw)
	}
	if info.CategoryID != nil {
		db = db.Where("category_id = ?", *info.CategoryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&list).Error
	if err != nil {
		return
	}
	hotIds := make([]int64, 0)
	type Row struct{ ProductID int64 }
	rows := make([]Row, 0)
	// 选择排序指标
	metric := info.Metric
	if metric != "amount" {
		metric = "quantity"
	}
	sel := "product_id as product_id"
	if metric == "amount" {
		sel = sel
	}
	q := global.GVA_DB.WithContext(ctx).Table("app_order_items").Select(sel)
	if info.MerchantID != nil {
		q = q.Where("merchant_id = ?", *info.MerchantID)
	}
	if info.Days > 0 {
		since := time.Now().Add(-time.Duration(info.Days) * 24 * time.Hour)
		q = q.Where("created_at >= ?", since)
	}
	orderExpr := "SUM(quantity) desc"
	if metric == "amount" {
		orderExpr = "SUM(total_amount) desc"
	}
	q = q.Group("product_id").Order(orderExpr).Limit(10)
	_ = q.Scan(&rows).Error
	for _, r := range rows {
		hotIds = append(hotIds, r.ProductID)
	}
	if len(hotIds) > 0 {
		_ = global.GVA_DB.WithContext(ctx).Model(&model.Product{}).Where("id in ?", hotIds).Find(&hot).Error
	}
	gdb := global.GVA_DB.WithContext(ctx).Model(&model.Product{})
	if info.MerchantID != nil {
		gdb = gdb.Where("merchant_id = ?", *info.MerchantID)
	}
	_ = gdb.Order("created_at desc").Limit(10).Find(&guess).Error
	return
}
