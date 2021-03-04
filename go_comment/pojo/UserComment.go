package pojo

type UserComment struct {
	ShowId     string
	FromUserId string
	ToUserId   string
	ToId       string
	Content    string
	Like       uint8
	Status     uint8
	NickName   string
	Avatar     string
}
