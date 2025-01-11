package main

import (
	"eventBooking.com/event-booking/db"
	"eventBooking.com/event-booking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080

}
