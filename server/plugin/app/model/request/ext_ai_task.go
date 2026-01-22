package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ExtAiTaskSearch struct {
	request.PageInfo
}

type ExtAiImage struct {
	Prompt      string `json:"prompt"`
	ImageUrl    string `json:"image_url"`
	AspectRatio string `json:"aspect_ratio"`
}
