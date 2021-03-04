package main

import (
	"blog_comment/database"
	"blog_comment/routes"
)

func main() {
	database.InitDB()
	router := routes.InitRouter()
	_ = router.Run(":9876")
}
