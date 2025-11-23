
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// AppRelease App客户端版本升级 结构体
type AppRelease struct {
    global.GVA_MODEL
  VersionName  *string `json:"versionName" form:"versionName" gorm:"comment:版本名;column:version_name;size:64;" binding:"required"`  //版本名
  VersionCode  *int64 `json:"versionCode" form:"versionCode" gorm:"comment:版本号（整型递增）;column:version_code;" binding:"required"`  //版本号
  Platform  *string `json:"platform" form:"platform" gorm:"default:android;comment:平台：android/ios;column:platform;" binding:"required"`  //平台
  Status  *string `json:"status" form:"status" gorm:"default:enabled;comment:发布状态;column:status;" binding:"required"`  //状态
  ForceUpdate  *bool `json:"forceUpdate" form:"forceUpdate" gorm:"comment:是否强制更新;column:force_update;"`  //强制更新
  MinSupport  *int64 `json:"minSupport" form:"minSupport" gorm:"comment:低于该版本必须强更;column:min_support;"`  //最低支持版本号
  PackageName  *string `json:"packageName" form:"packageName" gorm:"comment:Android 包名;column:package_name;size:128;"`  //安卓包名
  BundleId  *string `json:"bundleId" form:"bundleId" gorm:"comment:iOS Bundle ID;column:bundle_id;size:128;"`  //iOS Bundle ID
  AndroidPackage  datatypes.JSON `json:"androidPackage" form:"androidPackage" gorm:"comment:安卓安装包文件;column:android_package;" swaggertype:"array,object"`  //安卓安装包
  IosPackage  datatypes.JSON `json:"iosPackage" form:"iosPackage" gorm:"comment:iOS 安装包或跳转链接文件;column:ios_package;" swaggertype:"array,object"`  //iOS安装包/链接
  AndroidUrl  *string `json:"androidUrl" form:"androidUrl" gorm:"comment:安卓下载URL;column:android_url;size:255;"`  //安卓下载链接
  IosUrl  *string `json:"iosUrl" form:"iosUrl" gorm:"comment:iOS 下载URL;column:ios_url;size:255;"`  //iOS下载链接
  Changelog  *string `json:"changelog" form:"changelog" gorm:"comment:更新日志;column:changelog;type:text;"`  //更新日志
  ReleaseAt  *time.Time `json:"releaseAt" form:"releaseAt" gorm:"comment:发布时间;column:release_at;"`  //发布时间
}


// TableName App客户端版本升级 AppRelease自定义表名 app_releases
func (AppRelease) TableName() string {
    return "app_releases"
}







