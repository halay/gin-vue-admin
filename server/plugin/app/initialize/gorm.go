package initialize

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
    err := global.GVA_DB.WithContext(ctx).AutoMigrate(model.AppUsers{}, model.Merchants{}, model.Banner{}, model.MerchantAnnouncement{}, model.MerchantAdmin{}, model.ProductCategory{}, model.Product{}, model.ProductSku{}, model.ProductSpec{}, model.ProductSpecOption{}, model.ProductSkuOption{}, model.AppConsultation{}, model.PointsConfig{}, model.Order{}, model.OrderItem{}, model.UserPointsAccount{}, model.UserPointsLog{}, model.AppRelease{}, model.MerchantCategory{}, model.MembershipLevel{}, model.Node{})
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
