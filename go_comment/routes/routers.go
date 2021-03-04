package routes

import (
	"blog_comment/apps"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/comments/:showId", apps.CommentsGet)
	r.GET("/comment/:id", apps.CommentGet)
	r.POST("/comments", apps.CommentAdd)
	r.GET("/commentLike/:id", apps.CommentLike)

	return r
}
