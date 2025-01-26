package controllers

import (
	"Quiz-3/database"
	"Quiz-3/repository"
	"Quiz-3/structs"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
    var (
       result gin.H
    )

    category, err := repository.GetAllCategories(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": category,
       }
    }

    c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
    var category structs.Categories

    err := c.BindJSON(&category)
    if err != nil {
       panic(err)
    }

    err = repository.InsertCategory(database.DbConnection, category)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
   var category structs.Categories
   id, _ := strconv.Atoi(c.Param("id"))

   // Bind JSON ke objek category
   err := c.BindJSON(&category)
   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
       return
   }

   category.ID = id

   now := time.Now()
   category.ModifiedAt = &now

   err = repository.UpdateCategory(database.DbConnection, category)
   if err != nil {
       c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
       return
   }

   c.JSON(http.StatusOK, category)
}



func DeleteCategory(c *gin.Context) {
    var category structs.Categories
    id, _ := strconv.Atoi(c.Param("id"))

    category.ID = id
    err := repository.DeleteCategory(database.DbConnection, category)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, category)
}

