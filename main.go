package main

import (
	. "to-do-listik/database"
	. "to-do-listik/routes"
)

func main() {
	ConnectDB()

	r := SetupRouter()
	r.Run(":8080")
}
