package initialize

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/plugin"
)

func Viper() {
	err := global.GVA_VP.UnmarshalKey("app-user", &plugin.Config)
	if err != nil {
		err = errors.Wrap(err, "初始化配置文件失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
