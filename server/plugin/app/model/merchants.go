package model

import (
	"gorm.io/datatypes"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Merchants 商家信息 结构体
type Merchants struct {
	global.GVA_MODEL
	MerchantName    *string        `json:"merchantName" form:"merchantName" gorm:"comment:商家名称;column:merchant_name;" binding:"required"`        //商家名
	Logo            string         `json:"logo" form:"logo" gorm:"comment:商家LOGO图片;column:logo;" binding:"required"`                             //商户logo
	BusinessScope   *string        `json:"businessScope" form:"businessScope" gorm:"comment:商家经营范围;column:business_scope;"`                      //经营范围描述
	Address         *string        `json:"address" form:"address" gorm:"comment:商家地址;column:address;"`                                           //地址
	ContactPerson   *string        `json:"contactPerson" form:"contactPerson" gorm:"comment:联系人姓名;column:contact_person;"`                       //联系人
	ContactPhone    *string        `json:"contactPhone" form:"contactPhone" gorm:"comment:联系电话号码;column:contact_phone;"`                         //联系电话
	Rating          *float64       `json:"rating" form:"rating" gorm:"comment:商家评分;column:rating;"`                                              //评分
	Email           *string        `json:"email" form:"email" gorm:"column:email;size:50;"`                                                      //邮箱
	BusinessLicense string         `json:"businessLicense" form:"businessLicense" gorm:"comment:营业执照;column:business_license;"`                  //营业执照
	IdCardImage     datatypes.JSON `json:"idCardImage" form:"idCardImage" gorm:"comment:身份证照片;column:id_card_image;" swaggertype:"array,object"` //身份证照片
	Status          *string        `json:"status" form:"status" gorm:"index;default:1;comment:状态: 1-正常, 2-禁用;column:status;"`                    //状态
	// 在结构体中新增如下字段
	Founder             *string        `json:"founder" form:"founder" gorm:"comment:创始人/核心人员;column:founder;"`                                             //创始人
	FounderPhone        *string        `json:"founderPhone" form:"founderPhone" gorm:"comment:创始人电话;column:founder_phone;"`                                //创始人电话
	FounderDescriptions *string        `json:"founderDescriptions" form:"founderDescriptions" gorm:"comment:创始人简介;column:founder_descriptions;type:text;"` //创始人简介
	EnterpriseDesc      *string        `json:"enterpriseDesc" form:"enterpriseDesc" gorm:"column:enterprise_desc;type:text;"`                              //企业介绍
	CorporateVideo      string         `json:"corporateVideo" form:"corporateVideo" gorm:"comment:企业宣传片;column:corporate_video;"`                          //企业宣传片
	Certificate         datatypes.JSON `json:"certificate" form:"certificate" gorm:"comment:荣誉证书;column:certificate;" swaggertype:"array,object"`          //荣誉证书
	DigitalAssetName    *string        `json:"digitalAssetName" form:"digitalAssetName" gorm:"comment:数字资产名称;column:digital_asset_name;"`                  //数字资产名称
	TradingPair         *string        `json:"tradingPair" form:"tradingPair" gorm:"comment:交易对;column:trading_pair;"`                                     //交易对
	IsRecommended      *bool          `json:"isRecommended" form:"isRecommended" gorm:"comment:是否推荐;column:is_recommended;"`                                //是否推荐
	CategoryID         *int64         `json:"categoryId" form:"categoryId" gorm:"comment:所属分类;column:category_id;"`                                         //所属分类
}

// TableName 商家信息 Merchants自定义表名 app_merchants
func (Merchants) TableName() string {
	return "app_merchants"
}
