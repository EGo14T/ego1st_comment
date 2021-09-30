package vo

import (
	"blog_comment/entity"
)

type CommentsResponseVo struct {
	ReplyComments  *entity.Comments `json:"replyComments"`
	OriginComments *entity.Comments `json:"originComments"`
}
