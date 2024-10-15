package models

type UserTest struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"<-:create"`
	Email string `gorm:"<-:create"`
}
