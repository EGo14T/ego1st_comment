package service

import (
	"blog_comment/dao"
	"blog_comment/entity"
	"blog_comment/util"
	"blog_comment/vo"
	"time"
)

func QueryComments(showId string) *[]vo.CommentsResponseVo {

	comments := *dao.QueryAll(showId)

	commentMap := make(map[string]entity.Comments)

	var res []vo.CommentsResponseVo

	for i := 0; i < len(comments); i++ {
		//parse, _ := time.Parse("2006-01-02 15:04:05", comment.CreateTime)
		comments[i].CreateTime = util.TimParse(&comments[i].CreateTime)
		commentMap[comments[i].ID] = comments[i]
	}

	for i := 0; i < len(comments); i++ {
		resVo := vo.CommentsResponseVo{}
		if comments[i].ToId != "" {
			item := commentMap[comments[i].ToId]
			resVo.OriginComments = &item
		} else {
			resVo.OriginComments = nil
		}
		resVo.ReplyComments = &comments[i]

		res = append(res, resVo)
	}

	return &res
}

func QueryComment(commentID string) *entity.Comments {
	comment := dao.Query(commentID)
	comment.CreateTime = util.TimParse(&comment.CreateTime)
	return comment
}

func AddComment(comment *vo.AddCommentVo) string {
	node, err := util.NewWorker(1)
	if err != nil {
		panic(err)
	}

	commentAdd := entity.Comments{
		ID:         node.GetId(),
		ShowId:     comment.ShowId,
		FromUserId: comment.FromUserId,
		ToUserId:   comment.ToUserId,
		ToId:       comment.ToId,
		Content:    comment.Content,
		Like:       0,
		Status:     1,
		CreateTime: util.TimeFormat(time.Now()),
	}

	return dao.Insert(&commentAdd)
}

func LikeComment(commentID string) string{
	return dao.LikeComment(commentID)
}
