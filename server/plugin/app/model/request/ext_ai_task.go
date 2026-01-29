package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ExtAiTaskSearch struct {
	request.PageInfo
	UserId string `json:"userId" form:"userId"`
	Type   string `json:"type" form:"type"`
	Sort   string `json:"sort" form:"sort"`
}

type ExtAiImage struct {
	Prompt      string `json:"prompt"`
	ImageUrl    string `json:"image_url"`
	AspectRatio string `json:"aspect_ratio"`
}
