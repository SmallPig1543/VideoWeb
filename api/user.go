package api

import (
	"VideoWeb2/response"
	"VideoWeb2/service"
	"VideoWeb2/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRequest
		if err := ctx.ShouldBind(&req); err == nil {
			var userServ service.UserService
			resp, err := userServ.Register(ctx.Request.Context(), &req)
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
func LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.UserRequest
		if err := ctx.ShouldBind(&request); err == nil {
			var userServ service.UserService
			resp, err := userServ.Login(ctx.Request.Context(), &request)
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

func UploadHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.UserUploadRequest
		if err := ctx.ShouldBind(&request); err == nil {
			var userServ service.UserService
			resp, err := userServ.Upload(ctx.Request.Context(), &request)
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
