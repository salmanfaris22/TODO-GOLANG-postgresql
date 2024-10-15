package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"<-:create"`
}

func main() {
	dsn := "host=localhost user=postgres password=poomon dbname=management  port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	err = db.AutoMigrate(&Person{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	u := Person{Name: "rishid"}
	res := db.Create(&u)
	fmt.Println(res.Error)

}
