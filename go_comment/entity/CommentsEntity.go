package entity

type Comments struct {
	ID         string `json:"id"`
	ShowId     string `json:"showId"`
	FromUserId string `json:"fromUserId"`
	ToUserId   string `json:"toUserId"`
	ToId       string `json:"toId"`
	Content    string `json:"content"`
	Like       int64  `json:"like"`
	Status     uint8  `json:"status"`
	CreateTime string `json:"createTime"`
	Nickname   string `json:"nickName"`
	Avatar     string `json:"avatar"`
}

func (m *Comments) TableName() string {
	return "comments"
}
