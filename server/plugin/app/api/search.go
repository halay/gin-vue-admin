package api

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var Search = new(SEARCH)

type SEARCH struct{}

func (a *SEARCH) GetSearchPublic(c *gin.Context) {
    ctx := c.Request.Context()
    var info request.SearchPublic
    if err := c.ShouldBindQuery(&info); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    list, total, guess, hot, err := serviceSearch.GetSearchPublic(ctx, info)
    if err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:"+err.Error(), c)
        return
    }
    response.OkWithDetailed(gin.H{
        "list": gin.H{
            "items":    list,
            "total":    total,
            "page":     info.Page,
            "pageSize": info.PageSize,
        },
        "guess": guess,
        "hot":   hot,
    }, "获取成功", c)
}
