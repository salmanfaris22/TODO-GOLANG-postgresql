package services

import (
	"ORM-Go/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, name, email string) {
	user := models.UserTest{Name: name, Email: email}
	res := db.Create(&user)
	if res.Error != nil {
		fmt.Println("Error creating user:", res.Error)
	} else {
		fmt.Println("User added to the database successfully!")
	}
}

func FetchUserByID(db *gorm.DB, id int) {
	var user models.UserTest
	res := db.Where("ID = ?", id).Find(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		fmt.Println("Error fetching user or user not found:", res.Error)
	} else {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

func FetchAllUsers(db *gorm.DB) {
	var users []models.UserTest
	db.Find(&users)
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

func UpdateUser(db *gorm.DB, id int, name, email string) {
	res := db.Model(&models.UserTest{}).Where("id = ?", id).Updates(models.UserTest{Name: name, Email: email})
	if res.Error != nil {
		fmt.Println("Error updating user:", res.Error)
	} else {
		fmt.Println("User updated successfully!")
	}
}

func DeleteUser(db *gorm.DB, id int) {
	res := db.Where("id = ?", id).Delete(&models.UserTest{})
	if res.Error != nil {
		fmt.Println("Error deleting user:", res.Error)
	} else {
		fmt.Println("User deleted successfully!")
	}
}

func USerExitsted(db *gorm.DB, mail string) error {
	var user models.UserTest
	res := db.Where("email = ?", mail).Find(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil
	} else {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
		return errors.New("this mail alredy exited")
	}

}
