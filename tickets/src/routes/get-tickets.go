package routes

import (
	"github.com/gin-gonic/gin"
	"ticketing.io/src/customerrors"
	"ticketing.io/src/models"
)

func (r *Routes) GetTickets(ctx *gin.Context) {
	m := models.Tickets{}
	_, err := m.GetTicketsCollection().FindAll(r.DB)
	if err != nil {
		customerrors.SendInternalServerError(ctx, err)
		return
	}

	customerrors.SendBadRequestError(ctx, "invalid credentials")
}
