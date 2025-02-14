package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"size:255" json:"name"`
	Email      string `gorm:"unique;size:255" json:"email"`
	Image      string `gorm:"size:255" json:"image"`
	Password   string `gorm:"size:255" json:"-"`
	IsVerified bool   `gorm:"default:false" json:"is_verified"`
	Token      string `gorm:"size:255" json:"token"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
