package main

import (
	"log"
	"os"

	"github.com/alastor-4/sylcot/backend/controllers"
	"github.com/alastor-4/sylcot/backend/database"
	_ "github.com/alastor-4/sylcot/backend/docs"
	"github.com/alastor-4/sylcot/backend/middleware"
	"github.com/alastor-4/sylcot/backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	// if _, exists := os.LookupEnv("RUNNING_IN_DOCKER"); !exists {
	// 	log.Println("Loading .env file")
	// 	err := godotenv.Load(".env")
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file")
	// 	}
	// }
	// if os.Getenv("ENV") != "production" {
	// 	err := godotenv.Load(".env")
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file")
	// 	}
	// }

	config := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := database.NewConnection(&config)
	if err != nil {
		log.Fatal("Could not connect the database")
	}

	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatalf("Could not migrate user table in db")
	}
	err = models.MigrateTasks(db)
	if err != nil {
		log.Fatalf("Could not migrate task table in db")
	}
	err = models.MigrateCategories(db)
	if err != nil {
		log.Fatalf("Could not migrate category table in db")
	}

	var category models.Category
	if err := category.Setup(db); err != nil {
		panic("Failed to seed categories")
	}

	r := Repository{DB: db}

	authController := &controllers.AuthController{DB: db}

	app := gin.Default()
	r.SetupRoutes(app, authController)
	app.Run(":8080")
}

func (r *Repository) SetupRoutes(app *gin.Engine, authController *controllers.AuthController) {
	auth := app.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.GET("/verify-email/", authController.VerifyEmail)
	}

	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/tasks", controllers.GetTasks)
		api.POST("/tasks", controllers.CreateTask)
		api.PUT("/tasks/:id", controllers.UpdateTask)
		api.DELETE("/tasks/:id", controllers.DeleteTask)
		api.PATCH("/tasks/:id/complete", controllers.ToggleTask)
	}
	app.Use(cors.Default())
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
