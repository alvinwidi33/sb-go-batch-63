package controllers

import (
	"net/http"
	"proyek-akhir/database/connections"
	"proyek-akhir/repository"
	"proyek-akhir/structs"
	"strconv"
	"github.com/gin-gonic/gin"
)

func GetAllSaloon(s *gin.Context) {
    var (
       result gin.H
    )

    saloon, err := repository.GetAllSaloon(connection.DBConnections)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": saloon,
       }
    }

    s.JSON(http.StatusOK, result)
}

func GetAllSaloonCustomers(s *gin.Context) {
   var (
      result gin.H
   )

   saloon, err := repository.GetAllSaloonCustomers(connection.DBConnections)

   if err != nil {
      result = gin.H{
         "result": err.Error(),
      }
   } else {
      result = gin.H{
         "result": saloon,
      }
   }

   s.JSON(http.StatusOK, result)
}

func InsertSaloon(s *gin.Context) {
    var saloon structs.Saloon

    err := s.BindJSON(&saloon)
    if err != nil {
       panic(err)
    }

    err = repository.InsertSaloon(connection.DBConnections, saloon)
    if err != nil {
       panic(err)
    }

    s.JSON(http.StatusOK, saloon)
}

func UpdateSaloon(s *gin.Context) {
   var saloon structs.Saloon
   id, _ := strconv.Atoi(s.Param("id"))

   err := s.BindJSON(&saloon)
   if err != nil {
       s.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
       return
   }

   saloon.ID = id


   err = repository.UpdateSaloon(connection.DBConnections, saloon)
   if err != nil {
       s.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
       return
   }

   s.JSON(http.StatusOK, saloon)
}



func DeleteSaloon(s *gin.Context) {
    var saloon structs.Saloon
    id, _ := strconv.Atoi(s.Param("id"))

    saloon.ID = id
    err := repository.DeleteSaloon(connection.DBConnections, saloon)
    if err != nil {
       panic(err)
    }

    s.JSON(http.StatusOK, saloon)
}

func GetSaloonById(s *gin.Context) {
	id, err := strconv.Atoi(s.Param("id"))
	if err != nil {
		s.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	saloon, err := repository.GetSaloonById(connection.DBConnections, id)
	if err != nil {
		s.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	s.JSON(http.StatusOK, saloon)
}
