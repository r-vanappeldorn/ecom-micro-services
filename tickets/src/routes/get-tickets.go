package routes

import (
	"github.com/gin-gonic/gin"
	"ticketing.io/src/customerrors"
	"ticketing.io/src/models"
)

func (r *Routes) GetTickets(ctx *gin.Context) {
	m := models.Tickets{}
	res, err := m.GetTicketsCollection().FindAll(r.DB)
	if err != nil {
		customerrors.SendInternalServerError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"tickets": res})
}
