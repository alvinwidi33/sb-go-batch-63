package main

import (
    "database/sql"
    "fmt"
    "os"
    "Quiz-3/controllers"
    "Quiz-3/database"
    "Quiz-3/middleware"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

var (
    DB  *sql.DB
    err error
)

func main() {

    psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
        os.Getenv("PGHOST"),
        os.Getenv("PGPORT"),
        os.Getenv("PGUSER"),
        os.Getenv("PGPASSWORD"),
        os.Getenv("PGDATABASE"),
    )

    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer DB.Close()

    err = DB.Ping()
    if err != nil {
        panic(err)
    }

    database.DBMigrate(DB)

    router := gin.Default()

    protected := router.Group("/api")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/categories", controllers.GetAllCategories)
        protected.POST("/categories", controllers.InsertCategory)
        protected.PUT("/categories/:id", controllers.UpdateCategory)
        protected.DELETE("/categories/:id", controllers.DeleteCategory)
        protected.GET("/categories/:id/books", controllers.GetBookByCategoryID)

        protected.GET("/books", controllers.GetAllBooks)
        protected.POST("/books", controllers.InsertBook)
        protected.PUT("/books/:id", controllers.UpdateBook)
        protected.DELETE("/books/:id", controllers.DeleteBook)
    }

    router.POST("/api/login", controllers.Login(DB))
    router.POST("/api/register", controllers.Register(DB))

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router.Run(":" + port)
}
