package utils

import (
	"ORM-Go/services"
	"fmt"
	"gorm.io/gorm"
)

func ShowTodoList(db *gorm.DB) {
	for {
		var num int
		fmt.Println("1: ADD")
		fmt.Println("2: GET WITH SELECT")
		fmt.Println("3: GET ALL")
		fmt.Println("4: UPDATE WITH ID")
		fmt.Println("5: DELETE WITH ID")
		fmt.Println("6: EXIT")
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Please enter a valid number:", err.Error())
			return
		}
		switch num {
		case 1:
			var name, email string
			fmt.Println("Enter your name:")
			fmt.Scan(&name)
			fmt.Println("Enter your email:")
			fmt.Scan(&email)
			err := services.USerExitsted(db, email)
			if err != nil {
				fmt.Println(err)

			} else {
				services.CreateUser(db, name, email)
			}

		case 2:
			var id int
			fmt.Println("Enter the ID:")
			fmt.Scan(&id)
			services.FetchUserByID(db, id)
		case 3:
			services.FetchAllUsers(db)
		case 4:
			var id int
			var name, email string
			fmt.Println("Enter the ID to update:")
			fmt.Scan(&id)
			fmt.Println("Enter the new name:")
			fmt.Scan(&name)
			fmt.Println("Enter the new email:")
			fmt.Scan(&email)
			err := services.USerExitsted(db, email)
			if err != nil {
				fmt.Println(err)

			} else {
				services.UpdateUser(db, id, name, email)
			}

		case 5:
			var id int
			fmt.Println("Enter the ID to delete:")
			fmt.Scan(&id)
			services.DeleteUser(db, id)
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Please enter a valid number")
		}
	}
}
