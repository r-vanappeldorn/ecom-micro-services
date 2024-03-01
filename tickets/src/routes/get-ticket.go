package routes

import (
	"github.com/gin-gonic/gin"
	"ticketing.io/src/models"
)

func (r *Routes) GetTicket(ctx *gin.Context) {
	// Get ticket
	id := ctx.Param("id")

	ticket := &models.Tickets{}
	ticket.FindByID(r.DB, id)

	ctx.JSON(200, gin.H{"ticket": ticket})
}
