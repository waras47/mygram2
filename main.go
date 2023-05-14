package main

import (
	"final_project/database"
	"final_project/router"
)

func main() {
	database.StartDB()

	router.New().Run(":3000")
}
