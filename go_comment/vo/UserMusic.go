package vo

type AddCommentVo struct {
	ShowId     string `json:"showId"`
	FromUserId string `json:"fromUserId"`
	ToUserId   string `json:"toUserId"`
	ToId       string `json:"toId"`
	Content    string `json:"content"`
	Like       uint8  `json:"like"`
	Status     uint8  `json:"status"`
}
