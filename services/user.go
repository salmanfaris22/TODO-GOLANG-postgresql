package services

import (
	"ORM-Go/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// CreateUser adds a user to the database, with color-coded messages.
func CreateUser(db *gorm.DB, name, email string) {
	user := models.UserTest{Name: name, Email: email}
	res := db.Create(&user)
	if res.Error != nil {
		fmt.Println("\033[31mError creating user:\033[0m", res.Error) // Red for errors
	} else {
		fmt.Println("\033[32mUser added to the database successfully!\033[0m") // Green for success
	}
}

// FetchUserByID fetches a user by ID, with color-coded messages.
func FetchUserByID(db *gorm.DB, id int) {
	var user models.UserTest
	res := db.Where("ID = ?", id).Find(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		fmt.Println("\033[31mError fetching user or user not found:\033[0m", res.Error) // Red for errors
	} else {
		fmt.Printf("\033[36mID: %d, Name: %s, Email: %s\033[0m\n", user.ID, user.Name, user.Email) // Cyan for info
	}
}

// FetchAllUsers fetches and prints all users, with color-coded messages.
func FetchAllUsers(db *gorm.DB) {
	var users []models.UserTest
	db.Find(&users)
	if len(users) == 0 {
		fmt.Println("\033[31mNo users found.\033[0m") // Red for no records
	} else {
		fmt.Println("\033[36mFetching all users:\033[0m") // Cyan for info
		for _, user := range users {
			fmt.Printf("\033[36mID: %d, Name: %s, Email: %s\033[0m\n", user.ID, user.Name, user.Email)
		}
	}
}

// UpdateUser updates a user record by ID, with color-coded messages.
func UpdateUser(db *gorm.DB, id int, name, email string) {
	res := db.Model(&models.UserTest{}).Where("id = ?", id).Updates(models.UserTest{Name: name, Email: email})
	if res.Error != nil {
		fmt.Println("\033[31mError updating user:\033[0m", res.Error) // Red for errors
	} else {
		fmt.Println("\033[32mUser updated successfully!\033[0m") // Green for success
	}
}

// DeleteUser deletes a user by ID, with color-coded messages.
func DeleteUser(db *gorm.DB, id int) {
	res := db.Where("id = ?", id).Delete(&models.UserTest{})
	if res.Error != nil {
		fmt.Println("\033[31mError deleting user:\033[0m", res.Error) // Red for errors
	} else {
		fmt.Println("\033[32mUser deleted successfully!\033[0m") // Green for success
	}
}

func USerExitsted(db *gorm.DB, email string) error {
	var user models.UserTest
	res := db.Where("email = ?", email).Find(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil
	} else {
		fmt.Printf("\033[36mID: %d, Name: %s, Email: %s\033[0m\n", user.ID, user.Name, user.Email) // Cyan for info
		return errors.New("\033[31mThis email already exists!\033[0m")                             // Red for existing user error
	}
}
