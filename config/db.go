package config

import (
	"ORM-Go/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=poomon dbname=management port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return nil
	}

	err = db.AutoMigrate(&models.UserTest{})
	if err != nil {
		fmt.Println("Failed to migrate the database:", err)
		return nil
	}

	return db
}
