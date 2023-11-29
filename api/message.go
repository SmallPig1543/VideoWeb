package api

import (
	"VideoWeb2/response"
	"VideoWeb2/service"
	"VideoWeb2/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.MessageReq
		if err := ctx.ShouldBind(&req); err == nil {
			var chatServ service.ChatServ
			resp, err := chatServ.SendMessage(ctx.Request.Context(), &req)
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

func QueryMessagesHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter types.Filter
		if err := ctx.ShouldBind(&filter); err == nil {
			var chatServ service.ChatServ
			resp, err := chatServ.QueryMessages(ctx.Request.Context(), &filter)
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
