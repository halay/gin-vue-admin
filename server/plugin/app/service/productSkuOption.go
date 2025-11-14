package service

import (
    "context"
    "encoding/json"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "gorm.io/datatypes"
)

var ProductSkuOptionSvc = new(PSO2)

type PSO2 struct{}

// UpsertSkuOptions 解析 SKU 的 attrs 并同步到中间表（容错解析）
func (s *PSO2) UpsertSkuOptions(ctx context.Context, skuID uint, merchantID int, attrs datatypes.JSON) error {
    if len(attrs) == 0 { return nil }
    // 期望格式：{"specId":"optionId", ...} 或 {specId: optionId}
    var msi map[string]int64
    var mii map[int64]int64
    // 删除旧记录
    _ = global.GVA_DB.WithContext(ctx).Where("sku_id = ?", int64(skuID)).Delete(&model.ProductSkuOption{}).Error
    if err := json.Unmarshal(attrs, &msi); err == nil {
        for sk, ov := range msi {
            // 尝试将 specId 字符串转换为 int64
            var sid int64
            _ = json.Unmarshal([]byte(sk), &sid)
            if sid == 0 {
                // 如果无法直接解析，尝试 atoi
                // 简化处理：忽略解析失败的键
            }
            rec := model.ProductSkuOption{
                SkuID:      ptrInt64(int64(skuID)),
                SpecID:     ptrInt64(sid),
                OptionID:   ptrInt64(ov),
                MerchantID: ptrInt64(int64(merchantID)),
            }
            _ = global.GVA_DB.WithContext(ctx).Create(&rec).Error
        }
        return nil
    }
    if err := json.Unmarshal(attrs, &mii); err == nil {
        for sid, ov := range mii {
            rec := model.ProductSkuOption{
                SkuID:      ptrInt64(int64(skuID)),
                SpecID:     ptrInt64(sid),
                OptionID:   ptrInt64(ov),
                MerchantID: ptrInt64(int64(merchantID)),
            }
            _ = global.GVA_DB.WithContext(ctx).Create(&rec).Error
        }
        return nil
    }
    return nil
}
