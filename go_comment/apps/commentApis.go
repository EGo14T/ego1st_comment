package apps

import (
	"blog_comment/entity"
	"blog_comment/service"
	"blog_comment/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommentsGet 获取资源下的所有评论
func CommentsGet(c *gin.Context) {
	showId := c.Param("showId")
	comments := service.QueryComments(showId)
	c.JSON(200, entity.ResOk(comments))
}

// CommentGet 根据ID获取评论数据
func CommentGet(c *gin.Context) {
	commentID := c.Param("id")
	comment, err := service.QueryComment(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ResError(err.Error()))
	} else {
		c.JSON(200, entity.ResOk(comment))
	}
}

// CommentAdd 添加评论
func CommentAdd(c *gin.Context) {
	var addVo vo.AddCommentVo
	err := c.ShouldBind(&addVo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ResError(""))
	} else {
		c.JSON(http.StatusOK, entity.ResOk(service.AddComment(&addVo)))
	}
}

// CommentLike 点赞
func CommentLike(c *gin.Context) {
	commentID := c.Param("id")
	c.JSON(200, entity.ResOk(service.LikeComment(commentID)))
}
