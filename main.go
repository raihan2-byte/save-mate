package main

import (
	"SaveMate/database"
	"SaveMate/router"
)

func main() {
	db := database.InitDatabase()

	router := router.SetupRouter(db)

	router.Run(":8080")
}
