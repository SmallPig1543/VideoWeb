package api

import (
	"VideoWeb2/response"
	"VideoWeb2/service"
	"VideoWeb2/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateVideoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.VideoCreateRequest
		if err := ctx.ShouldBind(&request); err == nil {
			var videoServ service.VideoService
			resp, err := videoServ.VideoCreate(ctx.Request.Context(), &request)
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

func WatchVideoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.VideoWatchRequest
		if err := ctx.ShouldBind(&request); err == nil {
			var videoServ service.VideoService
			resp, err := videoServ.VideoWatch(ctx.Request.Context(), &request)
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

func RankVideoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.VideoRankRequest
		if err := ctx.ShouldBind(&request); err == nil {
			var videoServ service.VideoService
			resp, err := videoServ.VideoRank(ctx.Request.Context(), &request)
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

func QueryVideosHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var filter types.Filter
		if err := ctx.ShouldBind(&filter); err == nil {
			var videoServ service.VideoService
			resp, err := videoServ.QueryVideos(ctx.Request.Context(), &filter)
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
