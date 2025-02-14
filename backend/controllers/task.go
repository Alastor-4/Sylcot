package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	// var tasks []models.Task
	// db := database.DB
	// userID := c.MustGet("userID").(uint)
	// db = db.Where("user_id = ?", userID)

	// if categoryID := c.Query("category_id"); categoryID != "" {
	// 	db = db.Where("categorie_id = ?", categoryID)
	// }

	// if status := c.Query("status"); status != "" {
	// 	db = db.Where("status = ?", status == "true")
	// }

	// if err := db.Preload("Categorie").Find(&tasks).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, "tasks")
}

func CreateTask(c *gin.Context) {

}
func UpdateTask(c *gin.Context) {

}
func DeleteTask(c *gin.Context) {

}
func ToggleTask(c *gin.Context) {

}
