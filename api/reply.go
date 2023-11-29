package api

import (
	"VideoWeb2/response"
	"VideoWeb2/service"
	"VideoWeb2/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendReplyHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ReplyRequest
		if err := ctx.ShouldBind(&req); err == nil {
			var replyServ service.ReplyServ
			resp, err := replyServ.ReplyCreate(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, resp)
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusBadRequest, response.BadResponse("参数有误"))
		}
	}
}
