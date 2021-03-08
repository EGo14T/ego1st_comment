package entity

type UserEntity struct {
	UUID       string `json:"uuid"`
	Nickname   string `json:"nickName"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	CreateTime string `json:"createTime"`
}

func (m *UserEntity) TableName() string {
	return "base_user"
}
