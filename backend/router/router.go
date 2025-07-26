package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"it-order-app/backend/handler"
	"it-order-app/backend/middleware"
	"it-order-app/backend/repository"
	"it-order-app/backend/service"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {

	jwtService := service.NewJWTService()
	// Inisialisasi repository
	userRepository := repository.NewUserRepository(db)
	orderRepository := repository.NewOrderRepository(db)
	statsRepository := repository.NewStatsRepository(db)

	// Inisialisasi handler
	authHandler := handler.NewAuthHandler(userRepository, jwtService)
	userHandler := handler.NewUserHandler(userRepository)
	orderHandler := handler.NewOrderHandler(orderRepository)
	statsHandler := handler.NewStatsHandler(statsRepository)

	// Grouping API v1
	api := r.Group("/api")
	{
		// === Rute Publik (tidak perlu login) ===
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		authRoutes := api.Group("")
		authRoutes.Use(middleware.AuthMiddleware(jwtService))
		{
			authRoutes.GET("/profile", userHandler.GetProfile)

			// Rute untuk Orders
			authRoutes.POST("/orders", orderHandler.CreateOrder)
			authRoutes.GET("/orders", orderHandler.GetUserOrders)
		}

		// Rute terproteksi HANYA untuk ADMIN
		adminRoutes := api.Group("/admin")
		adminRoutes.Use(middleware.AuthMiddleware(jwtService), middleware.AdminMiddleware()) // <-- DUA MIDDLEWARE
		{
			adminRoutes.GET("/orders", orderHandler.GetAllOrders)
			adminRoutes.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)
			adminRoutes.GET("/stats", statsHandler.GetStats)
		}
	}
}
