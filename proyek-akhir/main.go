package main

import (
	"database/sql"
	"os"
	"proyek-akhir/controllers"
	"proyek-akhir/database/connections"
	"proyek-akhir/database/migrations"
	"proyek-akhir/middleware"

	"github.com/gin-gonic/gin"
	"log"
	"time"
)
var (
    DB *sql.DB
)

func init() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("Failed to load timezone, defaulting to WIB (GMT+7)")
		loc = time.FixedZone("WIB", 7*60*60) // Default ke GMT+7
	}
	time.Local = loc
}

func main() {
	connection.Initiator()

	DB = connection.DBConnections
	if DB == nil {
		panic("Database connection is nil")
	}
	_, err := DB.Exec("SET TIMEZONE TO 'Etc/GMT-7';") 
	if err != nil {
		log.Println("Failed to set timezone in database:", err)
	}
	
	defer DB.Close()
	migration.Initiator(DB)

	router := gin.Default()

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware()) 
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
        protected.GET("/saloon/:id", middleware.AuthMiddleware("Admin","Customer"), controllers.GetSaloonById)

		protected.GET("/users", middleware.AuthMiddleware("Admin"), controllers.GetAllCustomer)
		protected.PATCH("/users/:id/active", middleware.AuthMiddleware("Admin"), controllers.ActivateUser)
		protected.PATCH("/users/:id/member", middleware.AuthMiddleware("Admin"), controllers.SetCustomerMembership)
        protected.POST("/admin", middleware.AuthMiddleware("Admin"), controllers.RegisterAdmin(DB))
        protected.GET("/admin/saloon", middleware.AuthMiddleware("Admin"), controllers.GetAllSaloon)
	}
	router.POST("/api/login", controllers.Login(DB))
	router.POST("/api/register", controllers.Register(DB))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
