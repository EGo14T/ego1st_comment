package util

import "time"

func TimParse(tt *string) string {
	parse, _ := time.Parse("2006-01-02T15:04:05Z07:00", *tt)
	*tt = parse.Format("2006-01-02 15:04:05")
	return *tt
}

func TimeFormat(tt time.Time) string {
	return tt.Format("2006-01-02 15:04:05")
}
