package main

import (
	"community-marketplace/config"
	"community-marketplace/controllers"
	"community-marketplace/middleware"
	"community-marketplace/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"community-marketplace/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Run Migrations
	err := config.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Item{}, &models.Message{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Database migration completed")

	// Seed Admin User
	seedAdmin()

	r := gin.Default()

	// CORS Setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Routes
	api := r.Group("/api")
	{
		// Public routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		items := api.Group("/items")
		{
			items.GET("/", controllers.GetItems)
			items.GET("/:id", controllers.GetItem)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", controllers.GetCategories)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// User Profile
			protected.GET("/auth/me", func(c *gin.Context) {
				userId := c.MustGet("userID").(uint)
				var user models.User
				config.DB.First(&user, userId)
				c.JSON(http.StatusOK, user)
			})

			// Item Management
			protected.POST("/items", controllers.CreateItem)
			protected.PUT("/items/:id", controllers.UpdateItem)
			protected.DELETE("/items/:id", controllers.DeleteItem)
			protected.GET("/my-listings", controllers.GetMyListings)

			// Messaging
			protected.POST("/messages", controllers.SendMessage)
			protected.GET("/messages", controllers.GetConversations)
			protected.GET("/messages/:other_user_id", controllers.GetThread)

			// Admin routes
			admin := protected.Group("/admin")
			admin.Use(middleware.AdminMiddleware())
			{
				admin.GET("/items", controllers.GetItems) // Admin can see all
				admin.POST("/categories", controllers.CreateCategory)
				admin.PUT("/categories/:id", controllers.UpdateCategory)
				admin.DELETE("/categories/:id", controllers.DeleteCategory)
			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	r.Run(":" + port)
}

func seedAdmin() {
	var count int64
	config.DB.Model(&models.User{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		hashedPassword, _ := utils.HashPassword("admin123")
		admin := models.User{
			Name:         "Admin",
			Email:        "admin@example.com",
			PasswordHash: hashedPassword,
			Role:         "admin",
		}
		config.DB.Create(&admin)
		fmt.Println("Default admin user created: admin@example.com / admin123")
	}
}
