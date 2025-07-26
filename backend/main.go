package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"it-order-app/backend/config"
	"it-order-app/backend/router"
	"log"
)

func main() {
	// Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi ke database
	config.ConnectDatabase()

	// Inisialisasi Gin
	r := gin.Default()

	// Setup CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Setup router
	router.SetupRouter(r, config.DB)

	// Jalankan server
	r.Run(":8080") // http://localhost:8080
}
