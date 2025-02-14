package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"size:50" json:"title"`
	Color    string `gorm:"size:50" json:"color"`
	IconName string `gorm:"size:50" json:"icon_name"`
}

func MigrateCategories(db *gorm.DB) error {
	err := db.AutoMigrate(&Category{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) Setup(db *gorm.DB) error {
	categories := []Category{
		{Title: "Work", Color: "#3498db", IconName: "work"},
		{Title: "Home", Color: "#2ecc71", IconName: "home"},
		{Title: "Love", Color: "#e74c3c", IconName: "love"},
	}

	for _, categorie := range categories {
		result := db.FirstOrCreate(&categorie, Category{Title: categorie.Title})
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
