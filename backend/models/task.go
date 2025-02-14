package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string    `gorm:"size:255" json:"title"`
	Priority   uint      `gorm:"" json:"priority"`
	Status     bool      `gorm:"default:false"`
	CategoryID uint      `gorm:"foreignKey" json:"category_id"`
	UserID     uint      `gorm:"foreignKey" json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func MigrateTasks(db *gorm.DB) error {
	err := db.AutoMigrate(&Task{})
	if err != nil {
		return err
	}
	return nil
}
