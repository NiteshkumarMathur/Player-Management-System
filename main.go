package main

import (
	"github.com/gin-gonic/gin"
)

type Footballer struct {
	UUID    int    `json:"uuid"`
	Name    string `json:"name"`
	Club    string `json:"club"`
	Country string `json:"country"`
}

var footballers = []Footballer{
	{UUID: 1, Name: "Messi", Club: "inter maimi", Country: "argentina"},
	{UUID: 2, Name: "Ronaldo", Club: "al nassr", Country: "portugal"},
	{UUID: 3, Name: "Neymar", Club: "barcelona", Country: "brazil"},
}
var current_id = 3

func main() {
	// Create a new Gin router
	router := gin.New()

	// Define a route for the root URL
	router.GET("/footballer", GetAllFootballers)

	router.POST("/footballer", PostNewFootballer)

	router.DELETE("/footballer/:UUID", DeleteFootballer)

	router.PUT("/footballer/:UUID", UpdateFootballerByUuid)

	router.GET("/search/:key", SearchFootballerAny)

	// Run the server on port 8080
	router.Run(":8080")
}
