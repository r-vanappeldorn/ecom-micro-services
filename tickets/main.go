package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()

	r.GET("/api/tickets", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Test",
		})
	})

	r.Run(":8080")
}