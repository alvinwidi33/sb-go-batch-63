package controllers

import (
	"Quiz-3/database"
	"Quiz-3/repository"
	"Quiz-3/structs"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
    var (
       result gin.H
    )

    book, err := repository.GetAllBooks(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": book,
       }
    }

    c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
    var book structs.Books

    err := c.BindJSON(&book)
    if err != nil {
       panic(err)
    }

    err = repository.InsertBook(database.DbConnection, book)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
    var book structs.Books
    id, _ := strconv.Atoi(c.Param("id"))

    err := c.BindJSON(&book)
    if err != nil {
       panic(err)
    }

    book.ID = id

    err = repository.UpdateBook(database.DbConnection, book)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
    var book structs.Books
    id, _ := strconv.Atoi(c.Param("id"))

    book.ID = id
    err := repository.DeleteBook(database.DbConnection, book)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, book)
}


type BooksController struct {
	BooksRepo *repository.BookRepository
}

func NewBooksController(repo *repository.BookRepository) *BooksController {
	return &BooksController{BooksRepo: repo}
}

func GetBookByCategoryID(c *gin.Context) {
   categoryIDParam := c.Param("id")
   categoryID, err := strconv.Atoi(categoryIDParam)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
       return
   }


   books, err := repository.GetBooksByCategoryID(database.DbConnection, categoryID)
   if err != nil {
       c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
       return
   }

   if len(books) == 0 {
       c.JSON(http.StatusNotFound, gin.H{"error": "no books found for this category"})
       return
   }

   c.JSON(http.StatusOK, gin.H{"result": books})
}

