package route

import (
	"VideoWeb2/api"
	"VideoWeb2/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret")) // 存储
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api")
	{
		//用户登录注册，上传或者修改头像
		user := v1.Group("user")
		{
			user.GET("register", api.RegisterHandler())
			user.GET("login", api.LoginHandler())
			user.PUT("upload", middleware.JWT(), api.UploadHandler())
		}
		//聊天模块
		chat := v1.Group("chat")
		chat.Use(middleware.JWT())
		{
			chat.POST("send", api.SendMessageHandler())     //发送消息
			chat.POST("search", api.QueryMessagesHandler()) //根据时间搜索聊天记录
		}
		//视频模块
		video := v1.Group("video")
		video.GET("watch", api.WatchVideoHandler()) //观看视频
		video.GET("rank", api.RankVideoHandler())   //返回排行榜，只返回vid和点击量
		video.Use(middleware.JWT())
		{
			video.PUT("create", api.CreateVideoHandler())
			video.POST("comment", api.SendCommentHandler()) //评论视频
			video.POST("reply", api.SendReplyHandler())     //回复评论
		}
		//搜索模块
		search := v1.Group("search")
		search.Use(middleware.JWT())
		{
			search.POST("video", api.QueryVideosHandler())
		}
	}
	return r
}
