package service

import (
	"blog_comment/dao"
	"blog_comment/entity"
	"blog_comment/util"
	"time"
)

func InitUser(uuid string) string {
	userEntity := entity.UserEntity{
		UUID:       uuid,
		Nickname:   uuid,
		Avatar:     "https://cdn.ego1st.cn/xinmusic/useravatar/defaultAvatar.jpg",
		Email:      "",
		CreateTime: util.TimeFormat(time.Now()),
	}
	return dao.InitUser(&userEntity)

}

func GetUserInfo(uuid string) *entity.UserEntity {
	return dao.GetUserInfo(uuid)
}
