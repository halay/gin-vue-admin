
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var Node = new(NODE)

type NODE struct {}
// CreateNode 创建节点策略记录
// Author [yourname](https://github.com/yourname)
func (s *NODE) CreateNode(ctx context.Context, NODE *model.Node) (err error) {
	err = global.GVA_DB.Create(NODE).Error
	return err
}

// DeleteNode 删除节点策略记录
// Author [yourname](https://github.com/yourname)
func (s *NODE) DeleteNode(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Node{},"id = ?",ID).Error
	return err
}

// DeleteNodeByIds 批量删除节点策略记录
// Author [yourname](https://github.com/yourname)
func (s *NODE) DeleteNodeByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Node{},"id in ?",IDs).Error
	return err
}

// UpdateNode 更新节点策略记录
// Author [yourname](https://github.com/yourname)
func (s *NODE) UpdateNode(ctx context.Context, NODE model.Node) (err error) {
	err = global.GVA_DB.Model(&model.Node{}).Where("id = ?",NODE.ID).Updates(&NODE).Error
	return err
}

// GetNode 根据ID获取节点策略记录
// Author [yourname](https://github.com/yourname)
func (s *NODE) GetNode(ctx context.Context, ID string) (NODE model.Node, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&NODE).Error
	return
}
// GetNodeInfoList 分页获取节点策略记录
// Author [yourname](https://github.com/yourname)
func (s *NODE) GetNodeInfoList(ctx context.Context, info request.NodeSearch) (list []model.Node, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Node{})
    var NODEs []model.Node
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.SystemLevel != nil && *info.SystemLevel != "" {
        db = db.Where("system_level = ?", *info.SystemLevel)
    }
    if info.LimitCount != nil {
        db = db.Where("limit_count >= ?", *info.LimitCount)
    }
    if info.JoinFee != nil {
        db = db.Where("join_fee >= ?", *info.JoinFee)
    }
    if info.SettlementCurrency != nil && *info.SettlementCurrency != "" {
        db = db.Where("settlement_currency = ?", *info.SettlementCurrency)
    }
    if info.CitySurchargePercent != nil {
        db = db.Where("city_surcharge_percent >= ?", *info.CitySurchargePercent)
    }
    if info.TokenDistributionX != nil {
        db = db.Where("token_distribution_x >= ?", *info.TokenDistributionX)
    }
    if info.ReleaseWeeks != nil {
        db = db.Where("release_weeks >= ?", *info.ReleaseWeeks)
    }
    if info.CashCouponPercent != nil {
        db = db.Where("cash_coupon_percent >= ?", *info.CashCouponPercent)
    }
    if info.PVBonusPercent != nil {
        db = db.Where("pv_bonus_percent >= ?", *info.PVBonusPercent)
    }
    if info.RwaBrandFeeSharePercent != nil {
        db = db.Where("rwa_brand_fee_share_percent >= ?", *info.RwaBrandFeeSharePercent)
    }
    if info.AnnualFeeSharePercent != nil {
        db = db.Where("annual_fee_share_percent >= ?", *info.AnnualFeeSharePercent)
    }
    if info.OnlineTradeFeeSharePercent != nil {
        db = db.Where("online_trade_fee_share_percent >= ?", *info.OnlineTradeFeeSharePercent)
    }
    if info.EquityPercent != nil {
        db = db.Where("equity_percent >= ?", *info.EquityPercent)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
    if info.SortValue != nil {
        db = db.Where("sort = ?", *info.SortValue)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["name"] = true
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
	err = db.Find(&NODEs).Error
	return  NODEs, total, err
}

func (s *NODE)GetNodePublic(ctx context.Context) {

}
