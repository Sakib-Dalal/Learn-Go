package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/rest_api/models"
)

func main() {
	server := gin.Default() // setup http server

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE // .GET("path", function())
	server.POST("/events", createEvents)

	server.Run(":8080") // listen to localhost:8080

}

func getEvents(context *gin.Context) {
	fmt.Println("Server Started :)")
	// context.HTML()
	events := models.GetAllEvents()
	context.JSON(200, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(400, gin.H{"message": "Could Not parse the data"})
		return
	}

	event.ID = 1
	event.UserID = 01

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

// var DB *sql.DB
 
// func InitDB() {
//     var err error
//     DB, err = sql.Open("sqlite3", "api.db")
 
//     if err != nil {
//         panic("Could not connect to database.")
//     }
 
//     DB.SetMaxOpenConns(10)
//     DB.SetMaxIdleConns(5)
 
//     createTables()
// }