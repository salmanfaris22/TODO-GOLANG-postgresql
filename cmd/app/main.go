package main

import (
	"ORM-Go/config"
	"ORM-Go/utils"
)

func main() {
	db := config.InitDB()  // Initialize the database connection
	utils.ShowTodoList(db) // Start the CLI menu
}
