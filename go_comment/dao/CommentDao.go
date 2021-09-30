package dao

import (
	"blog_comment/database"
	"blog_comment/entity"
	"errors"
)

func QueryAll(showId string) *[]entity.Comments {
	db := database.GetMysqlDb()
	var comments []entity.Comments

	db.Raw(
		"SELECT a.*,"+
			"b.nickname, "+
			"b.avatar "+
			"FROM comments a "+
			"LEFT JOIN base_user b "+
			"ON "+
			"b.uuid = a.from_user_id "+
			"WHERE a.show_id = ? ORDER BY a.create_time DESC", showId).Scan(&comments)
	return &comments
}

func Query(commentID string) (*entity.Comments, error) {
	db := database.GetMysqlDb()
	var comment entity.Comments

	scan := db.Raw(
		"SELECT a.*,"+
			"b.nickname, "+
			"b.avatar "+
			"FROM comments a "+
			"LEFT JOIN base_user b "+
			"ON "+
			"b.uuid = a.from_user_id "+
			"WHERE a.id = ?", commentID).Scan(&comment)

	if scan.RowsAffected == 0 {
		return nil, errors.New("查询数据为空")
	}
	return &comment, nil
}

func Insert(comment *entity.Comments) string {
	db := database.GetMysqlDb()
	result := db.Select(
		"id",
		"show_id",
		"from_user_id",
		"to_user_id",
		"to_id",
		"content",
		"like",
		"status",
		"create_time").Create(comment)
	if result.Error != nil {
		return ""
	} else {
		return comment.ID
	}
}

func LikeComment(commentID string) string {
	db := database.GetMysqlDb()
	result := db.Exec("UPDATE comments t SET t.LIKE = t.LIKE + 1 WHERE t.id =?", commentID)
	if result.Error != nil {
		return "-1"
	} else {
		return "1"
	}
}
