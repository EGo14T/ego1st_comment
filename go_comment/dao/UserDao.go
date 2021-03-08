package dao

import (
	"blog_comment/database"
	"blog_comment/entity"
)

func InitUser(userEntity *entity.UserEntity) string {
	db := database.GetMysqlDb()
	result := db.Select(
		"uuid",
		"nickname",
		"avatar",
		"email",
		"create_time").Create(userEntity)
	if result.Error != nil {
		return ""
	} else {
		return userEntity.UUID
	}
}

func GetUserInfo(uuid string) *entity.UserEntity {
	db := database.GetMysqlDb()
	var userEntity entity.UserEntity
	db.Raw("select * from base_user where uuid = ?", uuid).Scan(&userEntity)
	return &userEntity
}
