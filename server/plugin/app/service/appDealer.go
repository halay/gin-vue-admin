
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

var AppDealer = new(ADL)

type ADL struct {}
// CreateAppDealer 创建经销商管理记录
// Author [yourname](https://github.com/yourname)
func (s *ADL) CreateAppDealer(ctx context.Context, ADL *model.AppDealer) (err error) {
	err = global.GVA_DB.Create(ADL).Error
	return err
}

// DeleteAppDealer 删除经销商管理记录
// Author [yourname](https://github.com/yourname)
func (s *ADL) DeleteAppDealer(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.AppDealer{},"id = ?",ID).Error
	return err
}

// DeleteAppDealerByIds 批量删除经销商管理记录
// Author [yourname](https://github.com/yourname)
func (s *ADL) DeleteAppDealerByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.AppDealer{},"id in ?",IDs).Error
	return err
}

// UpdateAppDealer 更新经销商管理记录
// Author [yourname](https://github.com/yourname)
func (s *ADL) UpdateAppDealer(ctx context.Context, ADL model.AppDealer) (err error) {
	err = global.GVA_DB.Model(&model.AppDealer{}).Where("id = ?",ADL.ID).Updates(&ADL).Error
	return err
}

// GetAppDealer 根据ID获取经销商管理记录
// Author [yourname](https://github.com/yourname)
func (s *ADL) GetAppDealer(ctx context.Context, ID string) (ADL model.AppDealer, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ADL).Error
	return
}
// GetAppDealerInfoList 分页获取经销商管理记录
// Author [yourname](https://github.com/yourname)
func (s *ADL) GetAppDealerInfoList(ctx context.Context, info request.AppDealerSearch) (list []model.AppDealer, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AppDealer{})
    var ADLs []model.AppDealer
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.SalesCommission != nil {
        db = db.Where("sales_commission = ?", *info.SalesCommission)
    }
    if info.CommissionType != nil {
        db = db.Where("commission_type = ?", *info.CommissionType)
    }
    if info.ExpenseAllowance != nil {
        db = db.Where("expense_allowance = ?", *info.ExpenseAllowance)
    }
    if info.AllowanceType != nil {
        db = db.Where("allowance_type = ?", *info.AllowanceType)
    }
    if info.MerchantId != nil {
        db = db.Where("merchant_id = ?", *info.MerchantId)
    }
    if info.Remark != nil && *info.Remark != "" {
        db = db.Where("remark LIKE ?", "%"+ *info.Remark+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&ADLs).Error
	return  ADLs, total, err
}
func (s *ADL)GetAppDealerDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   merchantId := make([]map[string]any, 0)
	   global.GVA_DB.Table("app_merchants").Select("merchant_name as label,id as value").Scan(&merchantId)
	   res["merchantId"] = merchantId
	return
}

func (s *ADL)GetAppDealerPublic(ctx context.Context) {

}
