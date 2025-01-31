package controllers

import (
	"log"
	"net/http"
	"proyek-akhir/database/connections"
	"proyek-akhir/repository"
	"proyek-akhir/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllReservation(r *gin.Context) {
    var (
       result gin.H
    )

    saloon, err := repository.GetAllReservation(connection.DBConnections)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": saloon,
       }
    }

    r.JSON(http.StatusOK, result)
}

func InsertReservation(r *gin.Context) {
    var reservation structs.Reservation

    err := r.BindJSON(&reservation)
    if err != nil {
       panic(err)
    }

    err = repository.InsertReservation(connection.DBConnections, reservation)
    if err != nil {
       panic(err)
    }

    r.JSON(http.StatusOK, reservation)
}

func CancelReservation(r *gin.Context) {
   var reservation structs.Reservation

   id, err := strconv.Atoi(r.Param("id"))
   if err != nil || id < 0 { 
       log.Println("Invalid reservation ID:", id, "Error:", err)
       r.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
       return
   }

   reservation.ID = id

   err = repository.CancelReservation(connection.DBConnections, reservation)
   if err != nil {
       log.Println("Error updating reservation:", err)
       r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
       return
   }

   log.Println("Successfully cancel as done:", reservation)
   r.JSON(http.StatusOK, reservation)
}

func DoneReservation(r *gin.Context) {
    var reservation structs.Reservation


    id, err := strconv.Atoi(r.Param("id"))
    if err != nil || id < 0 {
        log.Println("Invalid reservation ID:", id, "Error:", err)
        r.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
        return
    }

    err = r.BindJSON(&reservation)
    if err != nil {
        log.Println("JSON Binding Error:", err)
        r.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
        return
    }
    reservation.ID = id

    err = repository.DoneReservation(connection.DBConnections, reservation)
    if err != nil {
        log.Println("Error updating reservation:", err)
        r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    log.Println("Successfully marked reservation as done, ID:", id)
    r.JSON(http.StatusOK, gin.H{"message": "Reservation marked as done", "id": id, "feedback": reservation.Feedback})
}


func GetAllReservationByCustomerID(r *gin.Context) {
   var (
       result gin.H
   )

   customerID := r.Param("id")

   reservation, err := repository.GetAllReservationByCustomerID(connection.DBConnections, customerID)

   if err != nil {
       result = gin.H{
           "result": err.Error(),
       }
   } else {
       result = gin.H{
           "result": reservation,
       }
   }

   r.JSON(http.StatusOK, result)
}
