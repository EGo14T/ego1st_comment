package routes

import (
	"blog_comment/apps"
	"blog_comment/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	// 评论模块
	r.GET("/comment/:id", apps.CommentGet)
	r.GET("/comments/:showId", apps.CommentsGet)
	r.POST("/comments", apps.CommentAdd)
	r.GET("/commentLike/:id", apps.CommentLike)

	// 用户模块
	r.GET("/uuid/get", apps.CommentUuidGet)
	r.GET("/getUserInfo/:uuid", apps.UserInfoGet)
	return r
}
