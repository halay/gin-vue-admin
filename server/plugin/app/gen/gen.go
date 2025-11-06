package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "app", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(new(model.AppUsers), //go:generate go mod tidy
		//go:generate go mod download
		//go:generate go run gen.go

		new(model.Merchants),
	)
	g.Execute()
}
