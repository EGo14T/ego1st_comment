package apps

import (
	"blog_comment/entity"
	"blog_comment/service"
	"blog_comment/util"
	"github.com/gin-gonic/gin"
)

// 第一次访问获取uuid存储到浏览器缓存里
func CommentUuidGet(c *gin.Context) {
	node, err := util.NewWorker(10)
	if err != nil {
		panic(err)
	}
	uuid := node.GetId()
	c.JSON(200, entity.ResOk(service.InitUser(uuid)))
}

// 获取用户信息
func UserInfoGet(c *gin.Context) {
	uuid := c.Param("uuid")
	c.JSON(200, entity.ResOk(service.GetUserInfo(uuid)))
}
