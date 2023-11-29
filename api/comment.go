package api

import (
	"VideoWeb2/response"
	"VideoWeb2/service"
	"VideoWeb2/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendCommentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CommentRequest
		if err := ctx.ShouldBind(&req); err == nil {
			var commentServ service.CommentServ
			resp, err := commentServ.CommentCreate(ctx.Request.Context(), &req)
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
