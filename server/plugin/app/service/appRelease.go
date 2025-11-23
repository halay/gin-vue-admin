
package service

import (
    "context"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "gorm.io/gorm"
    "encoding/json"
    "gorm.io/datatypes"
)

var AppRelease = new(AR)

type AR struct {}
// CreateAppRelease 创建App客户端版本升级记录
// Author [yourname](https://github.com/yourname)
func (s *AR) CreateAppRelease(ctx context.Context, AR *model.AppRelease) (err error) {
    if u := fileURLFromJSON(AR.AndroidPackage); u != "" { AR.AndroidUrl = toStrPtr(u) }
    if u := fileURLFromJSON(AR.IosPackage); u != "" { AR.IosUrl = toStrPtr(u) }
    err = global.GVA_DB.Create(AR).Error
    return err
}

// DeleteAppRelease 删除App客户端版本升级记录
// Author [yourname](https://github.com/yourname)
func (s *AR) DeleteAppRelease(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.AppRelease{},"id = ?",ID).Error
	return err
}

// DeleteAppReleaseByIds 批量删除App客户端版本升级记录
// Author [yourname](https://github.com/yourname)
func (s *AR) DeleteAppReleaseByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.AppRelease{},"id in ?",IDs).Error
	return err
}

// UpdateAppRelease 更新App客户端版本升级记录
// Author [yourname](https://github.com/yourname)
func (s *AR) UpdateAppRelease(ctx context.Context, AR model.AppRelease) (err error) {
    if u := fileURLFromJSON(AR.AndroidPackage); u != "" { AR.AndroidUrl = toStrPtr(u) }
    if u := fileURLFromJSON(AR.IosPackage); u != "" { AR.IosUrl = toStrPtr(u) }
    err = global.GVA_DB.Model(&model.AppRelease{}).Where("id = ?", AR.ID).Updates(&AR).Error
    return err
}

// GetAppRelease 根据ID获取App客户端版本升级记录
// Author [yourname](https://github.com/yourname)
func (s *AR) GetAppRelease(ctx context.Context, ID string) (AR model.AppRelease, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&AR).Error
	return
}
// GetAppReleaseInfoList 分页获取App客户端版本升级记录
// Author [yourname](https://github.com/yourname)
func (s *AR) GetAppReleaseInfoList(ctx context.Context, info request.AppReleaseSearch) (list []model.AppRelease, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AppRelease{})
    var ARs []model.AppRelease
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.VersionName != nil && *info.VersionName != "" {
        db = db.Where("version_name LIKE ?", "%"+ *info.VersionName+"%")
    }
    if info.VersionCode != nil {
        db = db.Where("version_code >= ?", *info.VersionCode)
    }
    if info.Platform != nil && *info.Platform != "" {
        db = db.Where("platform = ?", *info.Platform)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
    if info.ForceUpdate != nil {
        db = db.Where("force_update = ?", *info.ForceUpdate)
    }
    if info.MinSupport != nil {
        db = db.Where("min_support >= ?", *info.MinSupport)
    }
    if info.PackageName != nil && *info.PackageName != "" {
        db = db.Where("package_name LIKE ?", "%"+ *info.PackageName+"%")
    }
    if info.BundleId != nil && *info.BundleId != "" {
        db = db.Where("bundle_id LIKE ?", "%"+ *info.BundleId+"%")
    }
    if info.AndroidPackage != "" {
        db = db.Where("android_package LIKE ?", "%"+ info.AndroidPackage+"%")
    }
    if info.IosPackage != "" {
        db = db.Where("ios_package LIKE ?", "%"+ info.IosPackage+"%")
    }
    if info.AndroidUrl != nil && *info.AndroidUrl != "" {
        db = db.Where("android_url LIKE ?", "%"+ *info.AndroidUrl+"%")
    }
    if info.IosUrl != nil && *info.IosUrl != "" {
        db = db.Where("ios_url LIKE ?", "%"+ *info.IosUrl+"%")
    }
    if info.Changelog != "" {
        // TODO 数据类型为复杂类型，请根据业务需求自行实现复杂类型的查询业务
    }
			if len(info.ReleaseAtRange) == 2 {
				db = db.Where("release_at BETWEEN ? AND ? ", info.ReleaseAtRange[0], info.ReleaseAtRange[1])
			}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
        orderMap["id"] = true
        orderMap["created_at"] = true
        orderMap["version_name"] = true
        orderMap["version_code"] = true
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
	err = db.Find(&ARs).Error
	return  ARs, total, err
}

func (s *AR)GetAppReleasePublic(ctx context.Context) {

}

// CheckUpdate 返回是否需要更新及下载信息
func (s *AR) CheckUpdate(ctx context.Context, platform string, versionCode int64) (need bool, release model.AppRelease, err error) {
    var ar model.AppRelease
    err = global.GVA_DB.WithContext(ctx).Where("platform = ? AND status = ?", platform, "enabled").Order("version_code DESC").First(&ar).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound { return false, ar, nil }
        return false, ar, err
    }
    cur := int64(0)
    if ar.VersionCode != nil { cur = *ar.VersionCode }
    need = cur > versionCode
    release = ar
    return
}

func fileURLFromJSON(j datatypes.JSON) string {
    if len(j) == 0 { return "" }
    var v any
    if err := json.Unmarshal(j, &v); err != nil { return "" }
    switch vv := v.(type) {
    case map[string]any:
        if s, ok := pickURL(vv); ok { return s }
    case []any:
        for _, it := range vv {
            if m, ok := it.(map[string]any); ok {
                if s, ok := pickURL(m); ok { return s }
            }
        }
    }
    return ""
}

func pickURL(m map[string]any) (string, bool) {
    keys := []string{"url", "path", "fileUrl"}
    for _, k := range keys {
        if v, ok := m[k]; ok {
            if s, ok2 := v.(string); ok2 && s != "" { return s, true }
        }
    }
    return "", false
}

func toStrPtr(s string) *string { return &s }
