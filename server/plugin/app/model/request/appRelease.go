
package request
import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)
type AppReleaseSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
       VersionName  *string `json:"versionName" form:"versionName"` 
       VersionCode  *int `json:"versionCode" form:"versionCode"` 
       Platform  *string `json:"platform" form:"platform"` 
       Status  *string `json:"status" form:"status"` 
       ForceUpdate  *bool `json:"forceUpdate" form:"forceUpdate"` 
       MinSupport  *int `json:"minSupport" form:"minSupport"` 
       PackageName  *string `json:"packageName" form:"packageName"` 
       BundleId  *string `json:"bundleId" form:"bundleId"` 
       AndroidPackage  string `json:"androidPackage" form:"androidPackage"` 
       IosPackage  string `json:"iosPackage" form:"iosPackage"` 
       AndroidUrl  *string `json:"androidUrl" form:"androidUrl"` 
       IosUrl  *string `json:"iosUrl" form:"iosUrl"` 
       Changelog  string `json:"changelog" form:"changelog"` 
       ReleaseAtRange  []time.Time  `json:"releaseAtRange" form:"releaseAtRange[]"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
