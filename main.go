package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func init() {
	dsn := "host=localhost user=postgres password=poomon dbname=management  port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	err = db.AutoMigrate(&UserTest{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
}

type UserTest struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"<-:create"`
	Email string `gorm:"<-:create"`
}

func main() {

	showTodoList()
}

func showTodoList() {
	println("Enter List")
	for {
		var num int
		println("1:ADD")
		println("2:GET WITH SELECT")
		println("3:GET ALL")
		println("4:UPDATE WITH ID")
		println("5:DELET WITH ID")
		println("6:EXIT")
		_, err = fmt.Scan(&num)
		if err != nil {
			fmt.Println("enter Valid Number", err.Error())
			return
		}
		switch num {
		case 1:
			fmt.Println("creating DB")
			creatdbmodel(db)
		case 2:
			fmt.Println("getting person")
			fetchwithContionrecored(db)
		case 3:
			fmt.Println("Getting All")
			fetchAllRecord(db)

		case 4:
			updateRecord(db)

		case 5:
			deleteRecord(db)

		case 6:
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("pleas valid number")
		}
	}
}

func creatdbmodel(db *gorm.DB) {
	var name, email string
	fmt.Println("Enter Your Name")
	_, err = fmt.Scan(&name)
	if err != nil {
		fmt.Println("enter Valid Name", err.Error())
		return
	}
	fmt.Println("Enter Your Email")
	_, err = fmt.Scan(&email)
	if err != nil {
		fmt.Println("enter Valid email", err.Error())
		return
	}

	u := UserTest{Name: name, Email: email}
	res := db.Create(&u)
	fmt.Println("Added To db Done !")
	fmt.Println(res.Error)
}

func fetchwithContionrecored(db *gorm.DB) {
	var id int
	fmt.Println("Enter ID ")
	_, err = fmt.Scan(&id)
	if err != nil {
		fmt.Println("enter Valid Number", err.Error())
		return
	}
	var u UserTest
	tc := db.Where("ID=?", id).Find(&u)
	fmt.Println(tc.RowsAffected)
	fmt.Println("____________")
	fmt.Println(u)
	fmt.Println("____________")
}

func fetchAllRecord(db *gorm.DB) {
	var u []UserTest
	db.Find(&u)
	for _, v := range u {
		fmt.Println(v.ID, ":", v.Name, "|", v.Email)
	}
	//fmt.Println(u)
}

func updateRecord(db *gorm.DB) {
	var id int
	var name, email string
	fmt.Println("Enter The ID For Update")
	_, err = fmt.Scan(&id)
	if err != nil {
		fmt.Println("enter Valid Number", err.Error())
		return
	}
	fmt.Println("Enter The new Name ")
	_, err = fmt.Scan(&name)
	if err != nil {
		fmt.Println("enter Valid Number", err.Error())
		return
	}
	fmt.Println("Enter The new Email ")
	_, err = fmt.Scan(&email)
	if err != nil {
		fmt.Println("enter Valid Number", err.Error())
		return
	}
	res := db.Table("user_tests").Where("id=?", id).Updates(map[string]interface{}{"name": name, "email": email})
	fmt.Println(res, res.Error, res.RowsAffected)
	fmt.Println("Updated")
}

func deleteRecord(db *gorm.DB) {
	var id int
	fmt.Println("Enter Id For Delete")
	_, err = fmt.Scan(&id)
	if err != nil {
		fmt.Println("enter Valid Number", err.Error())
		return
	}
	res := db.Where("id=?", id).Delete(&UserTest{})
	fmt.Println(res, res.Error, res.RowsAffected)
	fmt.Println("Deleted")
}
