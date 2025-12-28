
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var Settlement = new(settlement)

type settlement struct {}
// CreateSettlement 创建结算管理记录
// Author [yourname](https://github.com/yourname)
func (s *settlement) CreateSettlement(ctx context.Context, settlement *model.Settlement) (err error) {
	err = global.GVA_DB.Create(settlement).Error
	return err
}

// DeleteSettlement 删除结算管理记录
// Author [yourname](https://github.com/yourname)
func (s *settlement) DeleteSettlement(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Settlement{},"id = ?",ID).Error
	return err
}

// DeleteSettlementByIds 批量删除结算管理记录
// Author [yourname](https://github.com/yourname)
func (s *settlement) DeleteSettlementByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Settlement{},"id in ?",IDs).Error
	return err
}

// UpdateSettlement 更新结算管理记录
// Author [yourname](https://github.com/yourname)
func (s *settlement) UpdateSettlement(ctx context.Context, settlement model.Settlement) (err error) {
	err = global.GVA_DB.Model(&model.Settlement{}).Where("id = ?",settlement.ID).Updates(&settlement).Error
	return err
}

// GetSettlement 根据ID获取结算管理记录
// Author [yourname](https://github.com/yourname)
func (s *settlement) GetSettlement(ctx context.Context, ID string) (settlement model.Settlement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&settlement).Error
	return
}
// GetSettlementInfoList 分页获取结算管理记录
// Author [yourname](https://github.com/yourname)
func (s *settlement) GetSettlementInfoList(ctx context.Context, info request.SettlementSearch) (list []model.Settlement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Settlement{})
    var settlements []model.Settlement
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["serial_number"] = true
        orderMap["name"] = true
        orderMap["settlement_period"] = true
        orderMap["subsidy_amount"] = true
        orderMap["status"] = true
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
	err = db.Find(&settlements).Error
	return  settlements, total, err
}

func (s *settlement)GetSettlementPublic(ctx context.Context) {

}
