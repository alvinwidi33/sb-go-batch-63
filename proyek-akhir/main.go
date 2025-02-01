package main

import (
	"database/sql"
	"fmt"
	"os"
	"proyek-akhir/controllers"
	"proyek-akhir/database/connections"
	"proyek-akhir/database/migrations"
	"proyek-akhir/middleware"

	"github.com/gin-gonic/gin"
)

var (
	DB *sql.DB
)

func main() {
	connections.Initiator()
	DB = connections.DBConnections
	if DB == nil {
		panic("Database connection is nil")
	}

	_, err := DB.Exec("SET TIME ZONE 'Asia/Jakarta';")
	if err != nil {
		fmt.Println("Failed to set timezone:", err)
	}

	defer DB.Close()

	// Jalankan migrasi database
	migrations.Initiator(DB)

	router := gin.Default()

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware()) // Middleware diterapkan ke semua route dalam group
	{
		protected.GET("/reservation", middleware.AuthMiddleware("Customer", "Admin"), controllers.GetAllReservation)
		protected.GET("/reservation/:id", middleware.AuthMiddleware("Customer", "Admin"), controllers.GetAllReservationByCustomerID)

		protected.POST("/reservation", middleware.AuthMiddleware("Customer"), controllers.InsertReservation)
		protected.PATCH("/reservation/:id/cancel", middleware.AuthMiddleware("Customer"), controllers.CancelReservation)
		protected.PATCH("/reservation/:id/done", middleware.AuthMiddleware("Customer"), controllers.DoneReservation)

		protected.GET("/saloon/customer", middleware.AuthMiddleware("Customer"), controllers.GetAllSaloonCustomers)
		protected.POST("/saloon", middleware.AuthMiddleware("Admin"), controllers.InsertSaloon)
		protected.PUT("/saloon/:id", middleware.AuthMiddleware("Admin"), controllers.UpdateSaloon)
		protected.PATCH("/saloon/:id/delete", middleware.AuthMiddleware("Admin"), controllers.DeleteSaloon)
		protected.GET("/saloon/:id", middleware.AuthMiddleware("Admin", "Customer"), controllers.GetSaloonById)

		protected.GET("/users", middleware.AuthMiddleware("Admin"), controllers.GetAllCustomer)
		protected.PATCH("/users/:id/active", middleware.AuthMiddleware("Admin"), controllers.ActivateUser)
		protected.PATCH("/users/:id/member", middleware.AuthMiddleware("Admin"), controllers.SetCustomerMembership)
		protected.POST("/admin", middleware.AuthMiddleware("Admin"), controllers.RegisterAdmin(DB))
		protected.GET("/admin/saloon", middleware.AuthMiddleware("Admin"), controllers.GetAllSaloon)
	}

	// Route tanpa middleware
	router.POST("/api/login", controllers.Login(DB))
	router.POST("/api/register", controllers.Register(DB))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
