package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Order 订单 结构体
type Order struct {
	global.GVA_MODEL
	OrderNo        *string  `json:"orderNo" form:"orderNo" gorm:"comment:订单号;column:order_no;size:64;" binding:"required"`                          //订单号
	UserID         *int64   `json:"userId" form:"userId" gorm:"comment:app用户ID;column:user_id;" binding:"required"`                                 //下单用户
	MerchantID     *int64   `json:"merchantId" form:"merchantId" gorm:"comment:商户ID;column:merchant_id;" binding:"required"`                        //商户
	TotalAmount    *float64 `json:"totalAmount" form:"totalAmount" gorm:"comment:金额汇总(现金);column:total_amount;"`                                    //总金额
	TotalPoints    *int64   `json:"totalPoints" form:"totalPoints" gorm:"comment:积分汇总;column:total_points;"`                                        //总积分
	PayMethod      *string  `json:"payMethod" form:"payMethod" gorm:"default:points;comment:支付方式;column:pay_method;" binding:"required"`            //支付方式
	PayStatus      *string  `json:"payStatus" form:"payStatus" gorm:"default:unpaid;comment:未支付/已支付/退款中/已退款;column:pay_status;" binding:"required"` //支付状态
	Status         *string  `json:"status" form:"status" gorm:"default:created;comment:创建/已支付/已发货/已完成/已取消/已退款;column:status;" binding:"required"`   //订单状态
	Province       *string  `json:"province" form:"province" gorm:"comment:省;column:province;size:64;"`                                             //省
	City           *string  `json:"city" form:"city" gorm:"comment:市;column:city;size:64;"`                                                         //市
	District       *string  `json:"district" form:"district" gorm:"comment:区县;column:district;size:64;"`                                            //区县
	Country        *string  `json:"country" form:"country" gorm:"comment:国家;column:country;size:64;"`                                               //国家
	ConsigneeName  *string  `json:"consigneeName" form:"consigneeName" gorm:"comment:收货人;column:consignee_name;size:64;"`                           //收货人
	ConsigneePhone *string  `json:"consigneePhone" form:"consigneePhone" gorm:"comment:收货人手机号;column:consignee_phone;size:20;"`                     //收货人手机号
	Address        *string  `json:"address" form:"address" gorm:"comment:收货地址;column:address;size:255;"`                                            //收货地址
	PostalCode     *string  `json:"postalCode" form:"postalCode" gorm:"comment:邮政编码;column:postal_code;size:20;"`                                   //邮政编码
	DeliveryStatus *string  `json:"deliveryStatus" form:"deliveryStatus" gorm:"default:none;comment:未发货/已发货/已签收;column:delivery_status;"`           //发货状态
	ExpressName    *string  `json:"expressName" form:"expressName" gorm:"comment:快递公司;column:express_name;size:64;"`                                //快递公司
	ExpressNo      *string  `json:"expressNo" form:"expressNo" gorm:"comment:快递单号;column:express_no;size:64;"`                                      //快递单号
	Remark         *string  `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:255;"`                                                 //备注
}

// TableName 订单 Order自定义表名 app_orders
func (Order) TableName() string {
	return "app_orders"
}
