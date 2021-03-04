package dao

import (
	"blog_comment/database"
	"blog_comment/entity"
)

func QueryAll(showId string) *[]entity.Comments {
	db := database.GetMysqlDb()
	var comments []entity.Comments

	db.Raw(
		"SELECT a.*," +
			"b.nickname, " +
			"b.avatar " +
			"FROM xinmusic_comments a " +
			"LEFT JOIN xinmusic_base_user b " +
			"ON " +
			"b.userid = a.from_user_id " +
			"WHERE a.show_id = ? ORDER BY a.create_time DESC", showId).Scan(&comments)
	return &comments
}

func Query(commentID string) *entity.Comments {
	db := database.GetMysqlDb()
	var comment entity.Comments

	db.Raw(
		"SELECT a.*," +
			"b.nickname, " +
			"b.avatar " +
			"FROM xinmusic_comments a " +
			"LEFT JOIN xinmusic_base_user b " +
			"ON " +
			"b.userid = a.from_user_id " +
			"WHERE a.id = ?", commentID).Scan(&comment)
	return &comment
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
	}else {
		return comment.ID
	}
}

func LikeComment(commentID string) string {
	db := database.GetMysqlDb()
	result := db.Exec("UPDATE xinmusic_comments t SET t.LIKE = t.LIKE + 1 WHERE t.id =?", commentID)
	if result.Error != nil {
		return "-1"
	}else {
		return "1"
	}
}


