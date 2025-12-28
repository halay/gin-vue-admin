
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var AgentTransactionDetail = new(atd)

type atd struct {}
// CreateAgentTransactionDetail 创建代理交易明细记录
// Author [yourname](https://github.com/yourname)
func (s *atd) CreateAgentTransactionDetail(ctx context.Context, atd *model.AgentTransactionDetail) (err error) {
	err = global.GVA_DB.Create(atd).Error
	return err
}

// DeleteAgentTransactionDetail 删除代理交易明细记录
// Author [yourname](https://github.com/yourname)
func (s *atd) DeleteAgentTransactionDetail(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.AgentTransactionDetail{},"id = ?",ID).Error
	return err
}

// DeleteAgentTransactionDetailByIds 批量删除代理交易明细记录
// Author [yourname](https://github.com/yourname)
func (s *atd) DeleteAgentTransactionDetailByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.AgentTransactionDetail{},"id in ?",IDs).Error
	return err
}

// UpdateAgentTransactionDetail 更新代理交易明细记录
// Author [yourname](https://github.com/yourname)
func (s *atd) UpdateAgentTransactionDetail(ctx context.Context, atd model.AgentTransactionDetail) (err error) {
	err = global.GVA_DB.Model(&model.AgentTransactionDetail{}).Where("id = ?",atd.ID).Updates(&atd).Error
	return err
}

// GetAgentTransactionDetail 根据ID获取代理交易明细记录
// Author [yourname](https://github.com/yourname)
func (s *atd) GetAgentTransactionDetail(ctx context.Context, ID string) (atd model.AgentTransactionDetail, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&atd).Error
	return
}
// GetAgentTransactionDetailInfoList 分页获取代理交易明细记录
// Author [yourname](https://github.com/yourname)
func (s *atd) GetAgentTransactionDetailInfoList(ctx context.Context, info request.AgentTransactionDetailSearch) (list []model.AgentTransactionDetail, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AgentTransactionDetail{})
    var atds []model.AgentTransactionDetail
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
        orderMap["sequence_number"] = true
        orderMap["transaction_time"] = true
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
	err = db.Find(&atds).Error
	return  atds, total, err
}

func (s *atd)GetAgentTransactionDetailPublic(ctx context.Context) {

}
