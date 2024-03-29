package main

import (
	"github.com/gin-gonic/gin"
	"ticketing.io/src/config"
	"ticketing.io/src/database"
	"ticketing.io/src/middleware"
	"ticketing.io/src/routes"
)

func main() {
	app := gin.Default()

	conf := config.Init()
	db, close := database.Init(conf)

	r := routes.New(db)
	app.GET("/api/tickets", r.GetTickets)
	app.GET("/api/tickets/:id", r.GetTicket)

	auth := app.Group("/api", middleware.RequireAuth)
	{
		auth.POST("/tickets", r.CreateTicket)
		auth.PUT(
			"/tickets/:id",
			middleware.CheckPermission(db),
			r.UpdateTicket,
		)
	}

	defer close()
	app.Run(":8080")

}
