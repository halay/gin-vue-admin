package initialize

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(model.AppUsers{}, model.Merchants{}, model.Banner{}, model.MerchantAnnouncement{}, model.MerchantAdmin{})
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
