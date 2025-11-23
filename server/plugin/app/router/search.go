package router

import (
	"github.com/gin-gonic/gin"
)

var Search = new(S)

type S struct{}

func (r *S) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	group := public.Group("search")
	group.GET("public", apiSearch.GetSearchPublic)
}
