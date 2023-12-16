package main

import (
  "log"
  "os"

  "github.com/gin-gonic/gin"
  "jwt-gin/controllers"
  "jwt-gin/models"
  "jwt-gin/middlewares"
  "github.com/joho/godotenv"
)

func main() {
	// Create a new gin instance
	r := gin.Default()

	public := r.Group("/api")

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)


	public.POST("/register", controllers.Register)
  public.POST("/login",controllers.Login)
  protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user",controllers.CurrentUser)
	// Run the server

	r.Run(":8080")
}
