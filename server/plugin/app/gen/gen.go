package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "app", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(new(model.AppUsers), new(model.Merchants), new(model.Banner), new(model.MerchantAnnouncement), new(model.MerchantAdmin), new(model.ProductCategory), new(model.Product), new(model.ProductSku), new(model.ProductSpec), new(model.ProductSpecOption), new(model.Order), new(model.OrderItem), new(model.UserPointsAccount), new(model.UserPointsLog), new(model.AppRelease), new(model.MerchantCategory), new(model.MembershipLevel), new(model.Node), new(model.PointsSettings), new(model.MerchantPointsSettings), new(model.UserAddress), //go:generate go mod tidy
		//go:generate go mod download
		//go:generate go run gen.go

		new(model.AgentLevel),
	)
	g.Execute()
}
