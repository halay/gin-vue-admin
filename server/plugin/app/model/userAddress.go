
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserAddress 用户收货地址 结构体
type UserAddress struct {
    global.GVA_MODEL
  UserID  *int64 `json:"userId" form:"userId" gorm:"comment:app_users.ID;column:user_id;" binding:"required"`  //用户ID
  Receiver  *string `json:"receiver" form:"receiver" gorm:"comment:姓名;column:receiver;size:64;" binding:"required"`  //收货人姓名
  Phone  *string `json:"phone" form:"phone" gorm:"comment:手机;column:phone;size:32;" binding:"required"`  //联系电话
  Country  *string `json:"country" form:"country" gorm:"comment:国家;column:country;size:64;"`  //国家
  Province  *string `json:"province" form:"province" gorm:"comment:省;column:province;size:64;"`  //省
  City  *string `json:"city" form:"city" gorm:"comment:市;column:city;size:64;"`  //市
  District  *string `json:"district" form:"district" gorm:"comment:区县;column:district;size:64;"`  //区/县
  Address  *string `json:"address" form:"address" gorm:"comment:详细地址;column:address;size:255;" binding:"required"`  //详细地址
  Zipcode  *string `json:"zipcode" form:"zipcode" gorm:"comment:邮编;column:zipcode;size:16;"`  //邮编
  IsDefault  *bool `json:"isDefault" form:"isDefault" gorm:"comment:默认地址;column:is_default;"`  //是否默认
  Status  *string `json:"status" form:"status" gorm:"default:enabled;comment:启用/禁用;column:status;"`  //状态
  Remark  *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:255;"`  //备注
}


// TableName 用户收货地址 UserAddress自定义表名 app_user_addresses
func (UserAddress) TableName() string {
    return "app_user_addresses"
}







